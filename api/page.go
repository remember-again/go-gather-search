package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zexho994/go-gather-search/form"
	"zexho994/go-gather-search/service"
)

// AddFavoritePageApi storage page for user
func AddFavoritePageApi(c *gin.Context) {
	fpForm := form.AddFpForm{}
	err := c.BindJSON(&fpForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "parameter error"})
		return
	}

	service.AddFpService(fpForm)

	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	return
}
