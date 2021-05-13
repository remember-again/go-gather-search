package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zexho994/go-gather-search/form"
)

// AddFavoritePage storage page for user
func AddFavoritePage(c *gin.Context) {
	fpForm := form.AddFpForm{}
	err := c.BindJSON(&fpForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "parameter error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	return
}
