package main

import (
	"log"
	"os"

	"go-tenders/server"
	"go-tenders/storage"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Загрузить переменные из .env в окружение
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Считать параметры из окружения
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Подключиться к базе
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Создать хранилище
	storageInstance := storage.NewPostgresStorage(db)

	// Создать логгер (пример)
	// loggerInstance := logger.NewLogger()

	// Конфиг сервера
	config := &server.Config{
		Port: 8080,
	}

	// Создать сервер с внедрением storage и логгера
	srv := server.NewServer(storageInstance, loggerInstance, config)

	// Запустить HTTP сервер (например, Echo)
	srv.Start()
}
