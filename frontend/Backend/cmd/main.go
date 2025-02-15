package main

import (
	"gorm.io/gorm"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/usecase"
)

func main() {
	db := // ... DB初期化

	todoRepo := gorm.NewTodoRepository(db)
	todoUseCase := usecase.NewTodoUseCase(todoRepo)

	// ... HTTPハンドラーの設定など
} 