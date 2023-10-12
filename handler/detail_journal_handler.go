package handler

import (
	"fmt"

	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find detail journal by id
func findDetailJournalById(id string, detailJournal *model.DetailJournal) error {
	db := database.DB.Db
	// find single detail journal in the database by id
	db.Find(&detailJournal, "id = ?", id)
	// if no detail journal found, return an error
	if detailJournal.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a detail journal
func CreateDetailJournal(c *fiber.Ctx) error {
	db := database.DB.Db
	detailJournal := new(model.DetailJournal)
	journal := new(model.GeneralJournal)
	account := new(model.Account)
	// store the body in the detail journal and return error if encountered
	if err := c.BodyParser(detailJournal); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find journal in the database by id
	if err := findGeneralJournalById(fmt.Sprint(detailJournal.IdGeneralJournal), journal); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Journal not found"})
	}
	// find account in the database by id
	if err := findAccountById(fmt.Sprint(detailJournal.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// assign journal to detail journal
	detailJournal.GeneralJournal = *journal
	// assign account to detail journal
	detailJournal.Account = *account
	// create detail journal
	if err := db.Create(detailJournal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create detail journal", "data": err})
	}
	// return the created detail journal
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Detail journal has created", "data": detailJournal})
}

// get all detail journals from db
func GetAllDetailJournals(c *fiber.Ctx) error {
	db := database.DB.Db
	detailJournals := []model.DetailJournal{}
	// find all detail journals in the database
	if err := db.Find(&detailJournals).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail journals not found", "data": nil})
	}
	// if no detail journal found, return an error
	if len(detailJournals) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail journals not found", "data": nil})
	}
	responseDetailJournals := []model.DetailJournal{}
	for _, detailJournal := range detailJournals {
		account := new(model.Account)
		generalJournal := new(model.GeneralJournal)
		// convert id to string
		idAccount := fmt.Sprint(detailJournal.IdAccount)
		idGeneralJournal := fmt.Sprint(detailJournal.IdGeneralJournal)
		// find account in the database by id
		if err := findAccountById(idAccount, account); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// find general journal in the database by id
		if err := findGeneralJournalById(idGeneralJournal, generalJournal); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "General journal not found"})
		}
		// assign account to detail journal
		detailJournal.Account = *account
		// assign general journal to detail journal
		detailJournal.GeneralJournal = *generalJournal
		responseDetailJournals = append(responseDetailJournals, detailJournal)
	}
	// return detail journals
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail journals Found", "data": detailJournals})
}

// get single detail journal from db
func GetSingleDetailJournal(c *fiber.Ctx) error {
	detailJournal := new(model.DetailJournal)
	account := new(model.Account)
	generalJournal := new(model.GeneralJournal)
	// get id params
	id := c.Params("id")
	// find single detail journal in the database by id
	if err := findDetailJournalById(id, detailJournal); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail journal not found"})
	}
	// find account in the database by id
	if err := findAccountById(fmt.Sprint(detailJournal.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// find general journal in the database by id
	if err := findGeneralJournalById(fmt.Sprint(detailJournal.IdGeneralJournal), generalJournal); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "General journal not found"})
	}
	// assign account to detail journal
	detailJournal.Account = *account
	// assign general journal to detail journal
	detailJournal.GeneralJournal = *generalJournal
	// return detail journal
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail journal Found", "data": detailJournal})
}

// update a detail journal in db
func UpdateDetailJournal(c *fiber.Ctx) error {
	db := database.DB.Db
	detailJournal := new(model.DetailJournal)
	account := new(model.Account)
	generalJournal := new(model.GeneralJournal)
	// get id params
	id := c.Params("id")
	// find single detail journal in the database by id
	if err := findDetailJournalById(id, detailJournal); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail journal not found"})
	}
	// store the body in the detail journal and return error if encountered
	if err := c.BodyParser(detailJournal); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find account in the database by id
	if err := findAccountById(fmt.Sprint(detailJournal.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// find general journal in the database by id
	if err := findGeneralJournalById(fmt.Sprint(detailJournal.IdGeneralJournal), generalJournal); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "General journal not found"})
	}
	// assign account to detail journal
	detailJournal.Account = *account
	// assign general journal to detail journal
	detailJournal.GeneralJournal = *generalJournal
	// update detail journal
	if err := db.Save(detailJournal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update detail journal", "data": err})
	}
	// return the updated detail journal
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail journal has updated", "data": detailJournal})
}

// delete a detail journal in db
func DeleteDetailJournal(c *fiber.Ctx) error {
	db := database.DB.Db
	detailJournal := new(model.DetailJournal)
	// get id params
	id := c.Params("id")
	// find single detail journal in the database by id
	if err := findDetailJournalById(id, detailJournal); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail journal not found"})
	}
	// delete detail journal
	if err := db.Delete(detailJournal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete detail journal", "data": err})
	}
	// return the deleted detail journal
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail journal has deleted", "data": detailJournal})
}
