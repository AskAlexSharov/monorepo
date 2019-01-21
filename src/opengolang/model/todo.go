package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo"

	"github.com/AskAlexSharov/monorepo/src/todo/api/todo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoModel struct {
	Db *dgo.Dgraph
}

func (s TodoModel) List(ctx context.Context, req *todo.ListRequest) (*todo.ListResponse, error) {
	tx := s.Db.NewReadOnlyTxn()
	defer tx.Discard(ctx)
	// Query the balance for Alice and Bob.
	const q = `
		{
			all(func: anyofterms(name, "Alice Bob")) {
				uid
				balance
			}
		}
	`

	resp, err := tx.Query(ctx, q)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}

	// After we get the balances, we have to decode them into structs so that
	// we can manipulate the data.
	var decode struct {
		All []struct {
			Uid     string
			Balance int
		}
	}
	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}
	fmt.Printf("From Dgraph %+v\n", decode)
	return nil, nil
	//return &todo.ListResponse{Items: items}, nil
}

//// CreateTodo creates a todo given a description
//func (s TodoModel) Insert(ctx context.Context, item *todo.Todo) (string, error) {
//	item.Id = primitive.NewObjectID().String()
//
//	if _, err := s.Collection.InsertOne(ctx, item); err != nil {
//		return "", status.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
//	}
//
//	return item.Id, nil
//}
//
//// GetTodo retrieves a todo item from its ID
//func (s TodoModel) FindById(ctx context.Context, id string, result *todo.Todo) error {
//	filter := bson.M{"_id": id}
//	if err := s.Collection.FindOne(ctx, filter).Decode(result); err != nil {
//		return status.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
//	}
//	return nil
//}

//func (s TodoModel) RemoveById(ctx context.Context, id string) error {
//	filter := bson.M{"_id": id}
//	_, err := s.Collection.DeleteOne(ctx, filter)
//	if err != nil {
//		return status.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
//	}
//	return nil
//}
//
//func (s TodoModel) UpdateById(ctx context.Context, id string, data map[string]interface{}) error {
//	filter := bson.M{"_id": id}
//	data["updated_at"] = time.Now()
//	_, err := s.Collection.UpdateOne(ctx, filter, data)
//	if err != nil {
//		return status.Errorf(codes.Internal, "Could not update item from the database: %s", err)
//	}
//
//	return nil
//}
