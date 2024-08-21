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
	// mendapatkan data user yang sedang login melalui JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*models.JWTClaims)
	db := c.Locals("db").(*gorm.DB)

	// mencari user berdasarkan ID menggunakan repo
	userData, err := repo.GetUserByID(db, claims.IdUser)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"data":    userData,
	})
}

func Authenticate(c *fiber.Ctx) error {
	// mendapatkan token dari header authorization
	authHeader := c.Get("Authorization")
	token := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer" {
		token = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong authorization header",
		})
	}

	// verifikasi token
	claims := new(models.JWTClaims)
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token signature",
			})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Authentication token failed",
		})
	}

	if !tkn.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// menyimpan data user ke context fiber
	c.Locals("user", tkn)

	return c.Next()
}

func LogoutUser(c *fiber.Ctx) error {
	// menghapus token dari authorization header
	c.Set("Authorization", "")

	return c.JSON(fiber.Map{
		"message": "Logout success",
	})
}