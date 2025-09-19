package main

import (
	"log"
	"os"

	"go-tenders/api"
	"go-tenders/config"
	"go-tenders/server"
	"go-tenders/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig() // Имплементация LoadConfig должна возвращать структуру с Host и Port

	// Инициализируем хранилище (Postgres)
	dbURL := os.Getenv("DATABASE_URL")
	store, err := storage.NewPostgresStorage(dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	// Создаем сервер api.ServerInterface
	apiServer := server.NewServer(store, nil, cfg)

	// Создаем Echo
	e := echo.New()

	// Встроенный логгер Echo для HTTP запросов
	e.Use(middleware.Logger())

	// Восстановление после паники и возврат 500
	e.Use(middleware.Recover())

	// Регистрируем API маршруты
	api.RegisterHandlers(e, apiServer)

	// Формируем адрес из конфигурации
	addr := cfg.Host + ":" + cfg.Port
	log.Printf("Starting server at %s", addr)

	// Запускаем сервер
	if err := e.Start(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
