package controllers

import (
	"net/http"

	"github.com/GilangAndhika/bukuin_be/models"
	repo "github.com/GilangAndhika/bukuin_be/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllBooks(c *fiber.Ctx) error {
	// cek autentikasi token header
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header is not found")
	}

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// memanggil fungsi repo untuk mendapatkan semua data buku
	books, err := repo.GetAllBooks(db)
	if err != nil {
		// jika terjadi kesalahan saat mengambil data buku, mengembalikan respon err
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get books data"})
	}

	// jika tidak ada data buku yang ditemukan, mengembalikan respon not found
	if len(books) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"success": false,
			"status":  "error",
			"message": "No books found",
			"data":    nil})
	}

	// jika tidak ada kesalahan, mengembalikan data books sebagai respons JSON
	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    books,
	}
	return c.Status(http.StatusOK).JSON(response)
}

func GetBookByID(c *fiber.Ctx) error {
	// cek autentikasi token header
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header is not found")
	}

	// mendapatkan parameter ID dari URL
	id := c.Query("id_book")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID book cannot be empty"})
	}

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk mendapatkan task berdasarkan ID
	books, err := repo.GetBookByID(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Data not found"})
	}

	// jika tidak ada kesalahan, mengembalikan data book sebagai respons JSON
	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    books,
	})
}

func GetBookByIdUser(c *fiber.Ctx) error {
	// cek autentikasi token header
	tokenStr := c.Get("login")
	if tokenStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header is not found")
	}

	// parse token untuk mendapatkan id user
	token, err := jwt.ParseWithClaims(tokenStr, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	IdUser := claims.IdUser

	db := c.Locals("db").(*gorm.DB)

	// memanggil fungsi repo untuk mendapatkan data buku berdasarkan ID user
	books, err := repo.GetBookByIdUser(db, int(IdUser))
	if err != nil {
		// jika terjadi kesalahan saat mengambil data buku, mengembalikan respon err
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Data not found"})
	}

	// jika tidak ada kesalahan, mengembalikan data books sebagai respons JSON
	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    books,
	})
}

func CreateBook(c *fiber.Ctx) error {
	// cek autentikasi token header
	tokenStr := c.Get("login")
	if tokenStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header is not found")
	}

	// parse token untuk mendapatkan id user
	token, err := jwt.ParseWithClaims(tokenStr, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	IdUser := claims.IdUser

	// mendeklarasikan variabel untuk menyimpan data book dari body request
	var books models.Books

	// mem-parsing body request ke dalam variabel book
	if err := c.BodyParser(&books); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to process request"})
	}

	books.IdUser = int(IdUser)

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// memanggil fungsi repo untuk menambahkan data buku ke database
	if err := repo.CreateBook(db, &books); err != nil {
		// jika terjadi kesalahan saat menambahkan data buku, mengembalikan respon err
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// jika tidak ada kesalahan, mengembalikan respon sukses
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"code":    http.StatusCreated,
		"success": true,
		"status":  "success",
		"message": "Book created successfully",
		"data":    books,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	// cek autentikasi token header
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(http.StatusBadRequest, "Header is not found")
	}

	// mendapatkan parameter ID dari URL
	id := c.Query("id_book")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID book is not found"})
	}

	// mendeklarasikan variabel untuk menyimpan data books yang diperbarui dari body request
	var UpdatedBooks models.Books

	// mem-parsing body request ke dalam variabel UpdatedBooks
	if err := c.BodyParser(&UpdatedBooks); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to process request"})
	}

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk cek data buku berdasarkan ID
	_, err := repo.GetBookByID(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "ID book is not found"})
	}

	// Memanggil fungsi repo untuk memperbarui data buku di dalam database
	if err := repo.UpdateBook(db, id, UpdatedBooks); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update book"})
	}

	// jika tidak ada kesalahan, mengembalikan respon sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"message": "Book updated successfully",
	})
}

func DeleteBook(c *fiber.Ctx) error {
	// cek autentikasi token header
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header is not found")
	}

	// mendapatkan parameter ID dari URL
	id := c.Query("id_book")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID cannot be empty"})
	}

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk cek data buku berdasarkan ID
	_, err := repo.GetBookByID(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "ID book is not found"})
	}

	// Memanggil fungsi repo untuk menghapus data buku dari database
	if err := repo.DeleteBook(db, id); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	// jika tidak ada kesalahan, mengembalikan respon sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":       http.StatusOK,
		"success":    true,
		"status":     "success",
		"message":    "Book deleted successfully",
		"deleted_id": id,
	})
}
