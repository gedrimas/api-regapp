package main

import (
	"api-regapp/database"
	//"api-regapp/helpers"
	"github.com/gin-gonic/gin"
	"api-regapp/routes"
	//"github.com/gin-contrib/cors"
	"os"
	"github.com/joho/godotenv"
	"log"
)


func EnvFileVal(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("ERROR loading .env file")
	}
	val := os.Getenv(key)
	if val == "" {
		log.Fatal("ERROR no value in .env file")
	}
	return val
}

func main() {

	router := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	//router.Use(cors.New(config))

	database.Start()

	router.Use(gin.Logger())

	routes.EventRoutes(router)


	port := EnvFileVal("PORT")
	router.Run(":" + port)
}
