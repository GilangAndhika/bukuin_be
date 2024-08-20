package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gryzlegrizz/bukuin_be/controllers"
)

func SetupBooksRoute(app *fiber.App) {
	// Mengatur route untuk mengambil semua data buku
	app.Get("/books", controllers.GetAllBooks)
	// Mengatur route untuk mengambil data buku berdasarkan ID
	app.Get("/books/get/:id", controllers.GetBookByID)
	// Mengatur route untuk menambahkan data buku
	app.Post("/books/create", controllers.CreateBook)
	// Mengatur route untuk memperbarui data buku
	app.Put("/books/update/:id", controllers.UpdateBook)
	// Mengatur route untuk menghapus data buku
	app.Delete("/books/delete/:id", controllers.DeleteBook)
}