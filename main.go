package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	board := make([][]int, 10)
	for i := range board {
		board[i] = make([]int, 10)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, fmt.Sprint(board))
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
