package models

import (
	"time"

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

type Poll struct {
	gorm.Model
	Name string
}

type Question struct {
	gorm.Model
	MainText string
	Position int
	PollID   uint
}

type Answer struct {
	gorm.Model
	QuestionID uint
	MainText   string
}

type Instance struct {
	gorm.Model
	PollID uint
}

type Reply struct {
	InstanceID uint
	QuestionID uint
	// answer_id or main_text in reply
	AnswerID  uint
	MainText  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
