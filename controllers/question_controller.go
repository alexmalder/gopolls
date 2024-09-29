package controllers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /questions
// Create new questions
func PostQuestion(c *gin.Context) {
	var input models.Question
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

// PUT /questions
// Update a questions
func PutQuestion(c *gin.Context) {
	// Get model if exist
	var question models.Question
	if err := models.DB.Where("id = ?", c.Query("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&question).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": question})
}
