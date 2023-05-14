package http

import (
	"xo_connect5/interface/controller"
	repository "xo_connect5/interface/repository/inmemory"
	usecase "xo_connect5/usecase/board"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	bc := controller.NewBoardController(usecase.NewBoardInteractor(repository.NewBoardRepository()))
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, bc.GetBoard())
	})
	return r
}
