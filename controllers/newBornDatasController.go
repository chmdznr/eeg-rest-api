package controllers

import (
	"be-eeg/database"
	"be-eeg/models"
	"be-eeg/models/reqresp"
	"github.com/gofiber/fiber/v2"
)

// CreateNewbornData creates a new record in the newborn_datas table
func CreateNewbornData(c *fiber.Ctx) error {
	newbornData := new(models.NewbornData)
	if err := c.BodyParser(newbornData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := database.DB.Create(&newbornData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Cannot create record"})
	}

	return c.Status(201).JSON(&reqresp.SuccessResponse{
		Status:  "success",
		Message: "Berhasil menambahkan data bayi baru",
		Data:    newbornData,
	})
}

// GetNewbornData returns all records from the newborn_datas table
func GetNewbornData(c *fiber.Ctx) error {
	var newbornData []models.NewbornData
	if err := database.DB.Find(&newbornData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Cannot retrieve records"})
	}

	return c.JSON(&reqresp.SuccessResponse{
		Status:  "success",
		Message: "Berhasil mengambil data",
		Data:    newbornData,
	})
}
