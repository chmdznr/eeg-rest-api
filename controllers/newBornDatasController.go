package controllers

import (
	"be-eeg/database"
	"be-eeg/models"
	"be-eeg/models/reqresp"
	"github.com/gofiber/fiber/v2"
)

// CreateNewbornData creates a new record in the newborn_datas table
// @Summary Create a new record in the newborn_datas table
// @Description Create a new record in the newborn_datas table
// @Tags Bayi
// @Accept json
// @Produce json
// @Param newbornData body models.NewbornData true "Data Sensor pada bayi"
// @Success 201 {object} reqresp.SuccessResponse
// @Failure 400 {object} reqresp.ErrorResponse
// @Failure 500 {object} reqresp.ErrorResponse
// @Router /newborn-data [post]
func CreateNewbornData(c *fiber.Ctx) error {
	newbornData := new(models.NewbornData)
	if err := c.BodyParser(newbornData); err != nil {
		return c.Status(400).JSON(&reqresp.ErrorResponse{
			Status:  "error",
			Message: "Failed to parse request body: " + err.Error(),
		})
	}

	if err := database.DB.Create(&newbornData).Error; err != nil {
		return c.Status(500).JSON(&reqresp.ErrorResponse{
			Status:  "error",
			Message: "Failed to create newborn data: " + err.Error(),
		})
	}

	return c.Status(201).JSON(&reqresp.SuccessResponse{
		Status:  "success",
		Message: "Berhasil menambahkan data bayi baru",
		Data:    newbornData,
	})
}
