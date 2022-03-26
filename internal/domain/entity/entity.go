package entity

import "go-ddd-blog/internal/domain/event"

type Eventor interface {
	Events() []event.DomainEvent
}

type Entity struct {
	id     int64
	events []event.DomainEvent
}

func NewEntity(id int64) Entity {
	return Entity{id: id}
}

func (e Entity) ID() int64 {
	return e.id
}

func (e *Entity) AddEvent(event event.DomainEvent) {
	e.events = append(e.events, event)
}

func (e Entity) Events() []event.DomainEvent {
	return e.events
}
