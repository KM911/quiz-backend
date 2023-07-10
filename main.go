package main

import (
	"quiz/app"
	"quiz/config"
	"quiz/model"
)

func main() {
	config.LoadConfig()
	model.InitDao()
	app.Start()
}
