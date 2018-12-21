//go:generate go run scripts/gqlgen.go -v

package chat

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

type resolver struct {
	Rooms map[string]*Chatroom
	mu    sync.Mutex // nolint: structcheck
}

func (r *resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

func New() Config {
	return Config{
		Resolvers: &resolver{
			Rooms: map[string]*Chatroom{},
		},
	}
}

type Chatroom struct {
	Name      string
	messages  []Message
	Observers map[string]chan Message
}

type Message struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	Name string `json:"name"`
}

type mutationResolver struct{ *resolver }

func (r *mutationResolver) Post(ctx context.Context, text string, username string, roomName string) (Message, error) {
	room := r.Rooms[roomName]
	if room == nil {
		room = &Chatroom{Name: roomName, Observers: map[string]chan Message{}}
		r.Rooms[roomName] = room
	}

	message := Message{
		ID:        randString(8),
		CreatedAt: time.Now(),
		Text:      text,
		CreatedBy: username,
	}

	room.messages = append(room.messages, message)
	for _, observer := range room.Observers {
		observer <- message
	}
	return message, nil
}

type queryResolver struct{ *resolver }

func (r *Message) User(ctx context.Context) (*User, error) {
	return UserLoaderFromCtx(ctx).Load(r.CreatedBy)
}

func (r *Message) User2(ctx context.Context) (*User, error) {
	return UserLoaderFromCtx(ctx).Load(r.CreatedBy)
}

func (r *Message) User3(ctx context.Context) (*User, error) {
	return UserLoaderFromCtx(ctx).Load(r.CreatedBy)
}

func (r *Message) User4(ctx context.Context) (*User, error) {
	return UserLoaderFromCtx(ctx).Load(r.CreatedBy)
}

func (r *Message) User5(ctx context.Context) (*User, error) {
	return UserLoaderFromCtx(ctx).Load(r.CreatedBy)
}

func (r *queryResolver) Room(ctx context.Context, name string) (*Chatroom, error) {
	r.mu.Lock()
	room := r.Rooms[name]
	if room == nil {
		room = &Chatroom{Name: name, Observers: map[string]chan Message{}}
		r.Rooms[name] = room
	}
	r.mu.Unlock()

	return room, nil
}

func (r *queryResolver) User(ctx context.Context, name string) (*User, error) {
	return UserLoaderFromCtx(ctx).Load(name)
}

func (r *Chatroom) Messages(ctx context.Context) ([]Message, error) {
	return r.messages, nil
}

type subscriptionResolver struct{ *resolver }

func (r *subscriptionResolver) MessageAdded(ctx context.Context, roomName string) (<-chan Message, error) {
	r.mu.Lock()
	room := r.Rooms[roomName]
	if room == nil {
		room = &Chatroom{Name: roomName, Observers: map[string]chan Message{}}
		r.Rooms[roomName] = room
	}
	r.mu.Unlock()

	id := randString(8)
	events := make(chan Message, 1)

	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(room.Observers, id)
		r.mu.Unlock()
	}()

	r.mu.Lock()
	room.Observers[id] = events
	r.mu.Unlock()

	return events, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
