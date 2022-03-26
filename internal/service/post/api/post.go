package api

import (
	"go-ddd-blog/internal/service/post/api/command"
	"go-ddd-blog/internal/service/post/api/query"
)

type Application struct {
	command.Command
	query.Query
}

func NewApplication(command command.Command, query query.Query) *Application {
	return &Application{Command: command, Query: query}
}
