package command

import (
	"context"

	"go-ddd-blog/internal/errors"
	"go-ddd-blog/internal/service/post/domain"
	"go-ddd-blog/internal/service/post/domain/aggregate/post"
	domainevent "go-ddd-blog/internal/service/post/domain/event"
)

type command struct {
	repo domain.Repository
}

func NewCommand(repo domain.Repository) Command {
	return &command{repo: repo}
}

func (c command) Create(ctx context.Context, userID int64, title, content string) (id int64, err error) {
	newPost, err := post.NewPost(userID, title, content)
	if err != nil {
		return 0, err
	}
	err = c.repo.Save(ctx, *newPost)
	if err != nil {
		return 0, err
	}
	return newPost.ID(), nil
}

func (c command) Update(ctx context.Context, userID int64, id int64, title, content string) (err error) {
	gotPost, err := c.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if gotPost.UserID() != userID {
		return errors.ErrUnauthorized
	}

	if err := gotPost.SetTitle(title); err != nil {
		return err
	}

	if err := gotPost.SetContent(content); err != nil {
		return err
	}

	return c.repo.Save(ctx, *gotPost)
}

func (c command) AddComment(ctx context.Context, id int64, userID int64, content string) (err error) {
	gotPost, err := c.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	err = gotPost.AddComment(userID, content)
	if err != nil {
		return err
	}

	gotPost.AddEvent(domainevent.NewCommentAddedEvent(gotPost.ID(), gotPost.LastComment().ID()))

	return c.repo.Save(ctx, *gotPost)
}
