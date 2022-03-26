package user

import (
	"time"

	"go-ddd-blog/internal/domain/entity"
	"go-ddd-blog/internal/errors"
)

type User struct {
	entity.Entity

	name     string
	password string

	messages []Message

	createdAt time.Time
	updatedAt time.Time
}

func (u *User) Messages() []Message {
	return u.messages
}

func UnmarshalUser(id int64, name string, password string, messages []Message, createdAt, updatedAt time.Time) *User {
	return &User{
		Entity:    entity.NewEntity(id),
		name:      name,
		password:  password,
		messages:  messages,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Name() string {
	return u.name
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) SetName(name string) error {
	if len(name) == 0 {
		return errors.ErrBadRequest.WithMessage("empty name")
	}
	u.name = name
	return nil
}

func (u *User) SetPassword(password string) error {
	if len(password) < 6 {
		return errors.ErrBadRequest.WithMessage("invalid password")
	}
	u.password = password
	return nil
}

func (u *User) AddMessage(userID int64, content string) error {
	m, err := newMessage(userID, content)
	if err != nil {
		return err
	}

	u.messages = append(u.messages, *m)
	return nil
}

func NewUser(name, password string) (*User, error) {
	user := &User{}
	if err := user.SetName(name); err != nil {
		return nil, err
	}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	return user, nil
}
