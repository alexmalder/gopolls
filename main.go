package main

import (
	"main/auth"
	"main/controllers"
	"main/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.Seed()
	auth.ExampleConn_WhoAmI()

	r.GET("/api/polls", controllers.GetPolls)
	r.Run()
}
