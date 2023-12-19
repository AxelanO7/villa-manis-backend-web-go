package handler

import (
	"fmt"
	"slices"
	"time"

	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// get all input & output by date
func GetTransactionFilterByDate(c *fiber.Ctx) error {
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
	// adjust category, account, and input in detail input
	if len(detailInput) != 0 {
		for i := range detailInput {
			category := new(model.Category)
			account := new(model.Account)
			input := new(model.Input)
			// convert id to string
			idCategory := fmt.Sprint(detailInput[i].IdCategory)
			idAccount := fmt.Sprint(detailInput[i].IdAccount)
			idInput := fmt.Sprint(detailInput[i].IdInput)
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
			detailInput[i].Category = *category
			// assign account to detail input
			detailInput[i].Account = *account
			detailInput[i].Account.Category = *category
			// assign input to detail input
			detailInput[i].Input = *input
		}
	}
	// adjust category, account, and output in detail output
	if len(detailOutput) != 0 {
		for i := range detailOutput {
			category := new(model.Category)
			account := new(model.Account)
			output := new(model.Output)
			// convert id to string
			idCategory := fmt.Sprint(detailOutput[i].IdCategory)
			idAccount := fmt.Sprint(detailOutput[i].IdAccount)
			idOutput := fmt.Sprint(detailOutput[i].IdOutput)
			// find category in the database by id
			if err := FindCategoryByID(idCategory, category); err != nil {
				return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
			}
			// find account in the database by id
			if err := FindAccountById(idAccount, account); err != nil {
				return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
			}
			// find output in the database by id
			if err := FindOutputById(idOutput, output); err != nil {
				return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
			}
			// assign category to detail output
			detailOutput[i].Category = *category
			// assign account to detail output
			detailOutput[i].Account = *account
			detailOutput[i].Account.Category = *category
			// assign output to detail output
			detailOutput[i].Output = *output
		}
	}
	// if no detail input & output found, return an error
	if len(detailInput) == 0 && len(detailOutput) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input & Output not found", "data": nil})
	}
	// return detail input & output
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail Input & Output Found", "data": fiber.Map{"detail_input": detailInput, "detail_output": detailOutput}})
}

// get all input & output in db
func GetTransactions(c *fiber.Ctx) error {
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
	// adjust category, account, and input in detail input
	if len(detailInput) != 0 {
		for i := range detailInput {
			category := new(model.Category)
			account := new(model.Account)
			input := new(model.Input)
			// convert id to string
			idCategory := fmt.Sprint(detailInput[i].IdCategory)
			idAccount := fmt.Sprint(detailInput[i].IdAccount)
			idInput := fmt.Sprint(detailInput[i].IdInput)
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
			detailInput[i].Category = *category
			// assign account to detail input
			detailInput[i].Account = *account
			detailInput[i].Account.Category = *category
			// assign input to detail input
			detailInput[i].Input = *input
		}
	}
	// adjust category, account, and output in detail output
	if len(detailOutput) != 0 {
		for i := range detailOutput {
			category := new(model.Category)
			account := new(model.Account)
			output := new(model.Output)
			// convert id to string
			idCategory := fmt.Sprint(detailOutput[i].IdCategory)
			idAccount := fmt.Sprint(detailOutput[i].IdAccount)
			idOutput := fmt.Sprint(detailOutput[i].IdOutput)
			// find category in the database by id
			if err := FindCategoryByID(idCategory, category); err != nil {
				return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
			}
			// find account in the database by id
			if err := FindAccountById(idAccount, account); err != nil {
				return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
			}
			// find output in the database by id
			if err := FindOutputById(idOutput, output); err != nil {
				return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
			}
			// assign category to detail output
			detailOutput[i].Category = *category
			// assign account to detail output
			detailOutput[i].Account = *account
			detailOutput[i].Account.Category = *category
			// assign output to detail output
			detailOutput[i].Output = *output
		}
	}

	// if no detail input & output found, return an error
	if len(detailInput) == 0 && len(detailOutput) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input & Output not found", "data": nil})
	}
	// return detail input & output
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail Input & Output Found", "data": fiber.Map{"detail_input": detailInput, "detail_output": detailOutput}})
}

