package user

import (
	"time"

	"go-ddd-blog/internal/domain/entity"
	"go-ddd-blog/internal/domain/idgen"
	"go-ddd-blog/internal/errors"
)

type Message struct {
	entity.Entity

	userID  int64
	content string

	createdAt time.Time
	updatedAt time.Time
}

func (m Message) UserID() int64 {
	return m.userID
}

func (m Message) Content() string {
	return m.content
}

func (m Message) CreatedAt() time.Time {
	return m.createdAt
}

func (m Message) UpdatedAt() time.Time {
	return m.updatedAt
}

func UnmarshalMessage(userID int64, id int64, content string, createdAt time.Time, updatedAt time.Time) *Message {
	return &Message{Entity: entity.NewEntity(id), userID: userID, content: content, createdAt: createdAt, updatedAt: updatedAt}
}

func newMessage(userID int64, content string) (*Message, error) {
	if userID <= 0 {
		return nil, errors.ErrBadRequest.WithMessage("invalid user_id")
	}

	c := &Message{userID: userID, content: content}
	id, err := idgen.GenerateID()
	if err != nil {
		return nil, errors.ErrInternal.WithMessage("failed to generate id").WithCause(err)
	}
	c.Entity = entity.NewEntity(id)

	return c, nil
}
