package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gryzlegrizz/bukuin_be/controllers"
)

func SetupBooksRoute(app *fiber.App) {
	// Mengatur route untuk mengambil semua data buku
	app.Get("/books", controllers.GetAllBooks)
}