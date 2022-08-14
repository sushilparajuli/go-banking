package main

import (
	"github.com/sushilparajuli/go-banking/app"
	"github.com/sushilparajuli/go-banking/logger"
)

func main() {
	// define routes
	logger.Info("Starting the application")
	app.App()
}
