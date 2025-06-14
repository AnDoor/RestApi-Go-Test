package main

import (
	"API_GO/config"
	"API_GO/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")
	//inicializar configuracion
	err := config.Init()

	if err != nil {
		logger.Error("config initialization error: ", err)
		return
	}
	//inicializar router
	router.Initialize()
}
