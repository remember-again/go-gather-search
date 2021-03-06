package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zexho994/go-gather-search/form"
	service "zexho994/go-gather-search/service"
)

// AddFpApi storage page for user
func AddFpApi(c *gin.Context) {
	fpForm := form.AddFpForm{}
	err := c.BindJSON(&fpForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "parameter error"})
		return
	}

	service.AddFp(fpForm)

	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	return
}

// FindFpApi find user's fp
func FindFpApi(c *gin.Context) {
	res := service.FindFp(c.Param("email"))
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func RemoveFpApi(c *gin.Context) {
	fpForm := form.RemoveFpForm{}
	err := c.BindJSON(&fpForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "parameter error"})
		return
	}
	effect := service.RemoveFp(fpForm.Email, fpForm.Url)
	c.JSON(http.StatusOK, gin.H{"msg": effect})
	return
}

func ExistApi(c *gin.Context) {
	flag := service.Exist(c.Param("email"), c.Param("url"))
	c.JSON(http.StatusOK, gin.H{"data": flag})
}
