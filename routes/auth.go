package routes

import (
	"net/http"

	"codedolphin.io/users-service/models"
	"github.com/gofiber/fiber/v2"
)

func initAuthRoutes(router *fiber.Router) {
	mainRouter := *router
	authRouter := mainRouter.Group("/users/auth")
	authRouter.Post("/register", register)
	authRouter.Post("/login", login)
}

func register(c *fiber.Ctx) error {
	var userDto models.UserDTO
	if err := c.BodyParser(&userDto); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	userDto.HashPassword()
	user := userDto.ToUser()
	res := models.DB.Create(user)
	if res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}
	return c.JSON(user)
}

func login(c *fiber.Ctx) error {
	var uldto models.UserLoginDTO
	if err := c.BodyParser(&uldto); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var user models.User
	res := models.DB.Where("email = ?", uldto.Email).First(&user)
	if res.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "No user with that email address was found!",
		})
	}
	passCorrect := uldto.VerifyPassword(user.Password)
	if !passCorrect {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "Password incorrect!",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success!",
	})
}
