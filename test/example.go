package main

import (
	"minigo"
)

func main() {
	router := minigo.Default()

	router.Post("/hey", func(ctx *minigo.Context) {
		ctx.JSON("Hi~")
	})

	_ = router.Run(":9527")
}