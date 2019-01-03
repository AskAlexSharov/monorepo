package service

import (
	"context"

	"github.com/fatih/structs"
	"github.com/nizsheanez/monorepo/src/todo/api/todo/v2"
	"github.com/nizsheanez/monorepo/src/todo/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoService struct {
	Model *model.TodoModel
}

func (s TodoService) Create(ctx context.Context, req *todo.CreateRequest) (*todo.CreateResponse, error) {
	id, err := s.Model.Insert(ctx, req.Item)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}

	return &todo.CreateResponse{Id: id}, nil
}

func (s TodoService) CreateBulk(ctx context.Context, req *todo.CreateBulkRequest) (*todo.CreateBulkResponse, error) {
	ids := make([]string, len(req.Items))

	var err error
	for i, item := range req.Items {
		ids[i], err = s.Model.Insert(ctx, item)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Could not insert items into the database: %s", err)
		}
	}

	return &todo.CreateBulkResponse{Ids: ids}, nil
}

func (s TodoService) Get(ctx context.Context, req *todo.GetRequest) (*todo.GetResponse, error) {
	var item todo.Todo
	{ // find it
		if err := s.Model.FindById(ctx, req.Id, &item); err != nil {
			return nil, status.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
		}
	}
	return &todo.GetResponse{Item: &item}, nil
}

func (s TodoService) List(ctx context.Context, req *todo.ListRequest) (*todo.ListResponse, error) {
	resp, err := s.Model.List(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
	}
	return resp, nil
}

func (s TodoService) Delete(ctx context.Context, req *todo.DeleteRequest) (*todo.DeleteResponse, error) {
	if err := s.Model.RemoveById(ctx, req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
	}
	return &todo.DeleteResponse{}, nil
}

func (s TodoService) Update(ctx context.Context, req *todo.UpdateRequest) (*todo.UpdateResponse, error) {
	if err := s.Model.UpdateById(ctx, req.Item.Id, structs.New(req.Item).Map()); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update item from the database: %s", err)
	}
	return &todo.UpdateResponse{}, nil
}

func (s TodoService) UpdateBulk(ctx context.Context, req *todo.UpdateBulkRequest) (*todo.UpdateBulkResponse, error) {
	for _, item := range req.Items {
		if err := s.Model.UpdateById(ctx, item.Id, structs.New(item).Map()); err != nil {
			return nil, status.Errorf(codes.Internal, "Could not update items from the database: %s", err)
		}
	}

	return &todo.UpdateBulkResponse{}, nil
}
