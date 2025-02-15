.PHONY: migrate seed db-reset

# マイグレーションを実行
migrate:
	go run cmd/migration/main.go

# シードデータを投入
seed:
	go run cmd/migration/main.go -seed

# DBをリセットしてマイグレーションとシードを実行
db-reset:
	go run cmd/migration/main.go -reset -seed 