package handler

import (
	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find general journal by id
func findGeneralJournalById(id string, generalJournal *model.GeneralJournal) error {
	db := database.DB.Db
	// find single general journal in the database by id
	db.Find(&generalJournal, "id = ?", id)
	// if no general journal found, return an error
	if generalJournal.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a general journal
func CreateGeneralJournal(c *fiber.Ctx) error {
	db := database.DB.Db
	generalJournal := new(model.GeneralJournal)
	// store the body in the general journal and return error if encountered
	if err := c.BodyParser(generalJournal); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create general journal
	if err := db.Create(generalJournal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create general journal", "data": err})
	}
	// return the created general journal
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "General Journal has created", "data": generalJournal})
}

// get all General Journals from db
func GetAllGeneralJournals(c *fiber.Ctx) error {
	db := database.DB.Db
	generalJournals := []model.GeneralJournal{}
	// find all general journals in the database
	if err := db.Find(&generalJournals).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "General Journals not found", "data": nil})
	}
	// if no general journal found, return an error
	if len(generalJournals) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "General Journals not found", "data": nil})
	}
	// return general journals
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "General Journals Found", "data": generalJournals})
}

// get single general journal from db
func GetSingleGeneralJournal(c *fiber.Ctx) error {
	generalJournal := new(model.GeneralJournal)
	// get id params
	id := c.Params("id")
	// find single general journal in the database by id
	if err := findGeneralJournalById(id, generalJournal); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "General Journal not found"})
	}
	// return general journal
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "General Journal Found", "data": generalJournal})
}

// update general journal
func UpdateGeneralJournal(c *fiber.Ctx) error {
	db := database.DB.Db
	generalJournal := new(model.GeneralJournal)
	// get id params
	id := c.Params("id")
	// find single general journal in the database by id
	if err := findGeneralJournalById(id, generalJournal); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "General Journal not found"})
	}
	// store the body in the general journal and return error if encountered
	if err := c.BodyParser(generalJournal); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// update general journal
	if err := db.Save(generalJournal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update general journal", "data": err})
	}
	// return the updated general journal
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "General Journal has updated", "data": generalJournal})
}

// delete general journal
func DeleteGeneralJournal(c *fiber.Ctx) error {
	db := database.DB.Db
	generalJournal := new(model.GeneralJournal)
	// get id params
	id := c.Params("id")
	// find single general journal in the database by id
	if err := findGeneralJournalById(id, generalJournal); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "General Journal not found"})
	}
	// delete general journal
	if err := db.Delete(generalJournal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete general journal", "data": err})
	}
	// return the deleted general journal
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "General Journal has deleted", "data": generalJournal})
}
