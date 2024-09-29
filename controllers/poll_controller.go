package controllers

import (
	"net/http"

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
