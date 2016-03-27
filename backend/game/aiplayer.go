package game

import (
	"crypto/rand"
	"io"
	"log"
	"math/big"
)

//DefaultStrength of the ai
const DefaultStrength = 6

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
				fx := f % 3
				fy := f / 3

				if b.CanPutStone(fx, fy, x, y, c) {
					positions = append(positions, Position{fx, fy, x, y})
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
		log.Printf("%s\n", b.String())
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

	log.Printf("AI auf [%d|%d](%d|%d) => %s\n", bestMove.fx, bestMove.fy, bestMove.x, bestMove.y, r.color)
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
	board := b.Copy()
	enemyFieldsWon := 0
	ownFieldsWon := 0
	for _, f := range board.data {
		if f.HasWinner() {
			if f.GetWinner() == winColor {
				ownFieldsWon++
			} else {
				enemyFieldsWon++
			}
		}
	}

	err := board.PutStone(move.fx, move.fy, move.x, move.y, color)
	if err != nil {
		return rating + 1
	}

	if board.HasWinner() {
		if board.GetWinner() != winColor {
			return rating - 100000
		}

		return rating + 1000000
	}

	if level >= maxLevel {
		return level
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

	enemyFieldsWonAfter := 0
	for _, f := range board.data {
		if f.HasWinner() {
			if f.GetWinner() != winColor && f.GetWinner() != MoveNone {
				bestRating -= 400
			} else if f.GetWinner() == winColor {
				bestRating += 50
			}
		}
	}

	if enemyFieldsWon < enemyFieldsWonAfter {
		rating -= 100
	}

	return rating + bestRating
}

//SetColor to change the users move color
func (a *aiPlayer) SetColor(c Move) {
	a.color = c
}

//NewAIPlayer returns a brute force intelligent enemy
func NewAIPlayer(color Move, strength int) Player {
	return &aiPlayer{seed: rand.Reader, color: color, strength: strength}
}
