package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gryzlegrizz/bukuin_be/models"
	repo "github.com/gryzlegrizz/bukuin_be/repositories"
	"gorm.io/gorm"
	"net/http"
)

func GetAllBooks(c *fiber.Ctx) error {
	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// memanggil fungsi repo untuk mendapatkan semua data buku
	books := repo.GetAllBooks(db)
	if err != nil {
		// jika terjadi kesalahan saat mengambil data buku, mengembalikan respon err
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// jika tidak ada data buku yang ditemukan, mengembalikan respon not found
	if len(books) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":http.StatusNotFound,
			"success":false,
			"status": "error",
			"message": "No books found",
			"data": nil,})
	}

	// jika tidak ada kesalahan, mengembalikan data books sebagai respons JSON
	response := fiber.Map{
		"code":http.StatusOK,
		"success":true,
		"status": "success",
		data : books,
	}
	return c.Status(http.StatusOK).JSON(response)
}