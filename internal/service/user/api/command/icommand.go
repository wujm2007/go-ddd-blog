package command

import (
	"context"
)

type Command interface {
	Create(ctx context.Context, name, password string) (id int64, err error)
	Update(ctx context.Context, id int64, name string) (err error)
	CreateMessage(ctx context.Context, userID int64, content string) (err error)
}
