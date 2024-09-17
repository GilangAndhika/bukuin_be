package controllers

import (
	"net/http"

	"github.com/GilangAndhika/bukuin_be/models"
	repo "github.com/GilangAndhika/bukuin_be/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllRoles(c *fiber.Ctx) error {
	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// memanggil fungsi repo untuk mendapatkan semua data role
	roles, err := repo.GetAllRoles(db)
	if err != nil {
		// jika terjadi kesalahan saat mengambil data role, mengembalikan respon err
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// jika tidak ada data role yang ditemukan, mengembalikan respon not found
	if len(roles) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"success": false,
			"status":  "error",
			"message": "No roles found",
			"data":    nil})
	}

	// jika tidak ada kesalahan, mengembalikan data roles sebagai respons JSON
	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    roles,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func GetRoleByID(c *fiber.Ctx) error {
	// mendapatkan parameter ID dari URL
	id := c.Params("id_role")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID role is not found"})
	}

	// mendapatkan koneksi database dari cotext fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk mendapatkan role berdasarkan ID
	role, err := repo.GetRoleByID(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// jika tidak ada kesalahan, mengembalikan data role sebagai respons JSON
	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    role,
	})
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Roles

	// menghubungkan ke database
	db := c.Locals("db").(*gorm.DB)

	// mengambil data dari body request
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body invalid",
		})
	}

	// memanggil fungsi repo untuk menambahkan data role
	if err := repo.CreateRole(db, role); err != nil {
		// jika terjadi kesalahan saat menambahkan data role, mengembalikan respon err
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create role"})
	}

	// jika tidak ada kesalahan, mengembalikan respon sukses
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"code":    http.StatusCreated,
		"success": true,
		"status":  "success",
		"message": "Role created successfully",
		"data":    role,
	})
}

func UpdateRole(c *fiber.Ctx) error {
	// mendapatkan parameter ID dari URL
	id := c.Params("id_role")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID role is not found"})
	}

	var updatedRole models.Roles

	// menghubungkan ke database
	db := c.Locals("db").(*gorm.DB)

	// mengambil data dari body request
	if err := c.BodyParser(&updatedRole); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body invalid",
		})
	}

	// memanggil fungsi repo untuk memperbarui data role
	if err := repo.UpdateRole(db, id, updatedRole); err != nil {
		// jika terjadi kesalahan saat memperbarui data role, mengembalikan respon err
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update role"})
	}

	// jika tidak ada kesalahan, mengembalikan respon sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"message": "Role updated successfully",
	})
}

func DeleteRole(c *fiber.Ctx) error {
	// mendapatkan parameter ID dari URL
	id := c.Params("id_role")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID role is not found"})
	}

	// menghubungkan ke database
	db := c.Locals("db").(*gorm.DB)

	// memanggil fungsi repo untuk menghapus data role
	if err := repo.DeleteRole(db, id); err != nil {
		// jika terjadi kesalahan saat menghapus data role, mengembalikan respon err
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete role"})
	}

	// jika tidak ada kesalahan, mengembalikan respon sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":       http.StatusOK,
		"success":    true,
		"status":     "success",
		"message":    "Role deleted successfully",
		"deleted_id": id,
	})
}
