package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)
	// karena fiber.router adalah interface, jadi gaperlu di kasih pointer.
	// biar bisa memakai router.group
	userRoute := router.Group("users")
	{
		userRoute.Post("register", handler.register)
		userRoute.Post("login", handler.login)
	}
}
