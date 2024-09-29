package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=alexmalder password=alexmalder dbname=alexmalder port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

func ExecSqlScript(content string) {
	DB.Exec(content)
}

func Seed() {
	poll0 := Poll{Name: "My first poll"}
	DB.Create(&poll0)
	log.Println("Created poll0:", poll0.ID, poll0.Name)

	question0 := Question{MainText: "Fav country", Position: 1, PollID: poll0.ID}
	DB.Create(&question0)
	log.Println("Created question0:", question0.ID, question0.MainText)

	answer0 := Answer{QuestionID: question0.ID, MainText: "Germany"}
	answer1 := Answer{QuestionID: question0.ID, MainText: "Russia"}
	DB.Create(&answer0)
	log.Println("Created answer0:", answer0.ID, answer0.MainText)
	DB.Create(&answer1)
	log.Println("Created answer1:", answer1.ID, answer1.MainText)

	instance0 := Instance{PollID: poll0.ID}
	DB.Create(&instance0)
	log.Println("Created instance0:", instance0.ID, instance0.PollID)

	// reply is Germany
	reply0 := Reply{InstanceID: instance0.ID, QuestionID: question0.ID, AnswerID: answer0.ID}
	// reply is Russia
	reply1 := Reply{InstanceID: instance0.ID, QuestionID: question0.ID, AnswerID: answer1.ID}
	// reply is free form
	reply2 := Reply{InstanceID: instance0.ID, QuestionID: question0.ID, MainText: "Poland"}
	// create replies
	DB.Create(&reply0)
	log.Println("Created reply0:", reply0.QuestionID, reply0.AnswerID, reply0.InstanceID, reply0.MainText)
	DB.Create(&reply1)
	log.Println("Created reply1:", reply1.QuestionID, reply1.AnswerID, reply1.InstanceID, reply1.MainText)
	DB.Create(&reply2)
	log.Println("Created reply2:", reply2.QuestionID, reply2.AnswerID, reply2.InstanceID, reply2.MainText)
}
