package game_test

import (
	"fmt"

	. "github.com/manyminds/tictacnine/backend/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	Context("Test Area Code", func() {
		var (
			game Game
		)

		BeforeEach(func() {
			game = NewTicTacNineGame()
		})

		It("won't have winner on an empty board", func() {
			Expect(game.Board().HasWinner()).To(BeFalse())
		})

		It("should calculate the winner correctly", func() {
			board := game.Board()
			Expect(board.PutStone(0, 0, 1, 1, MoveCircle)).ToNot(HaveOccurred())
			fmt.Printf("%s\n", board)
			Expect(game.Board().HasWinner()).To(BeFalse())

			Expect(board.PutStone(0, 0, 1, 0, MoveCircle)).ToNot(HaveOccurred())
			fmt.Printf("%s\n", board)
			Expect(game.Board().HasWinner()).To(BeFalse())

			Expect(board.PutStone(0, 0, 1, 2, MoveCircle)).ToNot(HaveOccurred())
			fmt.Printf("%s\n", board)
			Expect(game.Board().HasWinner()).To(BeFalse())
		})
	})
})
