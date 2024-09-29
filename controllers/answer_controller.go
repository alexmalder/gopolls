package controllers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /answers
// Create new answers
func PostAnswer(c *gin.Context) {
	var input models.Answer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		entry, err := domain.FindByCn(input.OwnerCn)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	*/

	models.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// PUT /answers
// Update a answers
func PutAnswer(c *gin.Context) {
	// Get model if exist
	var answer models.Answer
	if err := models.DB.Where("id = ?", c.Query("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Answer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&answer).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": answer})
}
