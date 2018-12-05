package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var KafkaNodes = []string{"localhost:32788", "localhost:32789", "localhost:32790"}

const TOPIC = "my-topic"

var writeCnt int32 = 0
var readCnt int32 = 0

func main() {
	createTopicIfNeed()

	ProducersPoolSize := 5
	ConsumersPoolSize := 100

	for i := 0; i < ProducersPoolSize; i++ {
		go producer(i)
	}

	time.Sleep(10 * time.Second)

	wg := sync.WaitGroup{}
	for i := 0; i < ConsumersPoolSize; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			consumer(id)
		}(i)
	}

	wg.Wait()
}
func createTopicIfNeed() {
	c, err := kafka.Dial("tcp", KafkaNodes[0])
	if err != nil {
		panic(err)
	}

	_ = c.DeleteTopics(TOPIC)
	time.Sleep(time.Second)

	err = c.CreateTopics(kafka.TopicConfig{
		Topic:             TOPIC,
		NumPartitions:     10,
		ReplicationFactor: 1,
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
	time.Sleep(time.Second)
}

func producer(id int) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          KafkaNodes,
		Topic:            TOPIC,
		Balancer:         &kafka.LeastBytes{},
		QueueCapacity:    1e5,
		BatchTimeout:     time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	})
	defer w.Close()

	var batchSize int = 1e5
	for j := 0; j < 1e3; j++ {
		messages := make([]kafka.Message, batchSize)
		for i := 0; i < batchSize; i++ {
			k := strconv.Itoa(id) + "-" + strconv.Itoa(i)
			messages[i] = kafka.Message{
				Key:   []byte("Key-A-" + k),
				Value: []byte("Val-A-" + k),
			}
		}

		if err := w.WriteMessages(context.Background(), messages...); err != nil {
			panic(err)
		}
		atomic.AddInt32(&writeCnt, int32(batchSize))

		if j%100 == 0 {
			fmt.Printf("producer %d %d\n", id, j)
		}
	}

	fmt.Println("write done")
}

func consumer(id int) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  KafkaNodes,
		GroupID:  "consumer-group-id",
		Topic:    TOPIC,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()

	cnt := 0
	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			panic(err)
		}

		time.Sleep(10 * time.Millisecond)

		cnt++
		if cnt%100 == 0 {
			atomic.AddInt32(&readCnt, int32(100))
			fmt.Printf("Consumer %d %d\n", id, cnt)
		}

		if err := r.CommitMessages(context.Background(), m); err != nil {
			panic(err)
		}
	}

}
