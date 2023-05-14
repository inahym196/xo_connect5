package usecase

import "xo_connect5/domain"

type IGameRepository interface {
	GetGame() domain.Game
}
