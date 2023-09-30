//go:generate stringer -type=TodoStatus -output=value_string.go

package todo

import "github.com/google/uuid"

type (
	TodoID     uuid.UUID
	TodoStatus uint8
)

// Enum
const (
	TodoStatusPending TodoStatus = iota + 1
	TodoStatusInProgress
	TodoStatusDone
)

// New functions
func NewTodoID() TodoID {
	return TodoID(uuid.New())
}

// String functions
func (t TodoID) String() string {
	return uuid.UUID(t).String()
}
