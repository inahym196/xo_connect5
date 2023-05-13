package infra

import (
	"xo_connect5/domain"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	b := domain.Board{}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, b.String())
	})
	return r
}
