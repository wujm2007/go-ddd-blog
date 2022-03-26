package persist

import (
	"gorm.io/gorm"

	"go-ddd-blog/internal/domain/event"
	"go-ddd-blog/internal/infra/persist"
	"go-ddd-blog/internal/service/user/domain/aggregate/user"
)

type UserPO struct {
	persist.Model

	Name     string
	Password string

	entity user.User `gorm:"-:all"`
}

func (po UserPO) TableName() string {
	return "users"
}

func (po UserPO) AfterSave(db *gorm.DB) error {
	return event.Dispatch(db.Statement.Context, po.entity.Events())
}

func marshalUser(user user.User) *UserPO {
	return &UserPO{
		Model: persist.Model{
			ID:        user.ID(),
			CreatedAt: user.CreatedAt(),
		},
		Name:     user.Name(),
		Password: user.Password(),
	}
}

func unmarshalUser(po UserPO, messages []MessagePO) *user.User {
	msgs := make([]user.Message, len(messages))
	for i, m := range messages {
		msgs[i] = *unmarshalMessage(m)
	}
	return user.UnmarshalUser(po.ID, po.Name, po.Password, msgs, po.CreatedAt, po.UpdatedAt)
}

type MessagePO struct {
	persist.Model

	UserID  int64
	Content string

	entity user.Message
}

func (po MessagePO) AfterSave(db *gorm.DB) error {
	return event.Dispatch(db.Statement.Context, po.entity.Events())
}

func (po MessagePO) TableName() string {
	return "messages"
}

func marshalMessage(entity user.Message) *MessagePO {
	return &MessagePO{
		Model: persist.Model{
			ID:        entity.ID(),
			CreatedAt: entity.CreatedAt(),
		},
		UserID:  entity.UserID(),
		Content: entity.Content(),
	}
}

func unmarshalMessage(po MessagePO) *user.Message {
	return user.UnmarshalMessage(po.ID, po.UserID, po.Content, po.CreatedAt, po.UpdatedAt)
}
