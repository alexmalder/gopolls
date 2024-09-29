package controllers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /replies
// Create new replies
func PostReply(c *gin.Context) {
	var input models.Reply
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

// PUT /replies
// Update a replies
func PutReply(c *gin.Context) {
	// Get model if exist
	var reply models.Reply
	if err := models.DB.Where("id = ?", c.Query("id")).First(&reply).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Reply
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&reply).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": reply})
}
