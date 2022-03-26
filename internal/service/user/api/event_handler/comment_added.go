package event_handler

import (
	"context"
	"fmt"

	"go-ddd-blog/internal/domain/event"
	postQuery "go-ddd-blog/internal/service/post/api/query"
	events "go-ddd-blog/internal/service/post/domain/event"
	userCmd "go-ddd-blog/internal/service/user/api/command"
)

type commentNotificationHandler struct {
	userCommand userCmd.Command
	postQuery   postQuery.Query
}

func NewCommentNotificationHandler(userCmd userCmd.Command, postQuery postQuery.Query) CommentAddedHandler {
	return &commentNotificationHandler{userCommand: userCmd, postQuery: postQuery}
}

func (c commentNotificationHandler) Handle(ctx context.Context, t event.DomainEvent) error {
	e := t.(*events.CommentAddedEvent)
	post, err := c.postQuery.Get(ctx, e.PostID(), false)
	if err != nil {
		return err
	}
	return c.userCommand.CreateMessage(ctx, post.ID, fmt.Sprintf(
		"[New Comment] %v: %v",
		post.UserID,
		post.Comments,
	))
}
