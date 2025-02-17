package handler

import (
	"net/http"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/constants"
	"github.com/gorilla/mux"
)

func (h *TodoHandler) RegisterHandlers(r *mux.Router) {
	todoRouter := r.PathPrefix(constants.TodosPath).Subrouter()

	todoRouter.HandleFunc("", h.ListTodo).Methods(http.MethodGet, http.MethodOptions)
	todoRouter.HandleFunc("/{id}", h.GetTodo).Methods(http.MethodGet, http.MethodOptions)
	todoRouter.HandleFunc("", h.CreateTodo).Methods(http.MethodPost, http.MethodOptions)
	todoRouter.HandleFunc("/{id}", h.UpdateTodo).Methods(http.MethodPut, http.MethodOptions)
	todoRouter.HandleFunc("/{id}", h.DeleteTodo).Methods(http.MethodDelete, http.MethodOptions)
}