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
	app.Use(controllers.Authenticate)
	app.Get("/auth", controllers.GetUser)

	// Mengatur route untuk data buku
	app.Get("/books", controllers.GetAllBooks) 					// Mengatur route untuk mengambil semua data buku
	app.Get("/books/get/:id", controllers.GetBookByID)			// Mengatur route untuk mengambil data buku berdasarkan ID
	app.Post("/books/create", controllers.CreateBook) 			// Mengatur route untuk menambahkan data buku
	app.Put("/books/update/:id", controllers.UpdateBook) 		// Mengatur route untuk memperbarui data buku
	app.Delete("/books/delete/:id", controllers.DeleteBook) 	// Mengatur route untuk menghapus data buku

	// Mengatur route untuk data role
	app.Get("/roles", controllers.GetAllRoles) 					// Mengatur route untuk mengambil semua data role
	app.Get("/roles/get/:id", controllers.GetRoleByID)			// Mengatur route untuk mengambil data role berdasarkan ID
	app.Post("/roles/create", controllers.CreateRole) 			// Mengatur route untuk menambahkan data role
	app.Put("/roles/update/:id", controllers.UpdateRole) 		// Mengatur route untuk memperbarui data role
	app.Delete("/roles/delete/:id", controllers.DeleteRole) 	// Mengatur route untuk menghapus data role

	// Mengatur route untuk logout
	app.Post("/logout", controllers.LogoutUser)
}