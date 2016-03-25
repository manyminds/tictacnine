package game

import "errors"

type area struct {
	field  [9]Move
	winner Move
}

func (a area) GetStone(x, y int) Move {
	index := x + (y * 3)
	return a.field[index]
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

func (a area) IsFull() bool {
	for _, m := range a.field {
		if m == MoveNone {
			return false
		}
	}

	return true
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
