package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing Authorization header",
			})
		}

		// Format header: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token format",
			})
		}

		// Decode dan verifikasi token
		userID, roleID, err := DecodeToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired token",
				"error":   err.Error(),
			})
		}

		// Check if role is in requiredRoles
		allowed := false
		for _, role := range requiredRoles {
			if roleID == role {
				allowed = true
				break
			}
		}

		if !allowed {
			return fiber.NewError(fiber.StatusForbidden, "Access denied for this role")
		}

		// Simpan userID ke context agar bisa diakses di handler berikutnya
		c.Locals("user_id", userID)
		c.Locals("role_id", roleID)

		return c.Next()
	}
}
