package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/muasx88/library-api/config/middleware"
	"github.com/muasx88/library-api/internal/config"
)

func Init(router *fiber.App, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	bookRouter := router.Group("books")
	bookRouter.Get("", handler.GetAllBook)
	bookRouter.Get(":id", handler.GetBook)
	bookRouter.Post("add", middleware.CheckAuth(), middleware.CheckRoles([]string{string(config.ROLE_Admin)}), handler.AddBook)
	// bookRouter.Get(":id", handler.AddBook)
}
