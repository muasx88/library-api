package transaction

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

	transactionRouter := router.Group("transaction", middleware.CheckAuth())
	transactionRouter.Post("borrow", middleware.CheckRoles([]string{string(config.ROLE_User)}), handler.BorrowBook)
	transactionRouter.Post("return", middleware.CheckRoles([]string{string(config.ROLE_User)}), handler.ReturnBook)
	transactionRouter.Get(
		"histories/:user_id",
		middleware.CheckRoles([]string{string(config.ROLE_Admin)}),
		handler.GetUserHistoryTransaction,
	)
}
