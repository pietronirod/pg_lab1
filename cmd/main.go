package main

import (
	"log"

	"github.com/pietronirod/lab1/internal/framework/router"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, relying on system environment variables.")
	}
	if !viper.IsSet("WEATHER_API_KEY") {
		log.Fatal("WEATHER_API_KEY is not set.")
	}

	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
