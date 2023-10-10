package handler

import (
	"fmt"

	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find account by id
func findAccountById(id string, account *model.Account) error {
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

// get all Accounts from db
func GetAllAccounts(c *fiber.Ctx) error {
	db := database.DB.Db
	accounts := []model.Account{}
	// find all accounts in the database
	db.Find(accounts)
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
		account.Category = *category
		responseAccounts = append(responseAccounts, account)
	}
	// return accounts
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Accounts Found", "data": responseAccounts})
}

// get singleAccount from db
func GetSingleAccount(c *fiber.Ctx) error {
	account := new(model.Account)
	category := new(model.Category)
	// get id params
	id := c.Params("id")
	// find single account in the database by id
	if err := findAccountById(id, account); err != nil {
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
	if err := findAccountById(id, account); err != nil {
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

// delete account
func DeleteAccount(c *fiber.Ctx) error {
	db := database.DB.Db
	account := new(model.Account)
	// get id params
	id := c.Params("id")
	// find single account in the database by id
	if err := findAccountById(id, account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// delete account
	if err := db.Delete(account).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete account", "data": err})
	}
	// return deleted account
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Account deleted"})
}
