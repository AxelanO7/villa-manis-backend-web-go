package router

import (
	"github.com/AxelanO7/villa-manis-backend-web-go/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")

	// account
	account := api.Group("/account")
	// routes
	account.Get("/", handler.GetAllAccounts)
	account.Get("/:id", handler.GetSingleAccount)
	account.Post("/", handler.CreateAccount)
	account.Put("/:id", handler.UpdateAccount)
	account.Delete("/:id", handler.DeleteAccount)

	// category
	category := api.Group("/category")
	// routes
	category.Get("/", handler.GetAllCategories)
	category.Get("/:id", handler.GetSingleCategory)
	category.Post("/", handler.CreateCategory)
	category.Put("/:id", handler.UpdateCategory)
	category.Delete("/:id", handler.DeleteCategory)

	// user
	user := api.Group("/user")
	// routes
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	// input
	input := api.Group("/input")
	// routes
	input.Get("/", handler.GetAllInputs)
	input.Get("/:id", handler.GetSingleInput)
	input.Post("/", handler.CreateInput)
	input.Put("/:id", handler.UpdateInput)
	input.Delete("/:id", handler.DeleteInput)

	// output
	output := api.Group("/output")
	// routes
	output.Get("/", handler.GetAllOutputs)
	output.Get("/:id", handler.GetSingleOutput)
	output.Post("/", handler.CreateOutput)
	output.Put("/:id", handler.UpdateOutput)
	output.Delete("/:id", handler.DeleteOutput)

	// general journal
	generalJournal := api.Group("/general-journal")
	// routes
	generalJournal.Get("/", handler.GetAllGeneralJournals)
	generalJournal.Get("/:id", handler.GetSingleGeneralJournal)
	generalJournal.Post("/", handler.CreateGeneralJournal)
	generalJournal.Put("/:id", handler.UpdateGeneralJournal)
	generalJournal.Delete("/:id", handler.DeleteGeneralJournal)
}
