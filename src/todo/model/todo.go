package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"log"
	"time"

	"github.com/AskAlexSharov/monorepo/src/todo/api/todo"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoModel struct {
	Collection *mongo.Collection
}

// CreateTodo creates a todo given a description
func (s TodoModel) Insert(ctx context.Context, item *todo.Todo) (string, error) {
	item.Id = primitive.NewObjectID().String()

	if _, err := s.Collection.InsertOne(ctx, item); err != nil {
		return "", status.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}

	return item.Id, nil
}

// GetTodo retrieves a todo item from its ID
func (s TodoModel) FindById(ctx context.Context, id string, result *todo.Todo) error {
	filter := bson.M{"_id": id}
	if err := s.Collection.FindOne(ctx, filter).Decode(result); err != nil {
		return status.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}
	return nil
}

func (s TodoModel) List(ctx context.Context, req *todo.ListRequest) (*todo.ListResponse, error) {
	var items []*todo.Todo
	filter := bson.M{}
	//if req.Limit > 0 {
	//	query.Limit(int(req.Limit))
	//}
	//if req.NotCompleted {
	//	query.Where("completed = false")
	//}
	//err := query.Select()

	cur, err := s.Collection.Find(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
	}
	defer cur.Close(ctx)
	for i := 0; cur.Next(ctx); i++ {
		var result todo.Todo
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		items[i] = &result
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &todo.ListResponse{Items: items}, nil
}

func (s TodoModel) RemoveById(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := s.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
	}
	return nil
}

func (s TodoModel) UpdateById(ctx context.Context, id string, data map[string]interface{}) error {
	filter := bson.M{"_id": id}
	data["updated_at"] = time.Now()
	_, err := s.Collection.UpdateOne(ctx, filter, data)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not update item from the database: %s", err)
	}

	return nil
}
