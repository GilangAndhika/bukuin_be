package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gryzlegrizz/bukuin_be/config"
	"github.com/gryzlegrizz/bukuin_be/routes"
)

func main() {
	// Membuat instance aplikasi fiber
	app := fiber.New()

	// Membuat koneksi database
	db := config.CreateDBConnection()

	app.Use(logger.New(logger.Config{
		Format: "${status - ${method} ${path}\n",
	}))

	// Menyimpan koneksi database dalam context fiber
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Mengatur route untuk buku
	routes.SetupBooksRoute(app)

	// Menjalankan aplikasi pada port 3000
	app.Listen(":3000")
}
