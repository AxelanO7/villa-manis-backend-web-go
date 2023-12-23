package handler

import (
	"fmt"

	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find detail input by id
func findDetailInputById(id string, detailInput *model.DetailInput) error {
	db := database.DB.Db
	// find single detail input in the database by id
	db.Find(&detailInput, "id = ?", id)
	// if no detail input found, return an error
	if detailInput.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a detail input
func CreateDetailInput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInput := new(model.DetailInput)
	input := new(model.Input)
	account := new(model.Account)
	// store the body in the detail input and return error if encountereds
	if err := c.BodyParser(detailInput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find input in the database by id
	if err := FindInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
	}
	// find account in the database by id
	if err := FindAccountById(fmt.Sprint(detailInput.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// assign input to detail input
	detailInput.Input = *input
	// assign account to detail input
	detailInput.Account = *account
	// create detail input
	if err := db.Create(detailInput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create detail input", "data": err})
	}
	// return the created detail input
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Detail input has created", "data": detailInput})
}

// get all detail inputs from db
func GetAllDetailInputs(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInputs := []model.DetailInput{}
	// find all detail inputs in the database
	if err := db.Find(&detailInputs).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail inputs not found", "data": nil})
	}
	// if no detail input found, return an error
	if len(detailInputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail inputs not found", "data": nil})
	}
	responseDetailInputs := []model.DetailInput{}
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
		responseDetailInputs = append(responseDetailInputs, detailInput)
	}
	// return detail inputs
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail inputs Found", "data": responseDetailInputs})
}

// get single detail input from db
func GetSingleDetailInput(c *fiber.Ctx) error {
	detailInput := new(model.DetailInput)
	account := new(model.Account)
	input := new(model.Input)
	// get id params
	id := c.Params("id")
	// find single detail input in the database by id
	if err := findDetailInputById(id, detailInput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail input not found"})
	}
	// find account in the database by id
	if err := FindAccountById(fmt.Sprint(detailInput.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// find input in the database by id
	if err := FindInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
	}
	// assign account to detail input
	detailInput.Account = *account
	// assign input to detail input
	detailInput.Input = *input
	// return detail input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail input Found", "data": detailInput})
}

// update a detail input in db
func UpdateDetailInput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInput := new(model.DetailInput)
	account := new(model.Account)
	input := new(model.Input)
	// get id params
	id := c.Params("id")
	// find single detail input in the database by id
	if err := findDetailInputById(id, detailInput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail input not found"})
	}
	// store the body in the detail input and return error if encountereds
	if err := c.BodyParser(detailInput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find account in the database by id
	if err := FindAccountById(fmt.Sprint(detailInput.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// find input in the database by id
	if err := FindInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
	}
	// assign account to detail input
	detailInput.Account = *account
	// assign input to detail input
	detailInput.Input = *input
	// update detail input
	if err := db.Save(detailInput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update detail input", "data": err})
	}
	// return the updated detail input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail input has updated", "data": detailInput})
}

// delete a detail input in db
func DeleteDetailInput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInput := new(model.DetailInput)
	// get id params
	id := c.Params("id")
	// find single detail input in the database by id
	if err := findDetailInputById(id, detailInput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail input not found"})
	}
	// delete detail input
	if err := db.Delete(detailInput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete detail input", "data": err})
	}
	// return the deleted detail input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail input has deleted", "data": detailInput})
}

// create multiple detail inputs
func CreateMultipleDetailInputs(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInputs := new([]model.DetailInput)
	// store the body in the detail input and return error if encountereds
	if err := c.BodyParser(detailInputs); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	for _, detailInput := range *detailInputs {
		input := new(model.Input)
		account := new(model.Account)
		// find input in the database by id
		if err := FindInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
		}
		// find account in the database by id
		if err := FindAccountById(fmt.Sprint(detailInput.IdAccount), account); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign input to detail input
		detailInput.Input = *input
		// assign account to detail input
		detailInput.Account = *account
		// create detail input
		if err := db.Create(&detailInput).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create detail input", "data": err})
		}
	}
	// return the created detail input
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Detail input has created", "data": detailInputs})
}

// update multiple detail inputs
func UpdateMultipleDetailInputs(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInputs := new([]model.DetailInput)
	existingDetailInputs := []model.DetailInput{}
	id := c.Params("id")
	// store the body in the detail input and return error if encountereds
	if err := c.BodyParser(detailInputs); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find all detail inputs in the database by id input
	if err := db.Find(&existingDetailInputs, "id_input = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail inputs not found", "data": nil})
	}
	db.Delete(&existingDetailInputs)
	for _, detailInput := range *detailInputs {
		input := new(model.Input)
		account := new(model.Account)
		// find input in the database by id
		if err := FindInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
		}
		// find account in the database by id
		if err := FindAccountById(fmt.Sprint(detailInput.IdAccount), account); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign input to detail input
		detailInput.Input = *input
		// assign account to detail input
		detailInput.Account = *account
		// update detail input
		if err := db.Save(&detailInput).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update detail input", "data": err})
		}
	}
	// return the updated detail input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail input has updated", "data": detailInputs})
}

// get all detail inputs by id input
func GetDetailInputsByInputID(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInputs := []model.DetailInput{}
	// get id params
	id := c.Params("id")
	// find all detail inputs in the database by id input
	if err := db.Find(&detailInputs, "id_input = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail inputs not found", "data": nil})
	}
	// if no detail input found, return an error
	if len(detailInputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail inputs not found", "data": nil})
	}
	responseDetailInputs := []model.DetailInput{}
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
		responseDetailInputs = append(responseDetailInputs, detailInput)
	}
	// return detail inputs
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail inputs Found", "data": responseDetailInputs})
}
