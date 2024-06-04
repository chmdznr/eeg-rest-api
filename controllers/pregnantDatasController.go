package controllers

import (
	"be-eeg/database"
	"be-eeg/models"
	"be-eeg/models/reqresp"
	"github.com/gofiber/fiber/v2"
)

// CreatePregnantData creates a new record in the pregnant_datas table
// @Summary Create a new record in the pregnant_datas table
// @Description Create a new record in the pregnant_datas table
// @Tags Ibu hamil
// @Accept json
// @Produce json
// @Param pregnantData body models.PregnantData true "Data Sensor pada ibu hamil"
// @Success 201 {object} reqresp.SuccessResponse
// @Failure 400 {object} reqresp.ErrorResponse
// @Failure 500 {object} reqresp.ErrorResponse
// @Router /pregnant-data [post]
func CreatePregnantData(c *fiber.Ctx) error {
	pregnantData := new(models.PregnantData)
	if err := c.BodyParser(pregnantData); err != nil {
		return c.Status(400).JSON(&reqresp.ErrorResponse{
			Status:  "error",
			Message: "Failed to parse request body: " + err.Error(),
		})
	}

	if err := database.DB.Create(&pregnantData).Error; err != nil {
		return c.Status(500).JSON(&reqresp.ErrorResponse{
			Status:  "error",
			Message: "Failed to create pregnant data: " + err.Error(),
		})
	}

	return c.Status(201).JSON(&reqresp.SuccessResponse{
		Status:  "success",
		Message: "Berhasil menambahkan data ibu hamil baru",
		Data:    pregnantData,
	})
}
