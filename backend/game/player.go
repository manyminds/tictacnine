package game

import (
	"crypto/rand"
	"io"
	"log"
	"math/big"
)

//Position wrapper for one move
type Position struct {
	fx, fy, x, y int
	color        Move
}

//Field returns the field position
func (p Position) Field() (int, int) {
	return p.fx, p.fy
}

//Position returns the position in a field
func (p Position) Position() (int, int) {
	return p.x, p.y
}

//Player interface for playas
type Player interface {
	NextMove(b *Board) Position
	SetColor(color Move)
}

type randomPlayer struct {
	seed  io.Reader
	color Move
}

func (r randomPlayer) NextMove(b *Board) Position {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	for _, field := range b.GetAllowedFields() {
		fx := field % 3
		fy := field / 3

		for len(array) > 0 {
			n, _ := rand.Int(r.seed, big.NewInt(int64(len(array))))
			i := n.Int64()

			move := array[i]

			x := move % 3
			y := move / 3

			err := b.PutStone(fx, fy, x, y, r.color)
			if err == nil {
				log.Printf("Random auf [%d|%d](%d|%d) => %s\n", fx, fy, x, y, r.color)
				return Position{fx: fx, fy: fy, x: x, y: y}
			}
			array = append(array[:i], array[i+1:]...)
		}

	}

	panic("no move left")
}

func (r *randomPlayer) SetColor(c Move) {
	r.color = c
}

//NewRandomPlayer returns a stupid simple random enemy
func NewRandomPlayer(color Move) Player {
	return &randomPlayer{seed: rand.Reader, color: color}
}
