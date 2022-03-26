package domain

import (
	"context"

	"go-ddd-blog/internal/service/post/domain/aggregate/post"
)

type Repository interface {
	Save(ctx context.Context, post post.Post) (err error)
	Get(ctx context.Context, id int64) (post *post.Post, err error)
}
