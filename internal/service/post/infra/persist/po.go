package persist

import (
	"gorm.io/gorm"

	"go-ddd-blog/internal/domain/event"
	"go-ddd-blog/internal/infra/persist"
	"go-ddd-blog/internal/service/post/domain/aggregate/post"
)

type PostPO struct {
	persist.Model

	UserID  int64
	Title   string
	Content string

	entity post.Post `gorm:"-:all"`
}

func (po PostPO) TableName() string {
	return "posts"
}

func (po PostPO) AfterSave(db *gorm.DB) error {
	return event.Dispatch(db.Statement.Context, po.entity.Events())
}

func marshalPost(post post.Post) *PostPO {
	return &PostPO{
		UserID: post.UserID(),
		Model: persist.Model{
			ID:        post.ID(),
			CreatedAt: post.CreatedAt(),
		},
		Title:   post.Title(),
		Content: post.Content(),
		entity:  post,
	}
}

func unmarshalPost(po PostPO, commentPOs []CommentPO) *post.Post {
	comments := make([]post.Comment, len(commentPOs))
	for i, commentPO := range commentPOs {
		comments[i] = *unmarshalComment(commentPO)
	}
	return post.UnmarshalPost(po.UserID, po.ID, po.Title, po.Content, comments, po.CreatedAt, po.UpdatedAt)
}

type CommentPO struct {
	persist.Model

	UserID int64
	PostID int64

	Content string

	entity post.Comment `gorm:"-:all"`
}

func (po CommentPO) TableName() string {
	return "comments"
}

func (po CommentPO) AfterSave(db *gorm.DB) error {
	return event.Dispatch(db.Statement.Context, po.entity.Events())
}

func marshalComment(postID int64, comment post.Comment) *CommentPO {
	return &CommentPO{
		UserID: comment.UserID(),
		PostID: postID,
		Model: persist.Model{
			ID:        comment.ID(),
			CreatedAt: comment.CreatedAt(),
		},
		Content: comment.Content(),
		entity:  comment,
	}
}

func unmarshalComment(po CommentPO) *post.Comment {
	return post.UnmarshalComment(po.UserID, po.ID, po.Content, po.CreatedAt, po.UpdatedAt)
}
