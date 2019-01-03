package model

import (
	"context"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/nizsheanez/monorepo/src/todo/api/todo/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoModel struct {
	Collection *mgo.Collection
}

// CreateTodo creates a todo given a description
func (s TodoModel) Insert(ctx context.Context, item *todo.Todo) (string, error) {
	item.Id = bson.NewObjectId().String()

	if err := s.Collection.Insert(item); err != nil {
		return "", status.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}

	return item.Id, nil
}

// GetTodo retrieves a todo item from its ID
func (s TodoModel) FindById(ctx context.Context, id string, result *todo.Todo) error {
	if err := s.Collection.FindId(id).One(result); err != nil {
		return status.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}
	return nil
}

func (s TodoModel) List(ctx context.Context, req *todo.ListRequest) (*todo.ListResponse, error) {
	var items []*todo.Todo
	request := bson.M{}
	//if req.Limit > 0 {
	//	query.Limit(int(req.Limit))
	//}
	//if req.NotCompleted {
	//	query.Where("completed = false")
	//}
	//err := query.Select()

	if err := s.Collection.Find(request).All(&items); err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
	}
	return &todo.ListResponse{Items: items}, nil
}

func (s TodoModel) RemoveById(ctx context.Context, id string) error {
	if err := s.Collection.RemoveId(id); err != nil {
		return status.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
	}
	return nil
}

func (s TodoModel) UpdateById(ctx context.Context, id string, data map[string]interface{}) error {
	data["updated_at"] = time.Now()
	if err := s.Collection.UpdateId(id, data); err != nil {
		return status.Errorf(codes.Internal, "Could not update item from the database: %s", err)
	}
	return nil
}
