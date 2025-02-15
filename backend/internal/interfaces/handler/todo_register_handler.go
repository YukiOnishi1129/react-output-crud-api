package handler

import (
	"net/http"
	"os"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/constants"
	"github.com/gorilla/mux"
)

func (h *TodoHandler) RegisterHandlers(r *mux.Router) {
	todoRouter := r.PathPrefix(constants.TodosPath).Subrouter()

	todoRouter.HandleFunc("", h.ListTodo).Methods("GET")
	todoRouter.HandleFunc("/{id}", h.GetTodo).Methods("GET")
	todoRouter.HandleFunc("", h.CreateTodo).Methods("POST")
	todoRouter.HandleFunc("/{id}", h.UpdateTodo).Methods("PUT")
	todoRouter.HandleFunc("/{id}", h.DeleteTodo).Methods("DELETE")
	todoRouter.HandleFunc("/{id}", optionsHandler).Methods("OPTIONS")
}


func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusNoContent)
}