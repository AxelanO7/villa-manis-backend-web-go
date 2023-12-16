package handler

import (
	"fmt"

	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find cart by id
func findGeneralCartById(id string, cart *model.GeneralCart) error {
	db := database.DB.Db
	// find single cart in the database by id
	db.Find(&cart, "id = ?", id)
	// if no cart found, return an error
	if cart.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a cart
func CreateGeneralCart(c *fiber.Ctx) error {
	db := database.DB.Db
	account := new(model.Account)
	cart := new(model.GeneralCart)
	// store the body in the cart and return error if encountereds
	if err := c.BodyParser(cart); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find account in the database by id
	if err := FindAccountById(fmt.Sprint(cart.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// assign account to cart
	cart.Account = *account
	// create cart
	if err := db.Create(cart).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create cart", "data": err})
	}
	// return the created cart
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Cart has created", "data": cart})
}

// get all Carts from db
func GetAllGeneralCarts(c *fiber.Ctx) error {
	db := database.DB.Db
	carts := []model.GeneralCart{}
	// find all carts in the database
	if err := db.Find(&carts).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Carts not found", "data": nil})
	}
	// if no cart found, return an error
	if len(carts) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Carts not found", "data": nil})
	}
	responseCarts := []model.GeneralCart{}
	for _, cart := range carts {
		account := new(model.Account)
		// find account in the database by id
		if err := FindAccountById(fmt.Sprint(cart.IdAccount), account); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// assign account to cart
		cart.Account = *account
		responseCarts = append(responseCarts, cart)
	}
	// return carts
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Carts Found", "data": responseCarts})
}

// get single cart from db
func GetSingleGeneralCart(c *fiber.Ctx) error {
	account := new(model.Account)
	cart := new(model.GeneralCart)
	// get id params
	id := c.Params("id")
	// find single cart in the database by id
	if err := findGeneralCartById(id, cart); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Cart not found"})
	}
	// find account in the database by id
	if err := FindAccountById(fmt.Sprint(cart.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// assign account to cart
	cart.Account = *account
	// return cart
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Cart Found", "data": cart})
}

// update a cart in db
func UpdateGeneralCart(c *fiber.Ctx) error {
	db := database.DB.Db
	account := new(model.Account)
	cart := new(model.GeneralCart)
	// get id params
	id := c.Params("id")
	// find single cart in the database by id
	if err := findGeneralCartById(id, cart); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Cart not found"})
	}
	// store the body in the cart and return error if encountereds
	if err := c.BodyParser(cart); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find account in the database by id
	if err := FindAccountById(fmt.Sprint(cart.IdAccount), account); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}
	// assign account to cart
	cart.Account = *account
	// update cart
	if err := db.Save(cart).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update cart", "data": err})
	}
	// return the updated cart
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Cart has updated", "data": cart})
}

// delete a cart in db
func DeleteGeneralCart(c *fiber.Ctx) error {
	db := database.DB.Db
	cart := new(model.GeneralCart)
	// get id params
	id := c.Params("id")
	// find single cart in the database by id
	if err := findGeneralCartById(id, cart); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Cart not found"})
	}
	// delete cart
	if err := db.Delete(cart).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete cart", "data": err})
	}
	// return the deleted cart
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Cart has deleted", "data": nil})
}
