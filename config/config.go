package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config содержит параметры конфигурации приложения
type Config struct {
	ServerHost  string `envconfig:"SERVER_HOST" default:"localhost"` // хост сервера
	ServerPort  int    `envconfig:"SERVER_PORT" default:"8080"`      // порт сервера
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`    // URL подключения к базе данных
}

// LoadConfig загружает конфигурацию из .env и переменных окружения
func LoadConfig() (*Config, error) {
	// Подгружаем переменные окружения из файла .env (если он есть)
	err := godotenv.Load()
	if err != nil {
		log.Println("warning: .env file not found, relying on environment variables")
	}

	// Создаем структуру конфигурации
	var cfg Config

	// Заполняем поля структуры значениями из переменных окружения
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
