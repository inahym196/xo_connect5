package domain

import "encoding/json"

type Board struct {
	Pieces [10][10]int
}

type Game struct {
	Board *Board
	Turn  int
}

func NewGame() *Game {
	return &Game{Board: &Board{}}
}

func (g *Game) String() string {
	jsonData, err := json.Marshal(&g)
	if err != nil {
		panic(0)
	}
	return string(jsonData)
}
