package main

import (
	"github.com/broboredo/locapp-api/config"
	"github.com/broboredo/locapp-api/router"
)

var (
	logger config.Logger
)

// @contact.name   Bruno Roboredo
// @contact.email  roboredo.bruno@gmail.com
func main() {
	logger := config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("init main error: %v", err)
		return
	}

	router.Init()
}
