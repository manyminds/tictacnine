package game_test

import (
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
		playerCross = NewRandomPlayer(MoveCross)
	})

	Context("can play against random enemy", func() {
		It("should answer with a correct move", func() {
			b := game.Board()
			circle := true
			for !b.HasWinner() && !b.IsFull() {
				if circle {
					playerCircle.NextMove(b)
				} else {
					playerCross.NextMove(b)
				}
				circle = !circle
			}

			Expect(b.HasWinner() || b.IsFull()).To(Equal(true))
			log.Printf("Winner %s!\n", b.GetWinner().String())
		})
	})
})
