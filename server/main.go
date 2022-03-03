package main

import (
	"goapp-service/controller"
	log "goapp-service/logger"
)

func main() {
	log.Init()
	var app = controller.App{}
	app.RestHandlerController()
	app.Run()
}
