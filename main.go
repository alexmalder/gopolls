package main

import (
	"fmt"
	"io/ioutil"
	"main/controllers"
	"main/domain"
	"main/models"
)

func main() {
	models.Connect()

	content, err := ioutil.ReadFile("init.sql")
	if err != nil {
		fmt.Println("Err")
	}
	fmt.Println(string(content))

	models.ExecSqlScript(string(content))

	models.Seed()
	domain.ExampleConn_WhoAmI()

	controllers.Listen()
}
