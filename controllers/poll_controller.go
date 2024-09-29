package controllers

import (
	"net/http"

	"main/domain"
	"main/models"

	"github.com/gin-gonic/gin"
)

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

// GET /polls/:owner
// Find a poll
func GetPollByOwnerCn(c *gin.Context) { // Get model if exist
	var poll models.Poll

	if err := models.DB.Where("owner_cn = ?", c.Param("ownercn")).First(&poll).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": poll})
}
