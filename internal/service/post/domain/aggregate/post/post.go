package post

import (
	"time"

	"go-ddd-blog/internal/domain/entity"
	"go-ddd-blog/internal/domain/idgen"
	"go-ddd-blog/internal/errors"
)

type Post struct {
	entity.Entity

	userID int64

	title   string
	content string

	comments []Comment

	createdAt time.Time
	updatedAt time.Time
}

func UnmarshalPost(userID int64, id int64, title string, content string, comments []Comment, createdAt, updatedAt time.Time) *Post {
	return &Post{
		Entity:    entity.NewEntity(id),
		userID:    userID,
		title:     title,
		content:   content,
		comments:  comments,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewPost(userID int64, title string, content string) (*Post, error) {
	if userID <= 0 {
		return nil, errors.ErrBadRequest.WithMessage("invalid user_id")
	}
	newPost := &Post{userID: userID}

	if err := newPost.SetTitle(title); err != nil {
		return nil, err
	}
	if err := newPost.SetContent(content); err != nil {
		return nil, err
	}

	if id, err := idgen.GenerateID(); err != nil {
		return nil, errors.ErrInternal.WithMessage("failed to generate id").WithCause(err)
	} else {
		newPost.Entity = entity.NewEntity(id)
	}
	return newPost, nil
}

func (p Post) UserID() int64 {
	return p.userID
}

func (p Post) Title() string {
	return p.title
}

func (p Post) Content() string {
	return p.content
}

func (p Post) Comments() []Comment {
	return p.comments
}

func (p Post) LastComment() *Comment {
	if len(p.comments) == 0 {
		return nil
	}
	return &p.comments[len(p.comments)-1]
}

func (p *Post) SetTitle(title string) error {
	if len(title) == 0 {
		return errors.ErrBadRequest.WithMessage("empty title")
	}
	p.title = title
	return nil
}

func (p *Post) SetContent(content string) error {
	if len(content) == 0 {
		return errors.ErrBadRequest.WithMessage("empty content")
	}
	p.content = content
	return nil
}

func (p Post) CreatedAt() time.Time {
	return p.createdAt
}

func (p Post) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Post) AddComment(userID int64, content string) error {
	c, err := newComment(userID, content)
	if err != nil {
		return err
	}

	p.comments = append(p.comments, *c)
	return nil
}
