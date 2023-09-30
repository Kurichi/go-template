package domain

import (
	"context"

	"github.com/Kurichi/go-template/modules/todo/domain/todo"
)

type TodoRepository interface {
	GetList(ctx context.Context, filter *todo.Todo) ([]*todo.Todo, error)
	Store(context.Context, *todo.Todo) error
	GetByID(context.Context, todo.TodoID) (*todo.Todo, error)
	Update(context.Context, *todo.Todo) error
	Delete(context.Context, todo.TodoID) error
}
