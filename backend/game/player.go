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
}

//Player interface for playas
type Player interface {
	NextMove(b *Board)
}

type randomPlayer struct {
	seed  io.Reader
	color Move
}

func (r randomPlayer) NextMove(b *Board) {
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
				return
			}
			array = append(array[:i], array[i+1:]...)
		}

	}

	panic("no move left")
}

//NewRandomPlayer returns a stupid simple random enemy
func NewRandomPlayer(color Move) Player {
	return &randomPlayer{seed: rand.Reader, color: color}
}
