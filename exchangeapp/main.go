package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	// listen and serve on 0.0.0.0:8888(default: 8080)
	r.Run(port)
}
