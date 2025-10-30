package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	pgxrepo "yourapp/repository/pgx"
	fiberhttp "yourapp/repository/transport/http/fiber"
)

// точка входа (запуск Fiber, DI)
func connectDB() (*pgxpool.Pool, error) {
	dsn := os.Getenv("DATABASE_URL")
	return pgxpool.New(context.Background(), dsn)
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	bookRepo := pgxrepo.NewBookRepo(db)
	h := &fiberhttp.Handler{Books: bookRepo}
	app := fiber.New()
	fiberhttp.RegisterRoutes(app, h)
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
