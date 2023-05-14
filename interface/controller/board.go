package controller

import usecase "xo_connect5/usecase/board"

type BoardController struct {
	bi *usecase.BoardInteractor
}

func NewBoardController(bi *usecase.BoardInteractor) *BoardController {
	return &BoardController{bi}
}

func (bc *BoardController) GetBoard() string {
	return bc.bi.GetBoard()
}
