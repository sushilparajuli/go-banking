package main

import (
	"github.com/spf13/viper"
	"github.com/sushilparajuli/go-banking/app"
	"github.com/sushilparajuli/go-banking/logger"
)

func main() {
	// define routes
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	logger.Info("Starting the application")
	app.App()
}
