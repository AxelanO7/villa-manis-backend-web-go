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
	if err := db.Find(&detailInput, "input_date BETWEEN ? AND ?", startDate+" 00:00:00", endDate+" 23:59:59").Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
	}
	// find all detail output in the database by date
	if err := db.Find(&detailOutput, "output_date BETWEEN ? AND ?", startDate+" 00:00:00", endDate+" 23:59:59").Error; err != nil {
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

	typeCategories := []model.Category{}
	typeAccount := []model.Account{}

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
	categoryString := []string{}
	accountString := []string{}
	// find all type account & category
	for _, category := range categories {
		for _, account := range accounts {
			if !slices.Contains(categoryString, fmt.Sprint(category.NameCategory)) {
				categoryString = append(categoryString, fmt.Sprint(category.NameCategory))
				typeCategories = append(typeCategories, category)
			}
			if !slices.Contains(accountString, fmt.Sprint(account.NameAccount)) {
				accountString = append(accountString, fmt.Sprint(account.NameAccount))
				typeAccount = append(typeAccount, account)
			}
		}
	}

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
		if err := db.Find(&detailInputs, "input_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database by date
		if err := db.Find(&detailOutputs, "output_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}

	for _, typeCategory := range typeCategories {
		groupCategory := GroupCategory{}
		groupCategory.NamaCategory = typeCategory.NameCategory
		groupCategory.TotalDebit = 0
		groupCategory.TotalCredit = 0
		for _, typeAccount := range typeAccount {
			if typeAccount.IdCategory == int(typeCategory.ID) {
				groupAccount := GroupAccount{}
				groupAccount.NamaAccount = typeAccount.NameAccount
				groupAccount.Debit = 0
				groupAccount.Credit = 0
				groupAccount.DetailInput = []model.DetailInput{}
				groupAccount.DetailOutput = []model.DetailOutput{}
				for _, detailInput := range detailInputs {
					if detailInput.IdAccount == int(typeAccount.ID) && detailInput.IdCategory == int(typeCategory.ID) {
						category := new(model.Category)
						account := new(model.Account)
						input := new(model.Input)
						// convert id to string
						idCategory := fmt.Sprint(detailInput.IdCategory)
						idAccount := fmt.Sprint(detailInput.IdAccount)
						idInput := fmt.Sprint(detailInput.IdInput)
						// find category in the database by id
						if err := FindCategoryByID(idCategory, category); err != nil {
							continue
							return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
						}
						// find account in the database by id
						if err := FindAccountById(idAccount, account); err != nil {
							continue
							return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
						}
						// find input in the database by id
						if err := FindInputById(idInput, input); err != nil {
							continue
							return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
						}
						// assign category to detail input
						detailInput.Category = *category
						// assign account to detail input
						detailInput.Account = *account
						detailInput.Account.Category = *category
						// assign input to detail input
						detailInput.Input = *input
						if detailInput.IdCategory == int(typeCategory.ID) && detailInput.IdAccount == int(typeAccount.ID) {
							groupAccount.Debit += float64(detailInput.TotalPrice)
							groupCategory.TotalDebit += float64(detailInput.TotalPrice)
						}
						groupAccount.DetailInput = append(groupAccount.DetailInput, detailInput)
					}
				}
				for _, detailOutput := range detailOutputs {
					if detailOutput.IdAccount == int(typeAccount.ID) && detailOutput.IdCategory == int(typeCategory.ID) {
						category := new(model.Category)
						account := new(model.Account)
						output := new(model.Output)
						// convert id to string
						idCategory := fmt.Sprint(detailOutput.IdCategory)
						idAccount := fmt.Sprint(detailOutput.IdAccount)
						idOutput := fmt.Sprint(detailOutput.IdOutput)
						// find category in the database by id
						if err := FindCategoryByID(idCategory, category); err != nil {
							continue
							return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
						}
						// find account in the database by id
						if err := FindAccountById(idAccount, account); err != nil {
							continue
							return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
						}
						// find input in the database by id
						if err := FindOutputById(idOutput, output); err != nil {
							continue
							return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
						}
						// assign category to detail input
						detailOutput.Category = *category
						// assign account to detail input
						detailOutput.Account = *account
						detailOutput.Account.Category = *category
						// assign input to detail input
						detailOutput.Output = *output
						if detailOutput.IdCategory == int(typeCategory.ID) && detailOutput.IdAccount == int(typeAccount.ID) {
							groupAccount.Credit += float64(detailOutput.TotalPrice)
							groupCategory.TotalCredit += float64(detailOutput.TotalPrice)
						}
						groupAccount.DetailOutput = append(groupAccount.DetailOutput, detailOutput)
					}
				}
				groupCategory.Accounts = append(groupCategory.Accounts, groupAccount)
			}
		}
		groupCategorys = append(groupCategorys, groupCategory)
	}
	lengthaccount := 0
	for _, groupCategory := range groupCategorys {
		lengthaccount += len(groupCategory.Accounts)
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
		if err := db.Find(&detailInputs, "input_date BETWEEN ? AND ?", startDate+" 00:00:00", endDate+" 23:59:59").Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database by date
		if err := db.Find(&detailOutputs, "output_date BETWEEN ? AND ?", startDate+" 00:00:00", endDate+" 23:59:59").Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}
	groupDatesStrings := []string{}
	// find all date in detail input
	for _, detailInput := range detailInputs {
		if !slices.Contains(groupDatesStrings, fmt.Sprint(
			detailInput.InputDate,
		)) {
			groupDatesStrings = append(groupDatesStrings, fmt.Sprint(
				detailInput.InputDate,
			))

		}
	}
	// find all date in detail output
	for _, detailOutput := range detailOutputs {
		if !slices.Contains(groupDatesStrings, fmt.Sprint(
			detailOutput.OutputDate,
		)) {
			groupDatesStrings = append(groupDatesStrings, fmt.Sprint(
				detailOutput.OutputDate,
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
			if fmt.Sprint(detailInput.InputDate) == groupDate.Date {

				category := new(model.Category)
				account := new(model.Account)
				input := new(model.Input)
				// convert id to string
				idCategory := fmt.Sprint(detailInput.IdCategory)
				idAccount := fmt.Sprint(detailInput.IdAccount)
				idInput := fmt.Sprint(detailInput.IdInput)
				// find category in the database by id
				if err := FindCategoryByID(idCategory, category); err != nil {
					continue
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
				}
				// find account in the database by id
				if err := FindAccountById(idAccount, account); err != nil {
					continue
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
				}
				// find input in the database by id
				if err := FindInputById(idInput, input); err != nil {
					continue
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
			if fmt.Sprint(detailOutput.OutputDate) == groupDate.Date {

				category := new(model.Category)
				account := new(model.Account)
				output := new(model.Output)
				// convert id to string
				idCategory := fmt.Sprint(detailOutput.IdCategory)
				idAccount := fmt.Sprint(detailOutput.IdAccount)
				idOutput := fmt.Sprint(detailOutput.IdOutput)
				// find category in the database by id
				if err := FindCategoryByID(idCategory, category); err != nil {
					continue
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
				}
				// find account in the database by id
				if err := FindAccountById(idAccount, account); err != nil {
					continue
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
				}
				// find input in the database by id
				if err := FindOutputById(idOutput, output); err != nil {
					continue
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
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

// get total transactions
func GetTotalTransaction(c *fiber.Ctx) error {
	// month struct
	type Month struct {
		January   float32 `json:"january"`
		February  float32 `json:"february"`
		March     float32 `json:"march"`
		April     float32 `json:"april"`
		May       float32 `json:"may"`
		June      float32 `json:"june"`
		July      float32 `json:"july"`
		August    float32 `json:"august"`
		September float32 `json:"september"`
		October   float32 `json:"october"`
		November  float32 `json:"november"`
		December  float32 `json:"december"`
	}

	db := database.DB.Db
	detailInput := new([]model.DetailInput)
	detailOutput := new([]model.DetailOutput)
	totalDebit := 0
	totalCredit := 0
	totalDebitMonth := Month{}
	totalCreditMonth := Month{}

	dateStart := c.Query("date-start")
	dateEnd := c.Query("date-end")

	if dateStart == "" || dateEnd == "" {
		// find all detail input in the database
		if err := db.Find(&detailInput).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database
		if err := db.Find(&detailOutput).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}
	if dateStart != "" && dateEnd != "" {
		if err := db.Find(&detailInput, "input_date BETWEEN ? AND ?", dateStart, dateEnd).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database
		if err := db.Find(&detailOutput, "output_date BETWEEN ? AND ?", dateStart, dateEnd).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}
	// if no detail input & output found, return an error
	if len(*detailInput) == 0 && len(*detailOutput) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input & Output not found", "data": nil})
	}

	for i := 0; i < len(*detailInput); i++ {
		for j := 1; j <= 12; j++ {
			if time.Time((*detailInput)[i].CreatedAt).Month() == time.Month(j) && time.Time((*detailInput)[i].CreatedAt).Year() == time.Now().Year() {
				switch j {
				case 1:
					totalDebitMonth.January += float32((*detailInput)[i].TotalPrice)
				case 2:
					totalDebitMonth.February += float32((*detailInput)[i].TotalPrice)
				case 3:
					totalDebitMonth.March += float32((*detailInput)[i].TotalPrice)
				case 4:
					totalDebitMonth.April += float32((*detailInput)[i].TotalPrice)
				case 5:
					totalDebitMonth.May += float32((*detailInput)[i].TotalPrice)
				case 6:
					totalDebitMonth.June += float32((*detailInput)[i].TotalPrice)
				case 7:
					totalDebitMonth.July += float32((*detailInput)[i].TotalPrice)
				case 8:
					totalDebitMonth.August += float32((*detailInput)[i].TotalPrice)
				case 9:
					totalDebitMonth.September += float32((*detailInput)[i].TotalPrice)
				case 10:
					totalDebitMonth.October += float32((*detailInput)[i].TotalPrice)
				case 11:
					totalDebitMonth.November += float32((*detailInput)[i].TotalPrice)
				case 12:
					totalDebitMonth.December += float32((*detailInput)[i].TotalPrice)
				default:
					totalDebitMonth.January += float32((*detailInput)[i].TotalPrice)
				}
			}
			totalDebit += (*detailInput)[i].TotalPrice
		}
	}
	for i := 0; i < len(*detailOutput); i++ {
		for j := 1; j <= 12; j++ {
			if time.Time((*detailOutput)[i].CreatedAt).Month() == time.Month(j) && time.Time((*detailOutput)[i].CreatedAt).Year() == time.Now().Year() {
				switch j {
				case 1:
					totalCreditMonth.January += float32((*detailOutput)[i].TotalPrice)
				case 2:
					totalCreditMonth.February += float32((*detailOutput)[i].TotalPrice)
				case 3:
					totalCreditMonth.March += float32((*detailOutput)[i].TotalPrice)
				case 4:
					totalCreditMonth.April += float32((*detailOutput)[i].TotalPrice)
				case 5:
					totalCreditMonth.May += float32((*detailOutput)[i].TotalPrice)
				case 6:
					totalCreditMonth.June += float32((*detailOutput)[i].TotalPrice)
				case 7:
					totalCreditMonth.July += float32((*detailOutput)[i].TotalPrice)
				case 8:
					totalCreditMonth.August += float32((*detailOutput)[i].TotalPrice)
				case 9:
					totalCreditMonth.September += float32((*detailOutput)[i].TotalPrice)
				case 10:
					totalCreditMonth.October += float32((*detailOutput)[i].TotalPrice)
				case 11:
					totalCreditMonth.November += float32((*detailOutput)[i].TotalPrice)
				case 12:
					totalCreditMonth.December += float32((*detailOutput)[i].TotalPrice)
				default:
					totalCreditMonth.January += float32((*detailOutput)[i].TotalPrice)
				}
			}
			totalCredit += (*detailOutput)[i].TotalPrice
		}
	}
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail Input & Output Found", "data": fiber.Map{"total_debit": totalDebit, "total_credit": totalCredit,
		"total_debit_month": totalDebitMonth, "total_credit_month": totalCreditMonth}})
}

func GetCashFlow(c *fiber.Ctx) error {
	type Account struct {
		NameAccount  string  `json:"name_account"`
		TotalAccount float64 `json:"total_account"`
	}

	type Group struct {
		NameGroup  string    `json:"name_group"`
		Accounts   []Account `json:"accounts"`
		TotalGroup float64   `json:"total_group"`
	}

	type CashFlow struct {
		Group []Group `json:"groups"`
		Total float64 `json:"totals"`
	}

	db := database.DB.Db
	detailInputs := []model.DetailInput{}
	detailOutputs := []model.DetailOutput{}
	accounts := []model.Account{}
	groups := []Group{}
	cashFlow := CashFlow{}

	listAccountOperational := []string{
		"Beban Gaji", "Beban Perlengkapan", "Beban Listrik", "Beban Iklan", "Beban Asuransi", "Beban Pemeliharaan Peralatan", "Beban Penyusutan Peralatan", "Beban Lain-lain", "Pendapatan Usaha",
	}

	listAccountInvestation := []string{
		"Bank BRI", "Bank Mandiri",
	}

	listAccountFinancing := []string{
		"Kas",
	}

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

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
		if err := db.Find(&detailInputs, "input_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database by date
		if err := db.Find(&detailOutputs, "output_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}

	// if no detail input & output found, return an error
	if len(detailInputs) == 0 && len(detailOutputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input & Output not found", "data": nil})
	}

	// find all accounts in the database
	if err := db.Find(&accounts).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Accounts not found", "data": nil})
	}

	groups = []Group{
		{
			NameGroup:  "operational",
			Accounts:   []Account{},
			TotalGroup: 0,
		},
		{
			NameGroup:  "investation",
			Accounts:   []Account{},
			TotalGroup: 0,
		},
		{
			NameGroup:  "financing",
			Accounts:   []Account{},
			TotalGroup: 0,
		},
	}

	cashFlow = CashFlow{
		Group: groups,
		Total: 0,
	}

	for _, account := range accounts {
		if slices.Contains(listAccountOperational, account.NameAccount) {
			groups[0].Accounts = append(groups[0].Accounts, Account{
				NameAccount:  account.NameAccount,
				TotalAccount: 0,
			})
		}
		if slices.Contains(listAccountInvestation, account.NameAccount) {
			groups[1].Accounts = append(groups[1].Accounts, Account{
				NameAccount:  account.NameAccount,
				TotalAccount: 0,
			})
		}
		if slices.Contains(listAccountFinancing, account.NameAccount) {
			groups[2].Accounts = append(groups[2].Accounts, Account{
				NameAccount:  account.NameAccount,
				TotalAccount: 0,
			})
		}
	}

	for _, detailInput := range detailInputs {
		accountDetailInput := new(model.Account)
		// convert id to string
		idAccount := fmt.Sprint(detailInput.IdAccount)
		// find account in the database by id
		if err := FindAccountById(idAccount, accountDetailInput); err != nil {
			continue
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign account to detail input
		detailInput.Account = *accountDetailInput
		for i, group := range groups {
			for j, account := range group.Accounts {
				if detailInput.Account.NameAccount == account.NameAccount {
					groups[i].Accounts[j].TotalAccount += float64(detailInput.TotalPrice)
					groups[i].TotalGroup += float64(detailInput.TotalPrice)
					cashFlow.Total += float64(detailInput.TotalPrice)
				}
			}
		}
	}

	for _, detailOutput := range detailOutputs {
		accountDetailOutput := new(model.Account)
		// convert id to string
		idAccount := fmt.Sprint(detailOutput.IdAccount)
		// find account in the database by id
		if err := FindAccountById(idAccount, accountDetailOutput); err != nil {
			continue
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign account to detail input
		detailOutput.Account = *accountDetailOutput
		for i, group := range groups {
			for j, account := range group.Accounts {
				if detailOutput.Account.NameAccount == account.NameAccount {
					groups[i].Accounts[j].TotalAccount -= float64(detailOutput.TotalPrice)
					groups[i].TotalGroup -= float64(detailOutput.TotalPrice)
					cashFlow.Total -= float64(detailOutput.TotalPrice)
				}
			}
		}
	}
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Cash Flow Found", "data": cashFlow})
}

func GetProfitLoss(c *fiber.Ctx) error {

	type Item struct {
		Name   string  `json:"name"`
		Credit float64 `json:"credit"`
		Debit  float64 `json:"debit"`
	}

	type Total struct {
		Income  float64 `json:"income"`
		Burden  float64 `json:"burden"`
		Balance float64 `json:"balance"`
	}

	type ProfitLoss struct {
		Income []Item `json:"income"`
		Burden []Item `json:"burden"`
		Total  Total  `json:"total"`
	}

	db := database.DB.Db
	detailInputs := []model.DetailInput{}
	detailOutputs := []model.DetailOutput{}

	accounts := []model.Account{}
	categorys := []model.Category{}

	profitLoss := ProfitLoss{}

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	listIncome := []string{"Kas", "Pendapatan Usaha"}
	// listBurden := []string{"Beban Gaji", "Beban Perlengkapan", "Beban Listrik", "Beban Iklan", "Beban Asuransi", "Beban Pemeliharaan Peralatan", "Beban Penyusutan Peralatan", "Beban Lain-lain, Kas"}

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
		if err := db.Find(&detailInputs, "input_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database by date
		if err := db.Find(&detailOutputs, "output_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}

	// if no detail input & output found, return an error
	if len(detailInputs) == 0 && len(detailOutputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input & Output not found", "data": nil})
	}

	// find all accounts in the database
	if err := db.Find(&accounts).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Accounts not found", "data": nil})
	}

	if err := db.Find(&categorys).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Categorys not found", "data": nil})
	}

	profitLoss = ProfitLoss{
		Income: []Item{},
		Burden: []Item{},
		Total: Total{
			Income:  0,
			Burden:  0,
			Balance: 0,
		},
	}

	for _, account := range accounts {
		category := new(model.Category)
		FindCategoryByID(fmt.Sprint(account.IdCategory), category)
		account.Category = *category

		fmt.Println("accounts ", account.NameAccount)
		if slices.Contains(listIncome, account.NameAccount) {
			fmt.Println("incomes ", account.NameAccount)
			profitLoss.Income = append(profitLoss.Income, Item{
				Name:   account.NameAccount,
				Credit: 0,
				Debit:  0,
			})
		}
		// if slices.Contains(listBurden, account.NameAccount) {
		if account.Category.NameCategory == "Beban" {
			fmt.Println("burdens ", account.NameAccount)
			profitLoss.Burden = append(profitLoss.Burden, Item{
				Name:   account.NameAccount,
				Credit: 0,
				Debit:  0,
			})
		}
	}

	fmt.Println("profitLoss ", profitLoss.Income)
	fmt.Println("profitLoss ", profitLoss.Burden)

	for _, detailInput := range detailInputs {
		accountDetailInput := new(model.Account)
		// convert id to string
		idAccount := fmt.Sprint(detailInput.IdAccount)
		// find account in the database by id
		if err := FindAccountById(idAccount, accountDetailInput); err != nil {
			continue
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign account to detail input
		detailInput.Account = *accountDetailInput
		for i, item := range profitLoss.Income {
			if detailInput.Account.NameAccount == item.Name {
				profitLoss.Income[i].Debit += float64(detailInput.TotalPrice)
				profitLoss.Total.Income += float64(detailInput.TotalPrice)
				profitLoss.Total.Balance += float64(detailInput.TotalPrice)
			}
		}
	}

	for _, detailOutput := range detailOutputs {
		accountDetailOutput := new(model.Account)
		// convert id to string
		idAccount := fmt.Sprint(detailOutput.IdAccount)
		// find account in the database by id
		if err := FindAccountById(idAccount, accountDetailOutput); err != nil {
			continue
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign account to detail input
		detailOutput.Account = *accountDetailOutput
		for i, item := range profitLoss.Burden {
			if detailOutput.Account.NameAccount == item.Name {
				profitLoss.Burden[i].Credit += float64(detailOutput.TotalPrice)
				profitLoss.Total.Burden += float64(detailOutput.TotalPrice)
				profitLoss.Total.Balance -= float64(detailOutput.TotalPrice)
			}
		}
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Profit Loss Found", "data": profitLoss})
}

func GetCapitalChange(c *fiber.Ctx) error {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	db := database.DB.Db
	detailInputs := []model.DetailInput{}
	detailOutputs := []model.DetailOutput{}

	listIncome := []string{"Kas", "Pendapatan Usaha"}
	listBurden := []string{"Beban Gaji", "Beban Perlengkapan", "Beban Listrik", "Beban Iklan", "Beban Asuransi", "Beban Pemeliharaan Peralatan", "Beban Penyusutan Peralatan", "Beban Lain-lain, Kas"}

	accounts := []model.Account{}

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
		if err := db.Find(&detailInputs, "input_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input not found", "data": nil})
		}
		// find all detail output in the database by date
		if err := db.Find(&detailOutputs, "output_date BETWEEN ? AND ?", startDate, endDate).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
		}
	}

	// if no detail input & output found, return an error
	if len(detailInputs) == 0 && len(detailOutputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Input & Output not found", "data": nil})
	}

	// find all accounts in the database
	if err := db.Find(&accounts).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Accounts not found", "data": nil})
	}

	beginningCapital := 0
	netIncome := 0
	prive := 0
	addCapital := 0
	endCapital := 0

	for _, detailInput := range detailInputs {
		accountDetailInput := new(model.Account)
		// convert id to string
		idAccount := fmt.Sprint(detailInput.IdAccount)
		// find account in the database by id
		if err := FindAccountById(idAccount, accountDetailInput); err != nil {
			continue
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign account to detail input
		detailInput.Account = *accountDetailInput
		if detailInput.Account.NameAccount == "Modal Awal" {
			beginningCapital += detailInput.TotalPrice
		}
		if detailInput.Account.NameAccount == "Prive" {
			prive += detailInput.TotalPrice
		}
		if detailInput.Account.NameAccount == "Penambahan Modal" {
			addCapital += detailInput.TotalPrice
		}
	}

	for _, detailOutput := range detailOutputs {
		accountDetailOutput := new(model.Account)
		// convert id to string
		idAccount := fmt.Sprint(detailOutput.IdAccount)
		// find account in the database by id
		if err := FindAccountById(idAccount, accountDetailOutput); err != nil {
			continue
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign account to detail input
		detailOutput.Account = *accountDetailOutput
		if detailOutput.Account.NameAccount == "Prive" {
			prive -= detailOutput.TotalPrice
		}
		if detailOutput.Account.NameAccount == "Penambahan Modal" {
			addCapital -= detailOutput.TotalPrice
		}
	}

	for _, account := range accounts {
		if slices.Contains(listIncome, account.NameAccount) {
			for _, detailInput := range detailInputs {
				accountDetailInput := new(model.Account)
				// convert id to string
				idAccount := fmt.Sprint(detailInput.IdAccount)
				// find account in the database by id
				if err := FindAccountById(idAccount, accountDetailInput); err != nil {
					continue
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
				}
				// assign account to detail input
				detailInput.Account = *accountDetailInput
				if detailInput.Account.NameAccount == account.NameAccount {
					netIncome += detailInput.TotalPrice
				}
			}
		}
		if slices.Contains(listBurden, account.NameAccount) {
			for _, detailOutput := range detailOutputs {
				accountDetailOutput := new(model.Account)
				// convert id to string
				idAccount := fmt.Sprint(detailOutput.IdAccount)
				// find account in the database by id
				if err := FindAccountById(idAccount, accountDetailOutput); err != nil {
					continue
					return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
				}
				// assign account to detail input
				detailOutput.Account = *accountDetailOutput
				if detailOutput.Account.NameAccount == account.NameAccount {
					netIncome -= detailOutput.TotalPrice
				}
			}
		}
	}
	endCapital = beginningCapital + netIncome + addCapital - prive
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Capital Change Found", "data": fiber.Map{"beginning_capital": beginningCapital, "net_income": netIncome, "prive": prive, "additional_capital": addCapital, "end_capital": endCapital}})
}
