package main

import (
	"github.com/AmeyCC/go-task/api/app"
	"github.com/AmeyCC/go-task/api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
