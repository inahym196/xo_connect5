package usecase

type GameInteractor struct {
	gr IGameRepository
}

func (gi *GameInteractor) GetGame() string {
	g := gi.gr.GetGame()
	return g.String()
}

func NewGameInteractor(gr IGameRepository) *GameInteractor {
	return &GameInteractor{gr}
}
