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
	account.Get("/", handler.GetAllAccounts)
	account.Get("/:id", handler.GetSingleAccount)
	account.Post("/", handler.CreateAccount)
	account.Put("/:id", handler.UpdateAccount)
	account.Delete("/:id", handler.DeleteAccount)

	// category
	category := api.Group("/category")
	category.Get("/", handler.GetAllCategories)
	category.Get("/:id", handler.GetSingleCategory)
	category.Post("/", handler.CreateCategory)
	category.Put("/:id", handler.UpdateCategory)
	category.Delete("/:id", handler.DeleteCategory)

	// user
	user := api.Group("/user")
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	// input
	input := api.Group("/input")
	input.Get("/", handler.GetAllInputs)
	input.Get("/:id", handler.GetSingleInput)
	input.Post("/", handler.CreateInput)
	input.Put("/:id", handler.UpdateInput)
	input.Delete("/:id", handler.DeleteInput)

	master := api.Group("/master")
	master.Get("/input", handler.GetInputMaster)
	master.Get("/output", handler.GetOutputMaster)

	// detail input
	detailInput := api.Group("/detail-input")
	detailInput.Get("/", handler.GetAllDetailInputs)
	detailInput.Get("/:id", handler.GetSingleDetailInput)
	detailInput.Post("/", handler.CreateDetailInput)
	detailInput.Put("/:id", handler.UpdateDetailInput)
	detailInput.Delete("/:id", handler.DeleteDetailInput)
	detailInput.Get("/input/:id", handler.GetDetailInputsByInputID)

	// detail inputs
	detailInputs := api.Group("/detail-inputs")
	detailInputs.Post("/", handler.CreateMultipleDetailInputs)
	detailInputs.Put("/:id", handler.UpdateMultipleDetailInputs)

	// output
	output := api.Group("/output")
	output.Get("/", handler.GetAllOutputs)
	output.Get("/:id", handler.GetSingleOutput)
	output.Post("/", handler.CreateOutput)
	output.Put("/:id", handler.UpdateOutput)
	output.Delete("/:id", handler.DeleteOutput)

	// detail output
	detailOutput := api.Group("/detail-output")
	detailOutput.Get("/", handler.GetAllDetailOutputs)
	detailOutput.Get("/:id", handler.GetSingleDetailOutput)
	detailOutput.Post("/", handler.CreateDetailOutput)
	detailOutput.Put("/:id", handler.UpdateDetailOutput)
	detailOutput.Delete("/:id", handler.DeleteDetailOutput)
	detailOutput.Get("/output/:id", handler.GetDetailOutputsByOutputID)

	// detail outputs
	detailOutputs := api.Group("/detail-outputs")
	detailOutputs.Post("/", handler.CreateMultipleDetailOutputs)
	detailOutputs.Put("/:id", handler.UpdateMultipleDetailOutputs)

	// general journal
	generalJournal := api.Group("/general-journal")
	generalJournal.Get("/", handler.GetAllGeneralJournals)
	generalJournal.Get("/:id", handler.GetSingleGeneralJournal)
	generalJournal.Post("/", handler.CreateGeneralJournal)
	generalJournal.Put("/:id", handler.UpdateGeneralJournal)
	generalJournal.Delete("/:id", handler.DeleteGeneralJournal)

	// transaction
	transaction := api.Group("/transaction")
	transaction.Get("/", handler.GetTransactions)
	transaction.Get("/group", handler.GetTransactionGroupByAccount)
	transaction.Get("/date/filter", handler.GetTransactionFilterByDate)
	transaction.Get("/date/group", handler.GetTransactionGroupByDate)
	transaction.Get("/total-transaction", handler.GetTotalTransaction)
	transaction.Get("/cash-flow", handler.GetCashFlow)
	transaction.Get("/profit-loss", handler.GetProfitLoss)
	transaction.Get("/capital-change", handler.GetCapitalChange)

	// login
	login := api.Group("/login")
	login.Post("/", handler.Login)

	// logout
	logout := api.Group("/logout")
	logout.Post("/", handler.Logout)

	// user login
	userLogin := api.Group("/user-login")
	userLogin.Get("/", handler.GetLogedUser)
}
