package game_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

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
			winXs := 0
			winOs := 0
			winXn := 0
			winOn := 0

			start := time.Now()
			startColor := false
			for i := 0; i < 50; i++ {
				game = NewTicTacNineGame()
				b := game.Board()
				circle := i%2 == 0
				startColor = circle
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
					if startColor {
						winOs++
					} else {
						winOn++
					}
				case MoveCross:
					if !startColor {
						winXs++
					} else {
						winXn++
					}
				case MoveNone:
					draws++
				}

				fmt.Printf(
					"Draws %d, AI (X/weak): [%d](%d/%d), Random (O): [%d](%d/%d)\nTime Elapsed %s\n",
					draws, winXs+winXn, winXs, winXn, winOs+winOn, winOs, winOn, time.Now().Sub(start),
				)
			}
		})
	})
})
