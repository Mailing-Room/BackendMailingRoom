package user_controller

import (
	"backendmailingroom/config/middleware"
	"backendmailingroom/model"
	"backendmailingroom/pkg/password"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (u *UserHandler) InputUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "gagal membaca body request: " + err.Error(),
		})
	}

	// Validasi input
	if user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email tidak boleh kosong",
		})
	}

	if user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password tidak boleh kosong",
		})
	}

	passwordHash, err := password.HashingPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "gagal melakukan hash password: " + err.Error(),
		})
	}

	user.Password = passwordHash
	_, err = u.user.InputUser(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "gagal menyimpan user: " + err.Error(),
		})
	}
	return c.JSON(fiber.Map{"status": "success", "user": user})
}

func (u *UserHandler) Login(c *fiber.Ctx) error {
	req := new(model.Login)
	if err := c.BodyParser(req); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}
	user, err := u.user.GetUserForLogin(c.Context(), req.Email)
	if err != nil {
		return fmt.Errorf("failed to login user: %w", err)
	}
	passwordValid := password.CheckPassword(user.Password, req.Password)
	log.Println(passwordValid)
	if !passwordValid {
		return fmt.Errorf("invalid password")
	}

	token, err := middleware.EncodeToken(user.ID)
	if err != nil {
		return fmt.Errorf("failed to generate token: %w", err)
	}
	return c.JSON(fiber.Map{"status": "success", "user": user, "token": token})
}

func (u *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := u.user.GetAllUsers(c.Context())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	return c.JSON(fiber.Map{"status": "success", "users": users})
}

func (u *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := u.user.GetUserByID(c.Context(), id)
	if err != nil {
		return fmt.Errorf("failed to get user by ID: %w", err)
	}
	return c.JSON(fiber.Map{"status": "success", "user": user})
}

func (u *UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := u.user.GetUserByEmail(c.Context(), email)
	if err != nil {
		return fmt.Errorf("failed to get user by email: %w", err)
	}
	return c.JSON(fiber.Map{"status": "success", "user": user})
}

func (u *UserHandler) DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Validasi parameter ID
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "id parameter is required",
		})
	}

	// Panggil repository
	user, err := u.user.DeleteUserByID(c.Context(), id)
	if err != nil {
		// Tangani user tidak ditemukan
		if strings.Contains(err.Error(), "tidak ditemukan") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		// Tangani error umum
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("failed to delete user: %v", err),
		})
	}

	// Jika berhasil dihapus
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "user deleted successfully",
		"user":    user,
	})
}

func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "id tidak boleh kosong",
		})
	}

	var updatedData model.User
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "gagal membaca body request: " + err.Error(),
		})
	}

	// 🧩 Jika field password tidak kosong, hash dulu
	if updatedData.Password != "" {
		passwordHash, err := password.HashingPassword(updatedData.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "gagal melakukan hash password: " + err.Error(),
			})
		}
		updatedData.Password = passwordHash
	}

	// 🔄 Lanjutkan proses update
	user, err := u.user.UpdateUser(c.Context(), id, updatedData)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"user":   user,
	})
}
