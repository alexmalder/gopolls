package controllers

import (
	"main/domain"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /instances
// Create new instances
func PostInstance(c *gin.Context) {
	var input models.Instance
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

// PUT /instances
// Update a instances
func PutInstance(c *gin.Context) {
	// Get model if exist
	var instance models.Instance
	if err := models.DB.Where("id = ?", c.Query("id")).First(&instance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Instance
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&instance).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": instance})
}
