func NewTodoListOutput(todos *dto.TodoListOutput, total int64) *TodoListOutput {
	outputs := make([]TodoOutput, len(todos.Todos))
	for i, todo := range todos.Todos {
		outputs[i] = *NewTodoOutput(&todo)  // todo is already dto.TodoOutput
	}
	return &TodoListOutput{
		Todos: outputs,
		Total: total,
	}
} 