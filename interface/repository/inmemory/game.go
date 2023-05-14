package repository

import "xo_connect5/domain"

type GameRepository struct{}

func (bi *GameRepository) GetGame() domain.Game {
	g := domain.NewGame()
	return *g
}

func NewGameRepository() *GameRepository {
	return &GameRepository{}
}
