package domain

import "encoding/json"

type Board struct {
	Value [10][10]int
}

func (b *Board) String() string {
	jsonData, err := json.Marshal(b.Value)
	if err != nil {
		panic(0)
	}
	return string(jsonData)
}
