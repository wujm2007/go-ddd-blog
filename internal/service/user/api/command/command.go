package command

import (
	"context"

	"go-ddd-blog/internal/service/user/domain"
	"go-ddd-blog/internal/service/user/domain/aggregate/user"
)

type command struct {
	repo domain.Repository
}

func NewCommand(repo domain.Repository) Command {
	return &command{repo: repo}
}

func (c command) Create(ctx context.Context, name, password string) (id int64, err error) {
	newUser, err := user.NewUser(name, password)
	if err != nil {
		return 0, err
	}

	return c.repo.Save(ctx, *newUser)
}

func (c command) Update(ctx context.Context, id int64, name string) (err error) {
	gotUser, err := c.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	return gotUser.SetName(name)
}

func (c command) CreateMessage(ctx context.Context, userID int64, content string) (err error) {
	gotUser, err := c.repo.Get(ctx, userID)
	if err != nil {
		return err
	}

	if err = gotUser.AddMessage(userID, content); err != nil {
		return err
	}

	_, err = c.repo.Save(ctx, *gotUser)

	return err
}
