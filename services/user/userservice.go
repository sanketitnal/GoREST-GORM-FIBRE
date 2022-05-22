package user

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sanketitnal/gobasicrest/database"
	"github.com/sanketitnal/gobasicrest/models"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/user/getUserById/:id", getUserById)
	app.Get("/user/getAll", getAllUsers)
	app.Post("/user/create", createUser)
	app.Delete("/user/delete/:id", deleteUserById)
	app.Put("/user/update", updateUserById)
}

func getUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("/user/delete/:id \n id should be integer")
	}
	user := models.User{}
	result := database.PostgresDB.First(&user, "uid = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	} else if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(200).JSON(user)
}

func getAllUsers(c *fiber.Ctx) error {
	users := []models.User{}
	result := database.PostgresDB.Find(&users)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	} else if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(200).JSON(users)
}

func createUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	result := database.PostgresDB.Create(&user)
	if result.Error != nil {
		return c.Status(400).SendString(result.Error.Error())
	}
	return c.SendStatus(200)
}

func deleteUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("/user/delete/:id \n id should be integer")
	}
	var user models.User
	if result := database.PostgresDB.First(&user, "uid = ?", id); result.Error != nil {
		return c.SendStatus(404)
	}

	result := database.PostgresDB.Delete(&user, "uid = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	} else if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.SendStatus(200)
}

func updateUserById(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	result := database.PostgresDB.Model(&user).Where("uid", user.Uid).Updates(user)
	if result.Error != nil {
		return c.Status(400).SendString(result.Error.Error())
	}
	return c.SendStatus(200)
}
