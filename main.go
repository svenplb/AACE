package main

import (
	"example/chessbot/app"
)

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
