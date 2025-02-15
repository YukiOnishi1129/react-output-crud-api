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
	todoHandler.RegisterHandlers(r)

	r.Use(mux.CORSMethodMiddleware(r))


	log.Printf("Server started at http://localhost:%s", os.Getenv("BACKEND_CONTAINER_POST"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("BACKEND_CONTAINER_POST")), r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
}