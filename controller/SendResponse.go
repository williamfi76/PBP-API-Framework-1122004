package controller

import (
	m "fiber-api/model"

	"github.com/gofiber/fiber/v2"
)

func SendSuccessUsersResponse(c *fiber.Ctx, message string, data []m.User) *fiber.Ctx {
	var response m.UsersResponse
	response.Status = 200
	response.Message = message
	response.Data = data
	c.Status(fiber.StatusOK).JSON(response)
	return nil
}

func SendSuccessUserResponse(c *fiber.Ctx, message string, data m.User) *fiber.Ctx {
	var response m.UserResponse
	response.Status = 200
	response.Message = message
	response.Data = data
	c.Status(fiber.StatusOK).JSON(response)
	return nil
}

func SendErrorResponse(c *fiber.Ctx, message string) *fiber.Ctx {
	var response m.UsersResponse
	response.Status = 400
	response.Message = message
	c.Status(fiber.StatusOK).JSON(response)
	return nil
}
