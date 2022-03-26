package query

import (
	"context"

	"go-ddd-blog/internal/service/post/api/model"
)

type Query interface {
	Get(ctx context.Context, id int64, withComment bool) (post *model.PostDTO, err error)
	MGetByUserID(ctx context.Context, userID int64) (posts []model.PostDTO, err error)
}
