package models

import (
	"gorm.io/gorm"
)

type Poll struct {
	gorm.Model
	Name    string
	OwnerCn string
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
	PollID  uint
	OwnerCn string
}

type Reply struct {
	gorm.Model
	InstanceID uint
	QuestionID uint
	// answer_id or main_text in reply
	AnswerID uint
	MainText string
}
