package controllers

import (
	"net/http"

	"main/domain"
	"main/models"

	"github.com/gin-gonic/gin"
)

// GET /polls
// Get all polls
func GetPolls(c *gin.Context) {
	var polls []models.Poll
	models.DB.Find(&polls)
	c.JSON(http.StatusOK, gin.H{"data": polls})
}

// GET /polls/get_by_owner_cn
// Find a poll
func GetPollsByOwnerCn(c *gin.Context) {
	var poll models.Poll

	if err := models.DB.Where("owner_cn = ?", c.Query("ownercn")).First(&poll).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": poll})
}

// POST /polls
// Create new poll
func PostPoll(c *gin.Context) {
	var input models.Poll
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entry, err := domain.FindByCn(input.OwnerCn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input, "entry": entry})
}

// PUT /polls
// Update a polls
func PutPoll(c *gin.Context) {
	// Get model if exist
	var poll models.Poll
	if err := models.DB.Where("id = ?", c.Query("id")).First(&poll).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Poll
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&poll).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": poll})
}

// DELETE /polls
// Delete a poll
func DeletePoll(c *gin.Context) {
	// Get model if exist
	var poll models.Poll
	if err := models.DB.Where("id = ?", c.Query("id")).First(&poll).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&poll)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
