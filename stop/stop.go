package stop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Stop(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"stop": id,
	})
}
