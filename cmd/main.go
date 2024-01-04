package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/davecgh/go-spew/spew"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/muasx88/library-api/config/middleware"
	"github.com/muasx88/library-api/domains/auth"
	"github.com/muasx88/library-api/domains/book"
	"github.com/muasx88/library-api/domains/transaction"
	"github.com/muasx88/library-api/internal/config"
	"github.com/muasx88/library-api/internal/database"
)

func init() {
	err := config.LoadConfig("config")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	spew.Dump(config.Config)

	db, err := database.ConnectDB(context.Background())
	if err != nil {
		log.Fatalf("error connect db %s\n", err.Error())
	}

	app := fiber.New(fiber.Config{
		AppName:     config.Config.App.Name,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(middleware.LoggingMiddleware())

	auth.Init(app, db)
	book.Init(app, db)
	transaction.Init(app, db)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Fatalf("Server not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := app.Listen(":" + config.Config.App.Port); err != nil {
		log.Fatalf("Server not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
