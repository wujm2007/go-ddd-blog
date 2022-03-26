package domain

import (
	"context"

	"go-ddd-blog/internal/service/user/domain/aggregate/user"
)

type Repository interface {
	Save(ctx context.Context, user user.User) (id int64, err error)
	Get(ctx context.Context, id int64) (user *user.User, err error)
}
