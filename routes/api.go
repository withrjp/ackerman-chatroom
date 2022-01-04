package routes

import (
	"ackerman-chatroom/app/controllers/app"
	"ackerman-chatroom/app/middleware"
	"ackerman-chatroom/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/user/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.GET("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
	}
}
