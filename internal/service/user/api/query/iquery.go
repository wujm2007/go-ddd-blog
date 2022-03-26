package query

import (
	"context"

	"go-ddd-blog/internal/service/user/api/model"
)

type Query interface {
	GetByNamePassword(ctx context.Context, name, password string) (user *model.UserDTO, err error)
	Get(ctx context.Context, id int64) (user *model.UserDTO, err error)
}
