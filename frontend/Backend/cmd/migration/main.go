package migration

import (
	"flag"
	"log"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/config/database"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/pointer"
	"github.com/joho/godotenv"
)

func main() {
	// フラグを定義
	isSeed := flag.Bool("seed", false, "シードデータを投入する")
	isReset := flag.Bool("reset", false, "テーブルを再作成する")
	flag.Parse()

	log.Printf("Start migration")
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env.local file: %v", err)
		return
	}

	db, err := database.InitConnectDB()
	if err != nil {
		log.Fatalf("Error connect to database: %v", err)
		return
	}

	// リセットフラグが立っている場合、テーブルを削除
	if *isReset {
		log.Printf("Dropping tables...")
		err := db.Migrator().DropTable(&domain.Todo{})
		if err != nil {
			log.Fatalf("Error dropping tables: %v", err)
			return
		}
	}

	// UUID拡張機能を有効化
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// マイグレーションを実行
	db.AutoMigrate(&domain.Todo{})
	log.Printf("Migration completed")

	// シードフラグが立っている場合のみシードデータを投入
	if *isSeed {
		log.Printf("Start seeding")
		insertTodoList := []*domain.Todo{
			{
				Title:   "title1",
				Content: pointer.String("content1"),
			},
			{
				Title:   "title2",
				Content: pointer.String("content2"),
			},
		}

		result := db.Create(insertTodoList)
		if result.Error != nil {
			log.Fatalf("Error inserting todos: %v", result.Error)
			return
		}

		log.Printf("Successfully inserted %d todos", len(insertTodoList))
	}
} 