package event

import "context"

type DomainEvent interface {
	EventKey() string
}

var registered = make(map[string]any)

func Register(e DomainEvent, handler func(context.Context, DomainEvent) error) {
	eventType := e.EventKey()
	if _, ok := registered[eventType]; !ok {
		registered[eventType] = make([]func(context.Context, DomainEvent) error, 0)
	}
	registered[eventType] = append(registered[eventType].([]func(context.Context, DomainEvent) error), handler)
}

func Dispatch(ctx context.Context, events []DomainEvent) error {
	for _, e := range events {
		eventType := e.EventKey()
		if _, ok := registered[eventType]; !ok {
			continue
		}
		l := registered[eventType].([]func(context.Context, DomainEvent) error)
		for _, f := range l {
			//if err := f(ctx, e); err != nil {
			//	return err
			//}

			//goland:noinspection ALL
			go f(ctx, e)
		}
	}
	return nil
}
