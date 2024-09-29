package main

import (
	"fmt"
	"io/ioutil"
	"main/auth"
	"main/controllers"
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
	auth.ExampleConn_WhoAmI()

	r.GET("/api/polls", controllers.GetPolls)
	r.Run()
}
