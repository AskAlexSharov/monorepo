//go:generate dataloaden -keys string github.com/nizsheanez/monorepo/src/chat.User
//go:generate dataloaden -keys string github.com/nizsheanez/monorepo/src/chat.Message

package chat

import (
	"context"
	"time"
)

type userCtxKeyT struct{}

var userCtxKey userCtxKeyT

func UserLoaderFromCtx(ctx context.Context) *UserLoader {
	return ctx.Value(userCtxKey).(*UserLoader)
}

func UserLoaderToCtx(ctx context.Context, l *UserLoader) context.Context {
	return context.WithValue(ctx, userCtxKey, l)
}

func NewUserLoader() *UserLoader {
	return &UserLoader{
		wait:     1 * time.Microsecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*User, []error) {

			users := make([]*User, len(keys))
			errors := make([]error, len(keys))

			for i, _ := range keys {
				users[i] = &User{Name: "user"}
			}
			return users, errors
		},
	}
}

type msgCtxKeyT struct{}

var msgCtxKey msgCtxKeyT

func MsgsLoaderFromCtx(ctx context.Context) *MessageLoader {
	return ctx.Value(msgCtxKey).(*MessageLoader)
}

func MsgsLoaderToCtx(ctx context.Context, l *MessageLoader) context.Context {
	return context.WithValue(ctx, msgCtxKey, l)
}

func NewMsgsLoader() *MessageLoader {
	return &MessageLoader{
		//Rooms map[string]*Chatroom
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Message, []error) {
			messages := make([]*Message, len(keys))
			errors := make([]error, len(keys))

			for i, _ := range keys {
				messages[i] = &Message{}
			}
			return messages, errors
		},
	}
}