// get all input & output and group by account and category in db
func GetTransactionGroupByAccount(c *fiber.Ctx) error {
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

	// get date params
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

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
	// // find all detail input in the database
	// if err := db.Find(&detailInputs).Error; err != nil {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
	// }
	// // find all detail output in the database
	// if err := db.Find(&detailOutputs).Error; err != nil {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
	// }
	if startDate == "" || endDate == "" {
		// find all detail input in the database
		if err := db.Find(&detailInputs).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database
		if err := db.Find(&detailOutputs).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}
	if startDate != "" && endDate != "" {
		// find all detail input in the database by date
		if err := db.Find(&detailInputs, "created_at BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database by date
		if err := db.Find(&detailOutputs, "created_at BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
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

// get transactions group by date
func GetTransactionGroupByDate(c *fiber.Ctx) error {
	type GroupDate struct {
		Date         string               `json:"date"`
		TotalDebit   float64              `json:"total_debit"`
		TotalCredit  float64              `json:"total_credit"`
		DetailInput  []model.DetailInput  `json:"detail_input"`
		DetailOutput []model.DetailOutput `json:"detail_output"`
	}

	db := database.DB.Db
	detailInputs := []model.DetailInput{}
	detailOutputs := []model.DetailOutput{}

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	groupDates := []GroupDate{}

	if startDate == "" || endDate == "" {
		// find all detail input in the database
		if err := db.Find(&detailInputs).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database
		if err := db.Find(&detailOutputs).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}
	if startDate != "" && endDate != "" {
		// find all detail input in the database by date
		if err := db.Find(&detailInputs, "created_at BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database by date
		if err := db.Find(&detailOutputs, "created_at BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}
	groupDatesStrings := []string{}
	// find all date in detail input
	for _, detailInput := range detailInputs {
		if !slices.Contains(groupDatesStrings, fmt.Sprint(
			time.Time(detailInput.CreatedAt).Format("2006-01-02"),
		)) {
			groupDatesStrings = append(groupDatesStrings, fmt.Sprint(
				time.Time(detailInput.CreatedAt).Format("2006-01-02"),
			))

		}
	}
	// find all date in detail output
	for _, detailOutput := range detailOutputs {
		if !slices.Contains(groupDatesStrings, fmt.Sprint(
			time.Time(detailOutput.CreatedAt).Format("2006-01-02"),
		)) {
			groupDatesStrings = append(groupDatesStrings, fmt.Sprint(
				time.Time(detailOutput.CreatedAt).Format("2006-01-02"),
			))
		}
	}
	// group by date
	for _, groupDate := range groupDatesStrings {
		groupDate := GroupDate{
			Date:         groupDate,
			TotalDebit:   0,
			TotalCredit:  0,
			DetailInput:  []model.DetailInput{},
			DetailOutput: []model.DetailOutput{},
		}
		for _, detailInput := range detailInputs {
			if fmt.Sprint(time.Time(detailInput.CreatedAt).Format("2006-01-02")) == groupDate.Date {

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

				groupDate.TotalDebit += float64(detailInput.TotalPrice)
				groupDate.DetailInput = append(groupDate.DetailInput, detailInput)
			}
		}
		for _, detailOutput := range detailOutputs {
			if fmt.Sprint(time.Time(detailOutput.CreatedAt).Format("2006-01-02")) == groupDate.Date {

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

				groupDate.TotalCredit += float64(detailOutput.TotalPrice)
				groupDate.DetailOutput = append(groupDate.DetailOutput, detailOutput)
			}
		}
		groupDates = append(groupDates, groupDate)
	}
	// if no detail input & output found, return an error
	if len(groupDates) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input & Output not found", "data": nil})
	}
	// return detail input & output
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail Input & Output Found", "data": fiber.Map{"group_date": groupDates}})
}
