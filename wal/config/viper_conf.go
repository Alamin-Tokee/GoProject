package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// ReadConfig default
func ReadConfig() {
	// default ==> setx APP_ENVIRONMENT STAGING
	//Or in main.go init func set os.Setenv("APP_ENVIRONMENT", "STAGING")
	//fmt.Println("os.Getenv", os.Getenv)

	if os.Getenv("APP_ENVIRONMENT") == "STAGING" {
		viper.SetConfigName("properties-staging")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./resource")
		//viper.SetConfigFile("/resource/properties-staging")
	} else if os.Getenv("APP_DEVELOPMENT") == "PROD" {
		viper.SetConfigName("propertise-prod")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./resource")
		//viper.SetConfigFile("../resource")
	}

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Err viper_config")
		fmt.Println(err)
	}
}
