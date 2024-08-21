package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gryzlegrizz/bukuin_be/models"
	repo "github.com/gryzlegrizz/bukuin_be/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.Users

	// menghubungkan ke database
	db := c.Locals("db").(*gorm.DB)

	// mengambil data dari body request
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body invalid",
		})
	}

	// menyimpan user ke database menggunakan repository
	if err := repo.CreateUser(db, &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var user models.Users

	// menghubungkan ke database
	db := c.Locals("db").(*gorm.DB)

	// mengambil data dari body request
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body invalid",
		})
	}

	// mencari user berdasarkan username
	userData, err := repo.GetUserByUsername(db, user.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username or password is incorrect",
		})
	}

	// verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username or password is incorrect",
		})
	}

	// membuat token JWT menggunakan repo
	token, err := repo.CreateToken(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func GetUser(c *fiber.Ctx) error {
	// mendapatkan token JWT dari header Authorization
	tokenString := c.Get("login")
	if tokenString == "" {
		return fiber.NewError(fiber.StatusNotFound, "Token not found in header")
	}

	// parse token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return err
	}

	// memeriksa apakah token valid
	if !token.Valid {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid token")
	}

	// mengekstrak klaim dari token
	claims := token.Claims.(jwt.MapClaims)
	IdUser := uint(claims["id_user"].(float64))
	db := c.Locals("db").(*gorm.DB)

	// mencari user berdasarkan ID menggunakan repo
	userData, err := repo.GetUserByID(db, IdUser)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"user": userData,
	})
}