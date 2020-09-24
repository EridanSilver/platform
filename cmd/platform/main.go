package main

import (
	"context"

	"github.com/EridanSilver/platform/app"
)

func main() {
	ctx := context.Background()

	application := app.NewApp()
	if err := application.Run(ctx); err != nil {
		panic(err)
	}
}
