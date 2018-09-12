package service

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/nizsheanez/monorepo/todo/model"
	"github.com/nizsheanez/monorepo/todo/sdk/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"time"
)

type TodoService struct {
	Model *model.TodoModel
}

func (s TodoService) Create(ctx context.Context, req *todo_sdk.CreateRequest) (*todo_sdk.CreateResponse, error) {
	ids, err := s.Model.Insert(ctx, []*todo_sdk.Todo{req.Item})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}

	if len(ids) == 0 {
		return nil, status.Errorf(codes.Internal, "Created 0 entities")
	}

	return &todo_sdk.CreateResponse{Id: ids[0]}, nil
}

func (s TodoService) CreateBulk(ctx context.Context, req *todo_sdk.CreateBulkRequest) (*todo_sdk.CreateBulkResponse, error) {
	ids, err := s.Model.Insert(ctx, req.Items)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not insert items into the database: %s", err)
	}
	return &todo_sdk.CreateBulkResponse{Ids: ids}, nil
}

func (s TodoService) Get(ctx context.Context, req *todo_sdk.GetRequest) (*todo_sdk.GetResponse, error) {
	var item todo_sdk.Todo
	{ // find it
		if err := s.Model.FindById(ctx, req.Id, &item); err != nil {
			return nil, status.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
		}
	}
	return &todo_sdk.GetResponse{Item: &item}, nil
}

func (s TodoService) List(ctx context.Context, req *todo_sdk.ListRequest) (*todo_sdk.ListResponse, error) {
	grpclog.Info("asdf")
	resp, err := s.Model.List(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
	}
	return resp, nil
}

func (s TodoService) Delete(ctx context.Context, req *todo_sdk.DeleteRequest) (*todo_sdk.DeleteResponse, error) {
	if err := s.Model.RemoveById(ctx, req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
	}
	return &todo_sdk.DeleteResponse{}, nil
}

func (s TodoService) Update(ctx context.Context, req *todo_sdk.UpdateRequest) (*todo_sdk.UpdateResponse, error) {
	var err error
	req.Item.UpdatedAt, err = ptypes.TimestampProto(time.Now())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't create timestamp")
	}

	if err := s.Model.UpdateById(ctx, req.Item.Id, req.Item); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update item from the database: %s", err)
	}
	return &todo_sdk.UpdateResponse{}, nil
}

func (s TodoService) UpdateBulk(ctx context.Context, req *todo_sdk.UpdateBulkRequest) (*todo_sdk.UpdateBulkResponse, error) {
	for _, item := range req.Items {
		if err := s.Model.UpdateById(ctx, item.Id, item); err != nil {
			return nil, status.Errorf(codes.Internal, "Could not update items from the database: %s", err)
		}
	}

	return &todo_sdk.UpdateBulkResponse{}, nil
}
