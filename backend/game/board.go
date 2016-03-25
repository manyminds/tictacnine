package game

import (
	"errors"
	"fmt"
)

//Board the game board
type Board struct {
	data [9]PlaceAndWinable
	lastMoveX,
	lastMoveY int
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

	index := fx + (fy * 3)
	allowedFields := b.GetAllowedFields()
	b.lastMoveX = x
	b.lastMoveY = y

	for _, f := range allowedFields {
		if f == index {
			return b.data[index].PutStone(x, y, m)
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
