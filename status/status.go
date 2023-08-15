package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"status": id,
	})
}
