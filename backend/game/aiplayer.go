package game

import (
	"crypto/rand"
	"io"
	"math/big"
)

//DefaultStrength of the ai
const DefaultStrength = 6

type aiPlayer struct {
	seed     io.Reader
	color    Move
	strength int
}

//SetColor to change the users move color
func (r *aiPlayer) SetColor(c Move) {
	r.color = c
}

func getAllPossibleMoves(b Board, c Move) []Position {
	positions := []Position{}
	for _, f := range b.GetAllowedFields() {
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				fx := f % 3
				fy := f / 3

				if b.CanPutStone(fx, fy, x, y, c) {
					positions = append(positions, Position{fx: fx, fy: fy, x: x, y: y})
				}
			}
		}
	}

	return positions
}

func (r aiPlayer) NextMove(b *Board) Position {
	rating := 0
	allMoves := getAllPossibleMoves(*b, r.color)
	if len(allMoves) == 0 {
		panic("no moves left")
	}

	bestMove := allMoves[0]
	for _, possibleMove := range allMoves {
		n, _ := rand.Int(r.seed, big.NewInt(int64(10)))
		startValue := int(n.Int64())

		thisRating := findBestMove(*b, possibleMove, r.color, r.color, startValue, 0, r.strength)
		if thisRating > rating {
			bestMove = possibleMove
			rating = thisRating
		}
	}

	/*
	 *log.Printf("Moving to [%d|%d](%d|%d) => %s Confidence[%d]\n", bestMove.fx, bestMove.fy, bestMove.x, bestMove.y, r.color, rating)
	 */
	b.PutStone(bestMove.fx, bestMove.fy, bestMove.x, bestMove.y, r.color)
	return bestMove
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
	err := b.PutStone(move.fx, move.fy, move.x, move.y, color)
	if err != nil {
		return 0
	}

	defer b.Undo()
	if b.HasWinner() {
		if b.GetWinner() != winColor {
			return -1
		}

		return 1
	}

	if level >= maxLevel {
		return level
	}

	nextColor := MoveCircle
	if color == MoveCircle {
		nextColor = MoveCross
	}

	bestRating := 0

	for _, f := range b.GetAllowedFields() {
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				fx := f % 3
				fy := f / 3

				pos := Position{fx: fx, fy: fy, x: x, y: y}
				thisRating := findBestMove(b, pos, nextColor, winColor, rating, level+1, maxLevel)
				if thisRating > bestRating {
					bestRating = thisRating
				}
			}
		}
	}

	return rating + bestRating
}

//NewAIPlayer returns a brute force intelligent enemy
func NewAIPlayer(color Move, strength int) Player {
	return &aiPlayer{seed: rand.Reader, color: color, strength: strength}
}
