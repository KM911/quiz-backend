package model

import "gorm.io/gorm"

type Choice struct {
	gorm.Model
	Title  string
	A      string
	B      string
	C      string
	D      string
	Answer string
}

type ChoiceModel struct {
	Title  string
	A      string
	B      string
	C      string
	D      string
	Answer string
}

func InsertChoice(choice Choice) {
	db.Create(&choice)
}
func ChoiceNums() int64 {
	//get the max id
	var choice Choice
	db.Last(&choice)
	return int64(choice.ID)
}

func GetChoice(id int64) Choice {
	var choice Choice
	db.First(&choice, id)
	return choice
}

func GetChoices() []Choice {
	var choices []Choice
	db.Find(&choices)
	return choices
}
