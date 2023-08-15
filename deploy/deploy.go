package deploy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Deploy(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"deploy": id,
	})
}
