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
	fp := r.Group("/api/user/page/favorite")
	{
		fp.POST("", a.AddFpApi)
		fp.GET("/email/:email", a.FindFpApi)
	}

	return r
}
