package command

import (
	"context"
)

type Command interface {
	Create(ctx context.Context, userID int64, title, content string) (id int64, err error)
	Update(ctx context.Context, userID int64, id int64, title, content string) (err error)
	AddComment(ctx context.Context, id int64, userID int64, content string) (err error)
}
