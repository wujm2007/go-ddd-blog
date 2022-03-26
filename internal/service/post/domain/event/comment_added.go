package event

import (
	"context"
	"fmt"

	"go-ddd-blog/internal/domain/event"
)

type CommentAddedEvent struct {
	postID    int64
	commentID int64
}

func (p CommentAddedEvent) PostID() int64 {
	return p.postID
}

func (p CommentAddedEvent) CommentID() int64 {
	return p.commentID
}

func (p CommentAddedEvent) EventKey() string {
	return fmt.Sprintf("%T", p)
}

func NewCommentAddedEvent(postID int64, commentID int64) *CommentAddedEvent {
	return &CommentAddedEvent{postID: postID, commentID: commentID}
}

func RegisterPostNewCommentEventHandler(h func(ctx context.Context, t event.DomainEvent) error) {
	event.Register(&CommentAddedEvent{}, h)
}
