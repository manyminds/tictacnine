package game_test

import (
	"fmt"
	"io/ioutil"
	"log"

	. "github.com/manyminds/tictacnine/backend/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
	var (
		game         Game
		playerCircle Player
		playerCross  Player
	)

	BeforeEach(func() {
		game = NewTicTacNineGame()
		playerCircle = NewRandomPlayer(MoveCircle)
		playerCross = NewAIPlayer(MoveCross, DefaultStrength)
		log.SetOutput(ioutil.Discard)
	})

	Context("can play against random enemy", func() {
		It("should answer with a correct move", func() {
			draws := 0
			winX := 0
			winO := 0
			for i := 0; i < 100; i++ {
				game = NewTicTacNineGame()
				b := game.Board()
				circle := i%2 == 0
				for !b.HasWinner() && !b.IsFull() {
					if circle {
						playerCircle.NextMove(b)
					} else {
						playerCross.NextMove(b)
					}
					circle = !circle
				}

				Expect(b.HasWinner() || b.IsFull()).To(Equal(true))
				switch b.GetWinner() {
				case MoveCircle:
					winO++
				case MoveCross:
					winX++
				case MoveNone:
					draws++
				}

			}
			fmt.Printf("Draws %d, X: %d, O: %d\n", draws, winX, winO)
		})
	})
})
