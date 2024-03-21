package controller

import (
	m "fiber-api/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := connectGorm()
	var users []m.User

	if db.Find(&users).Error != nil {
		SendErrorResponse(c, "Gagal mencari Data")
		return nil
	}

	SendSuccessUsersResponse(c, "Data berhasil ditemukan", users)
	return nil
}

func GetUserByID(c *fiber.Ctx) error {
	db := connectGorm()
	id := c.Params("id")
	var user m.User

	if db.Find(&user, id).Error != nil {
		SendErrorResponse(c, "Gagal mencari Data")
		return nil
	}

	SendSuccessUserResponse(c, "Data berhasil ditemukan", user)
	return nil
}

func InsertNewUser(c *fiber.Ctx) error {
	db := connectGorm()

	var user m.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":    err.Error(),
			"keterangan": "Gagal mengambil data dari input",
		})
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":    err.Error(),
			"keterangan": "Gagal memasukan data ke DB",
		})
	}

	SendSuccessUserResponse(c, "Data berhasil ditambahkan", user)
	return nil
}

func UpdateUser(c *fiber.Ctx) error {
	db := connectGorm()
	id := c.Params("id")

	var user m.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if db.Table("users").Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak dapat mengupdate data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	db := connectGorm()
	id := c.Params("id")

	var user m.User
	db.Table("users").Find(&user, id)
	if db.Table("users").Delete(&user, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data",
		})
	}

	SendSuccessUserResponse(c, "Data berhasil dihapus", user)
	return nil
}
