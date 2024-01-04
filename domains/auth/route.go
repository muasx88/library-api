package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router *fiber.App, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	authRouter := router.Group("auth")
	authRouter.Post("register", handler.register)
	authRouter.Post("login", handler.login)
}
