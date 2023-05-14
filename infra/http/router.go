package http

import (
	"xo_connect5/interface/controller"
	repository "xo_connect5/interface/repository/inmemory"
	usecase "xo_connect5/usecase/game"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gc := controller.NewGameController(usecase.NewGameInteractor(repository.NewGameRepository()))
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, gc.GetGame())
	})
	return r
}
