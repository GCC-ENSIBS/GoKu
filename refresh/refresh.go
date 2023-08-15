package refresh

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Refresh(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"refresh": id,
	})
}
