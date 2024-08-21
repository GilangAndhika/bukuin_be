package controllers

import (
	"github.com/gofiber/fiber/v2"
	repo "github.com/gryzlegrizz/bukuin_be/repository"
	"gorm.io/gorm"
	"net/http"
	"github.com/gryzlegrizz/bukuin_be/models"
)

func GetAllBooks(c *fiber.Ctx) error {
	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// memanggil fungsi repo untuk mendapatkan semua data buku
	books, err := repo.GetAllBooks(db)
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
		"data" : books,
	}
	return c.Status(http.StatusOK).JSON(response)
}

func GetBookByID(c *fiber.Ctx) error {
	// mendapatkan parameter ID dari URL
	id := c.Params("id_book")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID is not found"})
	}

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk mendapatkan task berdasarkan ID
	books, err := repo.GetBookByID(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} 
	
	// jika tidak ada kesalahan, mengembalikan data book sebagai respons JSON
	return c.JSON(fiber.Map{
		"code":http.StatusOK,
		"success":true,
		"status": "success",
		"data" : books,
	})
}

func CreateBook(c *fiber.Ctx) error {
	// mendeklarasikan variabel untuk menyimpan data book dari body request
	var books models.Books

	// mem-parsing body request ke dalam variabel book
	if err := c.BodyParser(&books); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to process request"})
	}

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// memanggil fungsi repo untuk menambahkan data buku ke database
	if err := repo.CreateBook(db, books); err != nil {
		// jika terjadi kesalahan saat menambahkan data buku, mengembalikan respon err
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// jika tidak ada kesalahan, mengembalikan respon sukses
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"code":http.StatusCreated,
		"success":true,
		"status": "success",
		"message": "Book created successfully",
		"data": books,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	// mendapatkan parameter ID dari URL
	id := c.Params("id_user")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID is not found"})
	}

	// mendeklarasikan variabel untuk menyimpan data books yang diperbarui dari body request
	var UpdatedBooks models.Books

	// mem-parsing body request ke dalam variabel UpdatedBooks
	if err := c.BodyParser(&UpdatedBooks); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to process request"})
	}

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk memperbarui data buku di dalam database
	if err := repo.UpdateBook(db, id, UpdatedBooks); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update book"})
	}

	// jika tidak ada kesalahan, mengembalikan respon sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":http.StatusOK,
		"success":true,
		"status": "success",
		"message": "Book updated successfully",
	})
}

func DeleteBook(c *fiber.Ctx) error {
	// mendapatkan parameter ID dari URL
	id := c.Params("id_user")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID is not found"})
	}

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk menghapus data buku dari database
	if err := repo.DeleteBook(db, id); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	// jika tidak ada kesalahan, mengembalikan respon sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":http.StatusOK,
		"success":true,
		"status": "success",
		"message": "Book deleted successfully",
		"deleted_id": id,
	})
}