package persist

import (
	"context"

	"gorm.io/gorm"

	"go-ddd-blog/internal/errors"
	"go-ddd-blog/internal/service/post/api/model"
	"go-ddd-blog/internal/service/post/api/query"
)

type postQuery struct {
	db *gorm.DB
}

func NewQuery(db *gorm.DB) query.Query {
	return &postQuery{db: db}
}

func (q postQuery) Get(ctx context.Context, id int64, withComments bool) (post *model.PostDTO, err error) {
	var po PostPO
	if err = q.db.WithContext(ctx).First(&po, "id = ?", id).Error; err != nil {
		return nil, errors.TransformGormErr(err)
	}

	var comments []model.CommentDTO
	if withComments {
		var commentPOs []CommentPO
		if err = q.db.WithContext(ctx).Find(&commentPOs, "post_id = ?", po.ID).Error; err != nil {
			return nil, errors.TransformGormErr(err)
		}
		comments = make([]model.CommentDTO, len(commentPOs))
		for i, commentPO := range commentPOs {
			comments[i] = model.CommentDTO{
				UserID:    commentPO.UserID,
				ID:        commentPO.ID,
				Content:   commentPO.Content,
				CreatedAt: commentPO.CreatedAt,
				UpdatedAt: commentPO.UpdatedAt,
			}
		}
	}

	return &model.PostDTO{
		UserID:    po.UserID,
		ID:        po.ID,
		Title:     po.Title,
		Content:   po.Content,
		Comments:  comments,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}, nil

}

func (q postQuery) MGetByUserID(ctx context.Context, userID int64) ([]model.PostDTO, error) {
	var postPOs []PostPO
	if err := q.db.WithContext(ctx).Find(&postPOs, "user_id = ?", userID).Error; err != nil {
		return nil, errors.TransformGormErr(err)
	}
	posts := make([]model.PostDTO, len(postPOs))
	for i, po := range postPOs {
		posts[i] = model.PostDTO{
			UserID:    po.UserID,
			ID:        po.ID,
			Title:     po.Title,
			Content:   po.Content,
			Comments:  nil,
			CreatedAt: po.CreatedAt,
			UpdatedAt: po.UpdatedAt,
		}
	}
	return posts, nil
}
