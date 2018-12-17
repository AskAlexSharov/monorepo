//go:generate dataloaden -keys string github.com/nizsheanez/monorepo/src/chat.User

package chat

import (
	"context"
	"fmt"
	"time"
)

var ctxKey = "user_loader"

func UserLoaderFromCtx(ctx context.Context) *UserLoader {
	return ctx.Value(ctxKey).(*UserLoader)
}

func UserLoaderToCtx(ctx context.Context, l *UserLoader) context.Context {
	return context.WithValue(ctx, ctxKey, l)
}

// NewLoader will collect user requests for 2 milliseconds and send them as a single batch to the fetch func
// normally fetch would be a database call.
func NewUserLoader() *UserLoader {
	return &UserLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*User, []error) {
			fmt.Println("Real Load")
			users := make([]*User, len(keys))
			errors := make([]error, len(keys))

			for i, key := range keys {
				users[i] = &User{Name: "user " + key}
			}
			return users, errors
		},
	}
}
