package todoserver

import (
	"context"

	todos "github.com/slaskis/tweact/pkg/todos/v1"
)

func (s *TodoServer) ListTodos(ctx context.Context, in *todos.ListTodosRequest) (*todos.ListTodoResponse, error) {
	return &todos.ListTodoResponse{
		Todos: s.todos,
	}, nil
}
