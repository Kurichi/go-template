package todo

type Todo struct {
	ID       TodoID
	Title    string
	AuthorID string
	Status   TodoStatus
}

func NewTodo(title string) *Todo {
	return &Todo{
		ID:     NewTodoID(),
		Title:  title,
		Status: TodoStatusPending,
	}
}
