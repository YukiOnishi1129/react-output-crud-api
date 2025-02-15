package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/usecase/todo"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TodoHandler struct {
	todoUseCase todo.TodoUseCase
}

func NewTodoHandler(todoUseCase todo.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

// GetTodos は、Todoリストを取得するハンドラーです
func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// クエリパラメータの取得
	input := &todo.GetTodosInput{
		Limit:  20, // デフォルト値
		Offset: 0,
	}

	if limit := r.URL.Query().Get("limit"); limit != "" {
		if n, err := strconv.Atoi(limit); err == nil {
			input.Limit = n
		}
	}

	if offset := r.URL.Query().Get("offset"); offset != "" {
		if n, err := strconv.Atoi(offset); err == nil {
			input.Offset = n
		}
	}

	// ユーザーIDの取得（認証ミドルウェアから）
	userID, ok := ctx.Value("userID").(uuid.UUID)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	input.UserID = userID

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.todoUseCase.GetTodos(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// GetTodo は、指定されたTodoを取得するハンドラーです
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	userID, ok := ctx.Value("userID").(uuid.UUID)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	input := &todo.GetTodoInput{
		ID:     todoID,
		UserID: userID,
	}

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.todoUseCase.GetTodo(ctx, todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// CreateTodo は、新しいTodoを作成するハンドラーです
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input todo.CreateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := ctx.Value("userID").(uuid.UUID)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	input.UserID = userID

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.todoUseCase.CreateTodo(ctx, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// UpdateTodo は、指定されたTodoを更新するハンドラーです
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	var input todo.UpdateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input.ID = todoID

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.todoUseCase.UpdateTodo(ctx, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// DeleteTodo は、指定されたTodoを削除するハンドラーです
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	userID, ok := ctx.Value("userID").(uuid.UUID)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	input := &todo.DeleteTodoInput{
		ID:     todoID,
		UserID: userID,
	}

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.todoUseCase.DeleteTodo(ctx, todoID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
} 