package main

import (
	"minigo"
)

func main() {
	api := minigo.Default()

	api.Post("/hey", func(ctx *minigo.Context) {
		ctx.JSON("Hi~")
	})

	_ = api.Run(":9527")
}