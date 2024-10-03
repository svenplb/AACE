package app

import (
	"example/chessbot/router"
)

func SetupAndRunApp() error {
	router.SetupRoutes()

	return nil

}
