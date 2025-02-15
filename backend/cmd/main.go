package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	gorm_persistence "github.com/YukiOnishi1129/react-output-crud-api/backend/internal/infrastructure/persistence/gorm"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/interfaces/handler"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/database"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/usecase"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Start server")
	db, err := database.InitConnectDB()
	if err != nil {
		log.Fatalf("Error connect to database: %v", err)
		return
	}
	r := mux.NewRouter()
	todoRepository := gorm_persistence.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUseCase(todoRepository)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	r.HandleFunc("/todos", todoHandler.ListTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods("GET")
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	log.Printf("Server started at http://localhost:%s", os.Getenv("BACKEND_CONTAINER_POST"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("BACKEND_CONTAINER_POST")), r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
}