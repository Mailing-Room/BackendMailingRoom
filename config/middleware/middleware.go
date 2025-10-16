package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
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
	userID, err := DecodeToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or expired token",
			"error":   err.Error(),
		})
	}

	// Simpan userID ke context agar bisa diakses di handler berikutnya
	c.Locals("user_id", userID)

	return c.Next()
}
