package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(gin.Logger())

	r.POST("/shorten", ShortenURL())
	r.GET("/:key", GetDetails())
	r.GET("/details", GetMultipleRecords())

	return r
}
