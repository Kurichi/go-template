// Code generated by "stringer -type=TodoStatus -output=value_string.go"; DO NOT EDIT.

package todo

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TodoStatusPending-1]
	_ = x[TodoStatusInProgress-2]
	_ = x[TodoStatusDone-3]
}

const _TodoStatus_name = "TodoStatusPendingTodoStatusInProgressTodoStatusDone"

var _TodoStatus_index = [...]uint8{0, 17, 37, 51}

func (i TodoStatus) String() string {
	i -= 1
	if i >= TodoStatus(len(_TodoStatus_index)-1) {
		return "TodoStatus(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TodoStatus_name[_TodoStatus_index[i]:_TodoStatus_index[i+1]]
}