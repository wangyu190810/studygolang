
package main

import (
	. "awesomeProject/middleware"
	"awesomeProject/middleware/jwt"
	"github.com/gin-gonic/gin"
	//"gopkg.in/gin-gonic/gin.v1"
	. "awesomeProject/api"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(Logger())
	router.GET("/", IndexApi)

	router.POST("/person", AddPersonApi)
	//
	router.GET("/persons", GetPersonsApi)
	//
	//router.GET("/person/:id", GetPersonApi)
	//
	//router.PUT("/person/:id", ModPersonApi)
	//
	//router.DELETE("/person/:id", DelPersonApi)

	router.POST("/loginFrom",LoginForm)

	taR := router.Group("/data")
	taR.Use(jwt.JWTAuth())

	{
		taR.GET("/dataByTime", GetDataByTime)
	}

	return router
}