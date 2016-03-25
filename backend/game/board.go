package game

import (
	"errors"
	"fmt"
)

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

//Board the game board
type Board struct {
	data [9]PlaceAndWinable
	lastMoveX,
	lastMoveY int
	lastMoveColor Move
}

func (b Board) Copy() Board {
	b2, _ := NewBoard(b.lastMoveX, b.lastMoveY)
	for k, v := range b.data {
		if a, ok := v.(*area); ok {
			b2.data[k] = a.Copy()
		} else {
			b2.data[k] = v
		}
	}

	b2.lastMoveColor = b.lastMoveColor
	return *b2
}

func (b Board) IsFull() bool {
	for _, f := range b.data {
		if !f.IsFull() {
			return false
		}
	}
	return true
}

//GetAllowedFields returns all allowed fields to play
func (b Board) GetAllowedFields() []int {
	nextField := b.lastMoveX + (b.lastMoveY * 3)
	if !b.data[nextField].IsFull() {
		return []int{nextField}
	}

	fields := []int{}
	for i, f := range b.data {
		if !f.IsFull() {
			fields = append(fields, i)
		}
	}

	return fields
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

	if b.lastMoveColor == m {
		return errors.New("invalid color to play")
	}

	index := fx + (fy * 3)
	allowedFields := b.GetAllowedFields()
	for _, f := range allowedFields {
		if f == index {
			err := b.data[index].PutStone(x, y, m)
			if err == nil {
				b.lastMoveX = x
				b.lastMoveY = y
				b.lastMoveColor = m
			}

			return err
		}
	}

	return errors.New("invalid field to play in")
}

//NewBoard initialize a new game board
func NewBoard(startX, startY int) (*Board, error) {
	if startX < 0 || startX >= 3 {
		return nil, errors.New("startX out of bounds")
	}

	if startY < 0 || startY >= 3 {
		return nil, errors.New("statY out of bounds")
	}

	return &Board{
		data: [9]PlaceAndWinable{
			&area{},
			&area{},
			&area{},
			&area{},
			&area{},
			&area{},
			&area{},
			&area{},
			&area{},
		},
		lastMoveX: startX,
		lastMoveY: startY,
	}, nil
}
