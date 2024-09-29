package main

import (
	"fmt"
	"io/ioutil"
	"main/controllers"
	"main/domain"
	"main/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.Connect()

	content, err := ioutil.ReadFile("init.sql")
	if err != nil {
		fmt.Println("Err")
	}
	fmt.Println(string(content))

	models.ExecSqlScript(string(content))

	models.Seed()
	domain.ExampleConn_WhoAmI()

	r.POST("/api/polls", controllers.PostPoll)
	r.GET("/api/polls/:ownercn", controllers.GetPollByOwnerCn)
	r.Run(":8000")
}
