package service

import (
	"fmt"
	"github.com/KM911/oslib/adt"
	"github.com/gofiber/fiber/v2"
	"quiz/model"
)

func GetChoiceNum(c *fiber.Ctx) error {
	return c.JSON(model.ChoiceNums())
}

func GetChoiceById(c *fiber.Ctx) error {
	id := adt.Str2int(c.Params("id", "1"))
	fmt.Println("id is", id)
	choice := model.GetChoice(int64(id))
	return c.JSON(model.ChoiceModel{
		Title:  choice.Title,
		A:      choice.A,
		B:      choice.B,
		C:      choice.C,
		D:      choice.D,
		Answer: choice.Answer,
	})
}
