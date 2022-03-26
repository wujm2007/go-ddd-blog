package post

import (
	"time"

	"go-ddd-blog/internal/domain/entity"

	"go-ddd-blog/internal/domain/idgen"

	"go-ddd-blog/internal/errors"
)

type Comment struct {
	entity.Entity

	userID int64

	content string

	createdAt time.Time
	updatedAt time.Time
}

func UnmarshalComment(userID int64, id int64, content string, createdAt, updatedAt time.Time) *Comment {
	return &Comment{Entity: entity.NewEntity(id), userID: userID, content: content, createdAt: createdAt, updatedAt: updatedAt}
}

func (c Comment) UserID() int64 {
	return c.userID
}

func (c Comment) Content() string {
	return c.content
}

func (c Comment) CreatedAt() time.Time {
	return c.createdAt
}

func (c Comment) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Comment) SetContent(content string) error {
	if len(content) == 0 {
		return errors.ErrBadRequest.WithMessage("empty content")
	}
	c.content = content
	return nil
}

func newComment(userID int64, content string) (*Comment, error) {
	if userID <= 0 {
		return nil, errors.ErrBadRequest.WithMessage("invalid user_id")
	}

	c := &Comment{userID: userID}
	if err := c.SetContent(content); err != nil {
		return nil, err
	}
	id, err := idgen.GenerateID()
	if err != nil {
		return nil, errors.ErrInternal.WithMessage("failed to generate id").WithCause(err)
	}
	c.Entity = entity.NewEntity(id)

	return c, nil
}
