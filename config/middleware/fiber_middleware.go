package middleware

import (
	"fmt"
	"log"
	"strings"

	cfg "github.com/muasx88/library-api/internal/config"
	"github.com/muasx88/library-api/internal/jwt_helper"
	"github.com/muasx88/library-api/internal/response"

	// "github.com/NooBeeID/go-logging/logger"
	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		// Get query params and path params
		queryArgs := make(map[string][]string)
		c.Context().QueryArgs().VisitAll(func(key, value []byte) {
			queryArgs[string(key)] = append(queryArgs[string(key)], string(value))
		})

		// Log request details
		log.Println("Request Log", map[string]interface{}{
			"IP":     c.IP(),
			"METHOD": c.Method(),
			"PATH":   c.Path(),
			"Query":  c.Context().QueryArgs(),
		})

		return err
	}
}

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return response.MsgWithCode(c, fiber.StatusUnauthorized, "token required")
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("token invalid")
			return response.MsgWithCode(c, fiber.StatusUnauthorized, "invalid token")
		}

		token := bearer[1]

		id, role, err := jwt_helper.ValidateToken(token, cfg.Config.App.Encryption.JWTSecret)
		if err != nil {
			log.Println(err.Error())
			return response.MsgWithCode(c, fiber.StatusUnauthorized, "invalid token")
		}

		c.Locals("role", role)
		c.Locals("id", id)

		return c.Next()
	}
}

func CheckRoles(authorizedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := fmt.Sprintf("%v", c.Locals("role"))

		isExists := false
		for _, authorizedRole := range authorizedRoles {
			if role == authorizedRole {
				isExists = true
				break
			}
		}

		if !isExists {
			return response.MsgWithCode(c, fiber.StatusForbidden, "forbidden")
		}

		return c.Next()
	}
}
