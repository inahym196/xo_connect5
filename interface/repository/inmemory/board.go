package repository

import "xo_connect5/domain"

type BoardRepository struct{}

func (bi *BoardRepository) GetBoard() domain.Board {
	b := domain.Board{}
	return b
}

func NewBoardRepository() *BoardRepository {
	return &BoardRepository{}
}
