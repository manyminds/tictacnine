package game

import (
	"errors"
	"fmt"
)

//Move to lay a stone
type Move int

const (
	//MoveNone none
	MoveNone Move = iota
	//MoveCross turn for X
	MoveCross

	//MoveCircle turn for O
	MoveCircle
)

func (m Move) String() string {
	if m == MoveNone {
		return "_"
	}
	if m == MoveCross {
		return "X"
	}

	return "O"
}

var winLines = [...][3]int{
	[3]int{0, 4, 8},
	[3]int{0, 1, 2},
	[3]int{0, 3, 6},
	[3]int{1, 4, 7},
	[3]int{2, 5, 8},
	[3]int{2, 4, 6},
	[3]int{3, 4, 5},
	[3]int{6, 7, 8},
}

//Game interface to control the whole board
type Game interface {
	Board() *Board
}

//Placeable you can put a stone there
type Placeable interface {
	PutStone(x, y int, m Move) error
	GetStone(x, y int) Move
}

//PlaceAndWinable placing possible an winning ;)
type PlaceAndWinable interface {
	Placeable
	Winable
}

//Winable if a board is winnable
type Winable interface {
	HasWinner() bool
	GetWinner() Move
}

//Board the game board
type Board struct {
	data [9]PlaceAndWinable
}

func (b Board) String() string {
	result := "%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n"
	variables := []interface{}{}

	for _, field := range b.data {
		fieldString := ""
		for y := 0; y < 3; y++ {
			fieldString += "|"
			for x := 0; x < 3; x++ {
				fieldString += field.GetStone(x, y).String()
				fieldString += "|"
			}

			fieldString += "\n"
		}

		if field.HasWinner() {
			fieldString = fieldString[:len(fieldString)-1]
			fieldString += fmt.Sprintf(" => [%s] \n", field.GetWinner())
		}

		variables = append(variables, fieldString)
	}

	return fmt.Sprintf(result, variables...)
}

//NewBoard initialize a new game board
func NewBoard() *Board {
	return &Board{data: [9]PlaceAndWinable{
		&area{},
		&area{},
		&area{},
		&area{},
		&area{},
		&area{},
		&area{},
		&area{},
		&area{},
	}}
}

type area struct {
	field  [9]Move
	winner Move
}

func (a area) GetStone(x, y int) Move {
	index := x + (y * 3)
	return a.field[index]
}

//HasWinner bool
func (b Board) HasWinner() bool {
	for _, line := range winLines {
		if b.data[line[0]].GetWinner() == MoveNone {
			continue // one field not placed yet
		}

		if b.data[line[0]].GetWinner() == b.data[line[1]].GetWinner() &&
			b.data[line[1]].GetWinner() == b.data[line[2]].GetWinner() {
			return true
		}
	}

	return false
}

//GetWinner for the whole board
func (b Board) GetWinner() Move {
	for _, line := range winLines {
		if b.data[line[0]].GetWinner() == MoveNone {
			continue // one field not placed yet
		}

		if b.data[line[0]].GetWinner() == b.data[line[1]].GetWinner() &&
			b.data[line[1]].GetWinner() == b.data[line[2]].GetWinner() {
			return b.data[line[0]].GetWinner()
		}
	}

	return MoveNone
}

//PutStone for the whole board
func (b *Board) PutStone(fx, fy, x, y int, m Move) error {
	if fx < 0 || fx >= 3 {
		return errors.New("fx out of bounds")
	}

	if fy < 0 || fy >= 3 {
		return errors.New("fy out of bounds")
	}

	index := fx + (fy * 3)
	return b.data[index].PutStone(x, y, m)
}

//PutStone adds a stone if possible and saves the winner
//if the field now has a winner and did not got one previously
func (a *area) PutStone(x, y int, m Move) error {
	if x < 0 || x >= 3 {
		return errors.New("x out of bounds")
	}

	if y < 0 || y >= 3 {
		return errors.New("y out of bounds")
	}

	index := x + (y * 3)
	if a.field[index] != MoveNone {
		return errors.New("place already has a stone")
	}

	a.field[index] = m
	if a.winner == MoveNone && a.HasWinner() {
		a.winner = m
	}

	return nil
}

func (a area) GetWinner() Move {
	return a.winner
}

//HasWinner for this area
func (a area) HasWinner() bool {
	for _, line := range winLines {
		if a.field[line[0]] == MoveNone {
			continue // one field not placed yet
		}

		if a.field[line[0]] == a.field[line[1]] &&
			a.field[line[1]] == a.field[line[2]] {
			return true
		}
	}

	return false
}

//Player interface for playas
type Player interface {
	GetNextMove() (x, y int, hasMove bool)
}

type humanPlayer struct {
}

func (h humanPlayer) GetNextMove() (int, int, bool) {
	return -1, -1, false
}

type randomPlayer struct {
}

type ticTacNineGame struct {
	b *Board
}

func (t ticTacNineGame) Board() *Board {
	return t.b
}

//NewTicTacNineGame returns a new instance of tic tac nine
func NewTicTacNineGame() Game {
	return &ticTacNineGame{b: NewBoard()}
}
