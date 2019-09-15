package routers

import (
	"Icarus/config"
	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.Engine,conf *config.Config) {
	router := gin.Default()
	user := router.Group("/user")
	{
		user.POST("/login")
	}
	excel := router.Group("/excel")
	{

	}
}
