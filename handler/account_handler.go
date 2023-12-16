package handler

import (
	"fmt"
	"slices"

	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find account by id
func FindAccountById(id string, account *model.Account) error {
	db := database.DB.Db
	// find single account in the database by id
	db.Find(&account, "id = ?", id)
	// if no account found, return an error
	if account.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a account
func CreateAccount(c *fiber.Ctx) error {
	db := database.DB.Db
	account := new(model.Account)
	category := new(model.Category)
	// store the body in the account and return error if encountereds
	if err := c.BodyParser(account); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find category in the database by id
	if err := FindCategoryByID(fmt.Sprint(account.IdCategory), category); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
	}
	// assign category to account
	account.Category = *category
	// create account
	if err := db.Create(account).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create account", "data": err})
	}
	// return the created account
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Account has created", "data": account})
}

// get all accounts from db
func GetAllAccounts(c *fiber.Ctx) error {
	db := database.DB.Db
	accounts := []model.Account{}
	// find all accounts in the database
	if err := db.Find(&accounts).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Accounts not found", "data": nil})
	}
	// if no account found, return an error
	if len(accounts) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Accounts not found", "data": nil})
	}
	responseAccounts := []model.Account{}
	for _, account := range accounts {
		category := new(model.Category)
		// find category in the database by id
		if err := FindCategoryByID(fmt.Sprint(account.IdCategory), category); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
		}
		// assign category to account
		account.Category = *category
		responseAccounts = append(responseAccounts, account)
	}
	// return accounts
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Accounts Found", "data": responseAccounts})
}

// get single account from db
func GetSingleAccount(c *fiber.Ctx) error {
	account := new(model.Account)
	category := new(model.Category)
	// get id params
	id := c.Params("id")
	// find single account in the database by id
	if err := FindAccountById(id, account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// find category in the database by id
	if err := FindCategoryByID(fmt.Sprint(account.IdCategory), category); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
	}
	// assign category to account
	account.Category = *category
	// return account
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Account Found", "data": account})
}

// update a account in db
func UpdateAccount(c *fiber.Ctx) error {
	db := database.DB.Db
	account := new(model.Account)
	category := new(model.Category)
	// get id params
	id := c.Params("id")
	// find single account in the database by id
	if err := FindAccountById(id, account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// store the body in the account and return error if encountereds
	if err := c.BodyParser(account); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find category in the database by id
	if err := FindCategoryByID(fmt.Sprint(account.IdCategory), category); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
	}
	// assign category to account
	account.Category = *category
	// update account
	if err := db.Save(account).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update account", "data": err})
	}
	// return the updated account
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Account has updated", "data": account})
}

// delete a account in db
func DeleteAccount(c *fiber.Ctx) error {
	db := database.DB.Db
	account := new(model.Account)
	// get id params
	id := c.Params("id")
	// find single account in the database by id
	if err := FindAccountById(id, account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// delete account
	if err := db.Delete(account).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete account", "data": err})
	}
	// return deleted account
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Account deleted"})
}

// get all input & output by date
func GetInputOutputByDate(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInput := []model.DetailInput{}
	detailOutput := []model.DetailOutput{}
	// get date params
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	// find all detail input in the database by date
	if err := db.Find(&detailInput, "input_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
	}
	// find all detail output in the database by date
	if err := db.Find(&detailOutput, "output_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
	}
	// if no detail input & output found, return an error
	if len(detailInput) == 0 && len(detailOutput) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input & Output not found", "data": nil})
	}
	// return detail input & output
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail Input & Output Found", "data": fiber.Map{"detail_input": detailInput, "detail_output": detailOutput}})
}

// get all input & output in db
func GetAllInputOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInput := []model.DetailInput{}
	detailOutput := []model.DetailOutput{}
	// find all detail input in the database
	if err := db.Find(&detailInput).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
	}
	// find all detail output in the database
	if err := db.Find(&detailOutput).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
	}
	// if no detail input & output found, return an error
	if len(detailInput) == 0 && len(detailOutput) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input & Output not found", "data": nil})
	}
	// return detail input & output
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail Input & Output Found", "data": fiber.Map{"detail_input": detailInput, "detail_output": detailOutput}})
}

