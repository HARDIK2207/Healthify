package main

import (
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	//get api to check working or not
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})

	app.Start()
}
