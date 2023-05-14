package usecase

import "xo_connect5/domain"

type IBoardRepository interface {
	GetBoard() domain.Board
}
