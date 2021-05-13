package main

import (
	"github.com/gin-gonic/gin"
	a "zexho994/go-gather-search/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode("release")
	api := r.Group("/api")
	{
		api.POST("/user/page/favorite", a.AddFavoritePage)
	}
	return r
}
