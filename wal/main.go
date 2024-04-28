package main

import (
	"log"
	"os"
	"wal/config"
	"wal/router"
	"wal/util"

	"github.com/spf13/viper"
)

func init() {
	os.Setenv("APP_ENVIRONMENT", "STAGING")

	// read config environment
	config.ReadConfig()

	util.Pool = util.SetupRedisJWT()
}

func main() {
	var err error
	config.DB, err = config.SetupDB()

	if err != nil {
		log.Fatal(err)
	}

	defer config.DB.Close()

	port := viper.GetString("PORT")

	// Setup router
	router := router.NewRoutes()

	log.Fatal(router.Run(":" + port))
}
