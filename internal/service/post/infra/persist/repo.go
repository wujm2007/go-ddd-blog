package persist

import (
	"context"

	"gorm.io/gorm"

	"go-ddd-blog/internal/errors"
	"go-ddd-blog/internal/service/post/domain"
	"go-ddd-blog/internal/service/post/domain/aggregate/post"
)

type postRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) domain.Repository {
	return &postRepo{db: db}
}

func (r postRepo) Save(ctx context.Context, post post.Post) error {
	po := marshalPost(post)
	db := r.db.WithContext(ctx)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(po).Error; err != nil {
			return err
		}
		commentPOs := make([]CommentPO, len(post.Comments()))
		for i, comment := range post.Comments() {
			commentPOs[i] = *marshalComment(post.ID(), comment)
		}
		if len(commentPOs) != 0 {
			return tx.Save(commentPOs).Error
		}
		return nil
	})
	return errors.TransformGormErr(err)
}

func (r postRepo) Get(ctx context.Context, id int64) (post *post.Post, err error) {
	var po PostPO
	if err = r.db.WithContext(ctx).First(&po, id).Error; err != nil {
		return nil, errors.TransformGormErr(err)
	}

	var commentPOs []CommentPO
	if err = r.db.WithContext(ctx).Find(&commentPOs, "post_id = ?", po.ID).Error; err != nil {
		return nil, errors.TransformGormErr(err)
	}

	return unmarshalPost(po, commentPOs), nil
}
