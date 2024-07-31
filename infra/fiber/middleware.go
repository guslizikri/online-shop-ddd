package infrafiber

import (
	"log"
	"online-shop-ddd/infra/response"
	"online-shop-ddd/internal/config"
	"online-shop-ddd/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("token invalid")
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		token := bearer[1]

		publicId, role, err := utility.ValidateToken(token, config.Cfg.App.Encryption.JwtSecret)
		if err != nil {
			log.Println(err.Error())
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		c.Locals("ROLE", role)
		c.Locals("PUBLIC_ID", publicId)

		return c.Next()
	}
}
