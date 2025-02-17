package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	persistence_gorm "github.com/YukiOnishi1129/react-output-crud-api/backend/internal/infrastructure/persistence/gorm"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/interfaces/handler"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/database"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Printf("Start server")
	db, err := database.InitConnectDB()
	if err != nil {
		log.Fatalf("Error connect to database: %v", err)
		return
	}
	r := mux.NewRouter()
	todoRepository := persistence_gorm.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUseCase(todoRepository)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	todoHandler.RegisterHandlers(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")}, // フロントエンドのオリジン
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)



	log.Printf("Server started at http://localhost:%s", os.Getenv("BACKEND_CONTAINER_POST"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("BACKEND_CONTAINER_POST")), handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
}