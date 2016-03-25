package game

import (
	"crypto/rand"
	"io"
	"log"
)

//DefaultStrength of the ai
const DefaultStrength = 5

const StrongStrength = 6

type aiPlayer struct {
	seed     io.Reader
	color    Move
	strength int
}

func getAllPossibleMoves(b Board, c Move) []Position {
	positions := []Position{}
	for _, f := range b.GetAllowedFields() {
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				copyBoard := b.Copy()
				fx := f % 3
				fy := f / 3

				err := copyBoard.PutStone(fx, fy, x, y, c)
				if err == nil {
					positions = append(positions, Position{fx, fy, x, y})
				}
			}
		}
	}

	return positions
}

func (r aiPlayer) NextMove(b *Board) {
	rating := 0
	allMoves := getAllPossibleMoves(b.Copy(), r.color)
	if len(allMoves) == 0 {
		panic("no moves left")
	}

	bestMove := allMoves[0]
	for _, possibleMove := range allMoves {
		thisRating := findBestMove(b.Copy(), possibleMove, r.color, r.color, 0, 0, r.strength)
		if thisRating > rating {
			bestMove = possibleMove
			rating = thisRating
		}
	}

	log.Printf("AI auf [%d|%d](%d|%d) => %s\n", bestMove.fx, bestMove.fy, bestMove.x, bestMove.y, r.color)
	b.PutStone(bestMove.fx, bestMove.fy, bestMove.x, bestMove.y, r.color)
}
func findBestMove(
	b Board,
	move Position,
	color,
	winColor Move,
	rating,
	level,
	maxLevel int,
) int {
	board := b.Copy()
	err := board.PutStone(move.fx, move.fy, move.x, move.y, color)
	if err != nil {
		return rating + 1
	}

	if board.HasWinner() {
		if board.GetWinner() != winColor {
			return rating - 333
		}

		return rating + 1000
	}

	if level >= maxLevel {
		return level
	}

	if board.IsFull() {
		return rating
	}

	nextColor := MoveCircle
	if color == MoveCircle {
		nextColor = MoveCross
	}

	bestRating := 0

	for _, f := range board.GetAllowedFields() {
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				fx := f % 3
				fy := f / 3

				pos := Position{fx: fx, fy: fy, x: x, y: y}
				thisRating := findBestMove(board, pos, nextColor, winColor, rating, level+1, maxLevel)
				if thisRating > bestRating {
					bestRating = thisRating
				}
			}
		}
	}

	for _, f := range b.data {
		if f.HasWinner() {
			if f.GetWinner() != winColor {
				bestRating -= 100
			}
		}
	}

	if color != winColor {
		return bestRating - rating
	}

	return bestRating + rating
}

//NewAIPlayer returns a brute force intelligent enemy
func NewAIPlayer(color Move, strength int) Player {
	return &aiPlayer{seed: rand.Reader, color: color, strength: strength}
}
