package controllers

import (
	"be-eeg/database"
	"be-eeg/models"
	"be-eeg/models/reqresp"
	"github.com/gofiber/fiber/v2"
)

// CreatePregnantEEG creates a new record in the pregnant_eegs table
// @Summary Create a new record in the pregnant_eegs table
// @Description Create a new record in the pregnant_eegs table
// @Tags Ibu hamil
// @Accept json
// @Produce json
// @Param pregnantEEG body models.PregnantEEG true "Data EEG pada ibu hamil"
// @Success 201 {object} reqresp.SuccessResponse
// @Failure 400 {object} reqresp.ErrorResponse
// @Failure 500 {object} reqresp.ErrorResponse
// @Router /pregnant-eeg [post]
func CreatePregnantEEG(c *fiber.Ctx) error {
	pregnantEEG := new(models.PregnantEEG)
	if err := c.BodyParser(pregnantEEG); err != nil {
		return c.Status(400).JSON(&reqresp.ErrorResponse{
			Status:  "error",
			Message: "Failed to parse request body: " + err.Error(),
		})
	}

	if err := database.DB.Create(&pregnantEEG).Error; err != nil {
		return c.Status(500).JSON(&reqresp.ErrorResponse{
			Status:  "error",
			Message: "Failed to create pregnant EEG data: " + err.Error(),
		})
	}

	return c.Status(201).JSON(&reqresp.SuccessResponse{
		Status:  "success",
		Message: "Berhasil menambahkan data EEG ibu hamil baru",
		Data:    pregnantEEG,
	})
}
