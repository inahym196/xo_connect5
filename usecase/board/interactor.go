package usecase

type BoardInteractor struct {
	br IBoardRepository
}

func (bi *BoardInteractor) GetBoard() string {
	b := bi.br.GetBoard()
	return b.String()
}

func NewBoardInteractor(br IBoardRepository) *BoardInteractor {
	return &BoardInteractor{br}
}
