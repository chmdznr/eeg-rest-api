package controllers

import (
	"be-eeg/database"
	"be-eeg/models"
	"be-eeg/models/reqresp"
	"github.com/gofiber/fiber/v2"
)

// CreateNewbornEEG creates a new record in the newborn_eegs table
// @Summary Create a new record in the newborn_eegs table
// @Description Create a new record in the newborn_eegs table
// @Tags Bayi
// @Accept json
// @Produce json
// @Param newbornEEG body models.NewbornEEG true "Data EEG pada bayi"
// @Success 201 {object} reqresp.SuccessResponse
// @Failure 400 {object} reqresp.ErrorResponse
// @Failure 500 {object} reqresp.ErrorResponse
// @Router /newborn-eeg [post]
func CreateNewbornEEG(c *fiber.Ctx) error {
	newbornEEG := new(models.NewbornEEG)
	if err := c.BodyParser(newbornEEG); err != nil {
		return c.Status(400).JSON(&reqresp.ErrorResponse{
			Status:  "error",
			Message: "Failed to parse request body: " + err.Error(),
		})
	}

	if err := database.DB.Create(&newbornEEG).Error; err != nil {
		return c.Status(500).JSON(&reqresp.ErrorResponse{
			Status:  "error",
			Message: "Failed to create newborn EEG data: " + err.Error(),
		})
	}

	return c.Status(201).JSON(&reqresp.SuccessResponse{
		Status:  "success",
		Message: "Berhasil menambahkan data EEG baru untuk bayi",
		Data:    newbornEEG,
	})
}
