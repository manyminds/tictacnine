package game

import "log"

//Game interface to control the whole board
type Game interface {
	Board() *Board
}

//Placeable you can put a stone there
type Placeable interface {
	PutStone(x, y int, m Move) error
	GetStone(x, y int) Move
	CanPutStone(x, y int, m Move) bool
}

//PlaceAndWinable placing possible an winning ;)
type PlaceAndWinable interface {
	Placeable
	Winable
	IsFull() bool
}

//Winable interface to
type Winable interface {
	HasWinner() bool
	GetWinner() Move
}

type ticTacNineGame struct {
	b *Board
}

func (t ticTacNineGame) Board() *Board {
	return t.b
}

//NewTicTacNineGame returns a new instance of tic tac nine
func NewTicTacNineGame() Game {
	b, err := NewBoard(1, 1)
	if err != nil {
		log.Fatal(err)
	}

	return &ticTacNineGame{b: b}
}