// get all input & output and group by account and category in db
func GetAllInputOutputGroupByAccountAndCategory(c *fiber.Ctx) error {

	type GroupAccount struct {
		NamaAccount  string               `json:"name_account"`
		Debit        float64              `json:"debit"`
		Credit       float64              `json:"credit"`
		DetailInput  []model.DetailInput  `json:"detail_input"`
		DetailOutput []model.DetailOutput `json:"detail_output"`
	}

	type GroupCategory struct {
		NamaCategory string         `json:"name_category"`
		Accounts     []GroupAccount `json:"accounts"`
		TotalDebit   float64        `json:"total_debit"`
		TotalCredit  float64        `json:"total_credit"`
	}

	db := database.DB.Db
	categories := []model.Category{}
	accounts := []model.Account{}

	typeAccount := []string{}
	typeCategory := []string{}

	detailInputs := []model.DetailInput{}
	detailOutputs := []model.DetailOutput{}

	groupCategorys := []GroupCategory{}

	// find all categories in the database
	if err := db.Find(&categories).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Categories not found", "data": nil})
	}
	// find all accounts in the database
	if err := db.Find(&accounts).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Accounts not found", "data": nil})
	}
	// if no category & account found, return an error
	if len(categories) == 0 && len(accounts) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Categories & Accounts not found", "data": nil})
	}

	// find all type account & category
	for _, category := range categories {
		for _, account := range accounts {
			if !slices.Contains(typeCategory, fmt.Sprint(category.NameCategory)) {
				typeCategory = append(typeCategory, fmt.Sprint(category.NameCategory))
			}
			if !slices.Contains(typeAccount, fmt.Sprint(account.NameAccount)) {
				typeAccount = append(typeAccount, fmt.Sprint(account.NameAccount))
			}
		}
	}

	// find all detail input in the database
	if err := db.Find(&detailInputs).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
	}

	// find all detail output in the database
	if err := db.Find(&detailOutputs).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
	}

	for _, typeCategory := range typeCategory {
		groupCategory := GroupCategory{}
		groupCategory.NamaCategory = typeCategory
		groupCategory.TotalDebit = 0
		groupCategory.TotalCredit = 0
		for _, typeAccount := range typeAccount {
			groupAccount := GroupAccount{}
			groupAccount.NamaAccount = typeAccount
			groupAccount.Debit = 0
			groupAccount.Credit = 0
			groupAccount.DetailInput = []model.DetailInput{}
			groupAccount.DetailOutput = []model.DetailOutput{}
			for _, detailInput := range detailInputs {
				category := new(model.Category)
				account := new(model.Account)
				input := new(model.Input)
				// convert id to string
				idCategory := fmt.Sprint(detailInput.IdCategory)
				idAccount := fmt.Sprint(detailInput.IdAccount)
				idInput := fmt.Sprint(detailInput.IdInput)
				// find category in the database by id
				if err := FindCategoryByID(idCategory, category); err != nil {
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
				}
				// find account in the database by id
				if err := FindAccountById(idAccount, account); err != nil {
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
				}
				// find input in the database by id
				if err := FindInputById(idInput, input); err != nil {
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
				}
				// assign category to detail input
				detailInput.Category = *category
				// assign account to detail input
				detailInput.Account = *account
				detailInput.Account.Category = *category
				// assign input to detail input
				detailInput.Input = *input
				if detailInput.Account.NameAccount == typeAccount && detailInput.Category.NameCategory == typeCategory {
					groupAccount.Debit += float64(detailInput.TotalPrice)
					groupCategory.TotalDebit += float64(detailInput.TotalPrice)
				}
				groupAccount.DetailInput = append(groupAccount.DetailInput, detailInput)
			}
			for _, detailOutput := range detailOutputs {
				category := new(model.Category)
				account := new(model.Account)
				output := new(model.Output)
				// convert id to string
				idCategory := fmt.Sprint(detailOutput.IdCategory)
				idAccount := fmt.Sprint(detailOutput.IdAccount)
				idOutput := fmt.Sprint(detailOutput.IdOutput)
				// find category in the database by id
				if err := FindCategoryByID(idCategory, category); err != nil {
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
				}
				// find account in the database by id
				if err := FindAccountById(idAccount, account); err != nil {
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
				}
				// find input in the database by id
				if err := FindOutputById(idOutput, output); err != nil {
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
				}
				// assign category to detail input
				detailOutput.Category = *category
				// assign account to detail input
				detailOutput.Account = *account
				detailOutput.Account.Category = *category
				// assign input to detail input
				detailOutput.Output = *output
				if detailOutput.Account.NameAccount == typeAccount && detailOutput.Category.NameCategory == typeCategory {
					groupAccount.Credit += float64(detailOutput.TotalPrice)
					groupCategory.TotalCredit += float64(detailOutput.TotalPrice)
				}
				groupAccount.DetailOutput = append(groupAccount.DetailOutput, detailOutput)
			}
			groupCategory.Accounts = append(groupCategory.Accounts, groupAccount)
		}
		groupCategorys = append(groupCategorys, groupCategory)
	}
	if len(groupCategorys) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input And Output not found", "data": nil})
	}
	// return detail input & output
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail Input And Output Found", "data": fiber.Map{
		"group_category": groupCategorys,
	}})
}
