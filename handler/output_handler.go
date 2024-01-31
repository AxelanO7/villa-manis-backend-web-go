package handler

import (
	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find output by id
func FindOutputById(id string, output *model.Output) error {
	db := database.DB.Db
	// find single output in the database by id
	db.Find(&output, "id = ?", id)
	// if no output found, return an error
	if output.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a output
func CreateOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	output := new(model.Output)
	// store the body in the output and return error if encountered
	if err := c.BodyParser(output); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your output", "data": err})
	}
	// create output
	if err := db.Create(output).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create output", "data": err})
	}
	// return the created output
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Output has created", "data": output})
}

// get all Outputs from db
func GetAllOutputs(c *fiber.Ctx) error {
	db := database.DB.Db
	outputs := []model.Output{}
	// find all outputs in the database
	if err := db.Find(&outputs).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Outputs not found", "data": nil})
	}
	// if no output found, return an error
	if len(outputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Outputs not found", "data": nil})
	}
	// return outputs
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Outputs Found", "data": outputs})
}

// get single output from db
func GetSingleOutput(c *fiber.Ctx) error {
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find single output in the database by id
	if err := FindOutputById(id, output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// return output
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Output Found", "data": output})
}

// update a output in db
func UpdateOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find single output in the database by id
	if err := FindOutputById(id, output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// store the body in the output and return error if encountered
	if err := c.BodyParser(output); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your output", "data": err})
	}
	// update output
	if err := db.Save(output).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update output", "data": err})
	}
	// return the updated output
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Output has updated", "data": output})
}

// delete a output in db
func DeleteOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find single output in the database by id
	if err := FindOutputById(id, output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// find detail output in the database by id
	detailOutput := []model.DetailOutput{}
	db.Find(&detailOutput, "output_id = ?", id)
	// if err := db.Find(&detailOutput, "output_id = ?", id).Error; err != nil {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Output not found", "data": nil})
	// }
	// delete detail output
	db.Delete(&detailOutput)
	// if err := db.Delete(&detailOutput).Error; err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete detail output", "data": err})
	// }
	// delete output
	if err := db.Delete(output).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete output", "data": err})
	}
	// return deleted output
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Output has deleted", "data": nil})
}

// get output master
func GetOutputMaster(c *fiber.Ctx) error {
	db := database.DB.Db
	outputs := []model.Output{}
	detailOutputs := []model.DetailOutput{}
	type OutputMaster struct {
		ID                uint     `json:"id"`
		NoOutput          string   `json:"no_output"`
		DateOutput        string   `json:"date_output"`
		StatusOutput      string   `json:"status_output"`
		OutputInformation []string `json:"output_information"`
	}
	outputMasters := []OutputMaster{}
	// find all outputs in the database
	if err := db.Find(&outputs).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Outputs not found", "data": nil})
	}
	if err := db.Find(&detailOutputs).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail Outputs not found", "data": nil})
	}
	for _, output := range outputs {
		OutputInformation := []string{}
		for _, detailOutput := range detailOutputs {
			if output.ID == uint(detailOutput.IdOutput) {
				OutputInformation = append(OutputInformation, detailOutput.OutputInformation)
			}
		}
		outputMasters = append(outputMasters, OutputMaster{
			ID:                output.ID,
			NoOutput:          output.NoOutput,
			DateOutput:        output.DateOutput,
			StatusOutput:      output.StatusOutput,
			OutputInformation: OutputInformation,
		})
	}
	// if no output found, return an error
	if len(outputMasters) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Outputs not found", "data": nil})
	}
	// return outputs
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Outputs Found", "data": outputMasters})
}
