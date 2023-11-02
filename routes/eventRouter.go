package routes

import (
	"api-regapp/controllers"
	"api-regapp/middleware"

	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.Engine){
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.AuthenticateUser())
	router.GET("/event", controllers.GetEvent())
	router.POST("/event", controllers.CreateEvent())
}