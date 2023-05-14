package controller

import usecase "xo_connect5/usecase/game"

type GameController struct {
	gi *usecase.GameInteractor
}

func NewGameController(bi *usecase.GameInteractor) *GameController {
	return &GameController{bi}
}

func (gc *GameController) GetGame() string {
	return gc.gi.GetGame()
}
