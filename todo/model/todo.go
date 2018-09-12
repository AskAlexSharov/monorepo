package model

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"github.com/nizsheanez/monorepo/todo/sdk/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type TodoModel struct {
	Collection *mgo.Collection
}

// CreateTodo creates a todo given a description
func (s TodoModel) Insert(ctx context.Context, items []*todo_sdk.Todo) ([]string, error) {
	for _, item := range items {
		item.Id = bson.NewObjectId().String()
	}

	err := s.Collection.Insert(items)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}

	ids := make([]string, len(items))
	for i := range items {
		ids[i] = items[i].Id
	}
	return ids, nil
}

// CreateTodos create todo items from a list of todo descriptions
func (s TodoModel) CreateBulk(ctx context.Context, req *todo_sdk.CreateBulkRequest) (*todo_sdk.CreateBulkResponse, error) {
	var ids []string
	for _, item := range req.Items {
		item.Id = bson.NewObjectId().String()
		ids = append(ids, item.Id)
	}
	//err := s.collection.Insert(&req.Items)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "Could not insert items into the database: %s", err)
	//}
	return &todo_sdk.CreateBulkResponse{Ids: ids}, nil
}

// GetTodo retrieves a todo item from its ID
func (s TodoModel) FindById(ctx context.Context, id string, result *todo_sdk.Todo) error {
	if err := s.Collection.FindId(id).One(result); err != nil {
		return status.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}
	return nil
}

func (s TodoModel) List(ctx context.Context, req *todo_sdk.ListRequest) (*todo_sdk.ListResponse, error) {
	var items []*todo_sdk.Todo
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
	return &todo_sdk.ListResponse{Items: items}, nil
}

func (s TodoModel) RemoveById(ctx context.Context, id string) error {
	err := s.Collection.RemoveId(id)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
	}
	return nil
}

func (s TodoModel) UpdateById(ctx context.Context, id string, item *todo_sdk.Todo) error {
	var err error
	item.UpdatedAt, err = ptypes.TimestampProto(time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Can't create timestamp")
	}

	if err := s.Collection.UpdateId(id, item); err != nil {
		return status.Errorf(codes.Internal, "Could not update item from the database: %s", err)
	}
	return nil
}
