package event_handler

import (
	"context"

	"go-ddd-blog/internal/domain/event"
)

type CommentAddedHandler interface {
	Handle(ctx context.Context, e event.DomainEvent) error
}
