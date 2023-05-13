package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	board := make([][]int, 10)
	for i := range board {
		board[i] = make([]int, 10)
	}
	jsonData, err := json.Marshal(board)
	if err != nil {
		panic(0)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Printf("%s", jsonData)
		c.String(200, string(jsonData))
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run("127.0.0.1:8080")
}
