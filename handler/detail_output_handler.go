package handler

import (
	"fmt"

	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find detail output by id
func findDetailOutputById(id string, detailOutput *model.DetailOutput) error {
	db := database.DB.Db
	// find single detail output in the database by id
	db.Find(&detailOutput, "id = ?", id)
	// if no detail output found, return an error
	if detailOutput.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a detail output
func CreateDetailOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutput := new(model.DetailOutput)
	output := new(model.Output)
	account := new(model.Account)
	// store the body in the detail output and return error if encountereds
	if err := c.BodyParser(detailOutput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find output in the database by id
	if err := FindOutputById(fmt.Sprint(detailOutput.IdOutput), output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// find account in the database by id
	if err := FindAccountById(fmt.Sprint(detailOutput.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// assign output to detail output
	detailOutput.Output = *output
	// assign account to detail output
	detailOutput.Account = *account
	// create detail output
	if err := db.Create(detailOutput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create detail output", "data": err})
	}
	// return the created detail output
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Detail output has created", "data": detailOutput})
}

// get all detail outputs from db
func GetAllDetailOutputs(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutputs := []model.DetailOutput{}
	// find all detail outputs in the database
	if err := db.Find(&detailOutputs).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail outputs not found", "data": nil})
	}
	// if no detail output found, return an error
	if len(detailOutputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No detail output found", "data": nil})
	}
	responseDetailOutputs := []model.DetailOutput{}
	for _, detailOutput := range detailOutputs {
		account := new(model.Account)
		output := new(model.Output)
		// get id params
		idAccount := fmt.Sprint(detailOutput.IdAccount)
		idOutput := fmt.Sprint(detailOutput.IdOutput)
		// find account in the database by id
		if err := FindAccountById(idAccount, account); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// find output in the database by id
		if err := FindOutputById(idOutput, output); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
		}
		// assign account to detail output
		detailOutput.Account = *account
		// assign output to detail output
		detailOutput.Output = *output
		responseDetailOutputs = append(responseDetailOutputs, detailOutput)
	}
	// return detail outputs
	return c.JSON(fiber.Map{"status": "success", "message": "Detail outputs found", "data": responseDetailOutputs})
}

// get a detail output by id
func GetSingleDetailOutput(c *fiber.Ctx) error {
	detailOutput := new(model.DetailOutput)
	account := new(model.Account)
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find single detail output in the database by id
	if err := findDetailOutputById(id, detailOutput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail output not found", "data": nil})
	}
	// find account in the database by id
	if err := FindAccountById(fmt.Sprint(detailOutput.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// find output in the database by id
	if err := FindOutputById(fmt.Sprint(detailOutput.IdOutput), output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// assign account to detail output
	detailOutput.Account = *account
	// assign output to detail output
	detailOutput.Output = *output
	// return the detail output
	return c.JSON(fiber.Map{"status": "success", "message": "Detail output found", "data": detailOutput})
}

// update a detail output by id
func UpdateDetailOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutput := new(model.DetailOutput)
	account := new(model.Account)
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find detail output in the database by id
	if err := findDetailOutputById(id, detailOutput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail output not found", "data": nil})
	}
	// store the body in the detail output and return error if encountereds
	if err := c.BodyParser(detailOutput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find account in the database by id
	if err := FindAccountById(fmt.Sprint(detailOutput.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// find output in the database by id
	if err := FindOutputById(fmt.Sprint(detailOutput.IdOutput), output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// assign account to detail output
	detailOutput.Account = *account
	// assign output to detail output
	detailOutput.Output = *output
	// update detail output
	if err := db.Save(&detailOutput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update detail output", "data": err})
	}
	// return the updated detail output
	return c.JSON(fiber.Map{"status": "success", "message": "Detail output successfully updated", "data": detailOutput})
}

// delete a detail output by id
func DeleteDetailOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutput := new(model.DetailOutput)
	// get id params
	id := c.Params("id")
	// find detail output in the database by id
	if err := findDetailOutputById(id, detailOutput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail output not found", "data": nil})
	}
	// delete detail output
	if err := db.Delete(&detailOutput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete detail output", "data": err})
	}
	// return the deleted detail output
	return c.JSON(fiber.Map{"status": "success", "message": "Detail output successfully deleted", "data": detailOutput})
}

// create multiple detail outputs
func CreateMultipleDetailOutputs(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutputs := new([]model.DetailOutput)
	// store the body in the detail outputs and return error if encountereds
	if err := c.BodyParser(detailOutputs); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	for _, detailOutput := range *detailOutputs {
		output := new(model.Output)
		account := new(model.Account)
		// find output in the database by id
		if err := FindOutputById(fmt.Sprint(detailOutput.IdOutput), output); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
		}
		// find account in the database by id
		if err := FindAccountById(fmt.Sprint(detailOutput.IdAccount), account); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign output to detail output
		detailOutput.Output = *output
		// assign account to detail output
		detailOutput.Account = *account
		// create detail output
		if err := db.Create(&detailOutput).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create detail output", "data": err})
		}
	}
	// return the created detail outputs
	return c.JSON(fiber.Map{"status": "success", "message": "Detail outputs successfully created", "data": detailOutputs})
}
