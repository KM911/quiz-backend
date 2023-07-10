package main

import (
	"demo/app"
	"demo/config"
)

func main() {
	config.LoadConfig()

	app.Start()

}
