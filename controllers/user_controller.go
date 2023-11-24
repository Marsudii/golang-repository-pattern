// controllers/user_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/request"
	"go-fiber/services"
)

// CLASS
type UserController struct {
	userService services.UserService
}

func NewUserController(userService *services.UserService) UserController {
	return UserController{
		userService: *userService,
	}
}

// GET ALL DATA
func (uc *UserController) GetAll(c *fiber.Ctx) error {
	users, _ := uc.userService.GetAllData()

	return c.JSON(fiber.Map{
		"users": users,
	})
}

// GET DETAIL DATA
func (us *UserController) Show(c *fiber.Ctx) error {

	id := c.Params("id")
	userDetail, _ := us.userService.GetDetailData(id)

	return c.JSON(fiber.Map{
		"data": userDetail,
	})
}

// STORE DATA
func (uc *UserController) Create(c *fiber.Ctx) error {
	userRequest := new(request.UserRequest)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	errValidation := uc.userService.Validation(*userRequest)
	if errValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": errValidation.Error(),
		})
	}

	user, errNew := uc.userService.New(*userRequest)
	if errNew != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"user":    user,
	})
}
