package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gryzlegrizz/bukuin_be/controllers"
)

func SetupBooksRoute(app *fiber.App) {
	// Mengatur route untuk Login dan Register
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)

	// Mengatur route untuk authentikasi
	app.Get("/auth", controllers.GetUser)

	// Mengatur route untuk data buku
	app.Get("/books", controllers.GetAllBooks) 						// Mengatur route untuk mengambil semua data buku
	app.Get("/books/get", controllers.GetBookByID)					// Mengatur route untuk mengambil data buku berdasarkan ID
	app.Post("/books/create", controllers.CreateBook) 				// Mengatur route untuk menambahkan data buku
	app.Put("/books/update", controllers.UpdateBook) 				// Mengatur route untuk memperbarui data buku
	app.Delete("/books/delete", controllers.DeleteBook) 			// Mengatur route untuk menghapus data buku

	// Mengatur route untuk data role
	app.Get("/roles", controllers.GetAllRoles) 						// Mengatur route untuk mengambil semua data role
	app.Get("/roles/get/:id_role", controllers.GetRoleByID)			// Mengatur route untuk mengambil data role berdasarkan ID
	app.Post("/roles/create", controllers.CreateRole) 				// Mengatur route untuk menambahkan data role
	app.Put("/roles/update/:id_role", controllers.UpdateRole) 		// Mengatur route untuk memperbarui data role
	app.Delete("/roles/delete/:id_role", controllers.DeleteRole) 	// Mengatur route untuk menghapus data role
}