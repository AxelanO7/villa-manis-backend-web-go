package handler

import (
	"github.com/AxelanO7/villa-manis-backend-web-go/database"
	"github.com/AxelanO7/villa-manis-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find category by id
func FindCategoryByID(id string, category *model.Category) error {
	db := database.DB.Db
	// find category in the database by id
	db.Find(&category, "id = ?", id)
	// if no category found, return an error
	if category.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a category
func CreateCategory(c *fiber.Ctx) error {
	db := database.DB.Db
	category := new(model.Category)
	// store the body in the category and return error if encountered
	if err := c.BodyParser(category); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create category
	if err := db.Create(category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create category", "data": err})
	}
	// return the created category
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Category has created", "data": category})
}

// get all Category from db
func GetAllCategories(c *fiber.Ctx) error {
	db := database.DB.Db
	categories := []model.Category{}
	// find all categories in the database
	db.Find(categories)
	// if no category found, return an error
	if len(categories) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Categories not found", "data": nil})
	}
	// return categories
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Categories Found", "data": categories})
}

// get single category from db
func GetSingleCategory(c *fiber.Ctx) error {
	category := new(model.Category)
	// get id params
	id := c.Params("id")
	// find single category in the database by id
	if err := FindCategoryByID(id, category); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
	}
	// return category
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Category Found", "data": category})
}

// update a category in db
func UpdateCategory(c *fiber.Ctx) error {
	db := database.DB.Db
	category := new(model.Category)
	// get id params
	id := c.Params("id")
	// find single category in the database by id
	if err := FindCategoryByID(id, category); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
	}
	// store the body in the category and return error if encountered
	if err := c.BodyParser(category); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// update category
	if err := db.Save(category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update category", "data": err})
	}
	// return the updated category
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Category has updated", "data": category})
}

// delete a category
func DeleteCategory(c *fiber.Ctx) error {
	db := database.DB.Db
	category := new(model.Category)
	// get id params
	id := c.Params("id")
	// find single category in the database by id
	if err := FindCategoryByID(id, category); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
	}
	// delete category
	if err := db.Delete(category, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete category", "data": err})
	}
	// return deleted category
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Category has deleted", "data": category})
}
