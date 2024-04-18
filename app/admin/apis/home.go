package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "index")
	})
}
