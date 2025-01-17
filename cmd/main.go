package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jamolkhon5/darkmode/internal/auth"
	"github.com/Jamolkhon5/darkmode/internal/config"
	"github.com/Jamolkhon5/darkmode/internal/handler"
	"github.com/Jamolkhon5/darkmode/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.NewConfig(".env.local")
	if err != nil {
		log.Fatal(err)
	}

	// Подключение к базе данных
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPassword, cfg.PgName)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS themes (
            id SERIAL PRIMARY KEY,
            user_id VARCHAR(255) UNIQUE NOT NULL,
            theme VARCHAR(50) NOT NULL DEFAULT 'light',
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	// Подключение к сервису auth
	authConfig, err := auth.NewConfig(".auth.env")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(authConfig.AuthAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	auth.InitClient(conn)

	// Инициализация репозитория и хендлера
	repo := repository.NewRepository(db)
	handler := handler.NewHandler(repo)

	// Настройка роутера
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Put("/v1/users/me/theme", handler.UpdateTheme)

	// Запуск сервера
	log.Printf("Server is running on port 5640")
	log.Fatal(http.ListenAndServe(":5640", r))
}
