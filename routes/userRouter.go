package routes

import (
	"movies_app/controllers"
	"movies_app/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router gin.Engine) {
	router.Use(middleware.AuthenticateUser())
	router.GET("/users/:user_id", controllers.GetUser())
	router.GET("/users", controllers.GetUsers())

}