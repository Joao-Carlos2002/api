package router

import (
	"github.com/Joao-Carlos2002/Api-User/router/routes"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/user/:id", routes.GetAccount)
	router.POST("/user", routes.CreateAccount)
	router.PUT("/user", routes.UpdateAccount)
	router.Run(":8000")
}
