package game_test

import (
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
			for i := 0; i < 3; i++ {
				Expect(board.PutStone(i, 0, 1, 1, MoveCircle)).ToNot(HaveOccurred())
				Expect(game.Board().HasWinner()).To(BeFalse())

				Expect(board.PutStone(i, 0, 1, 0, MoveCircle)).ToNot(HaveOccurred())
				Expect(game.Board().HasWinner()).To(BeFalse())

				Expect(board.PutStone(i, 0, 1, 2, MoveCircle)).ToNot(HaveOccurred())
				//after the last field the board should be won
				Expect(game.Board().HasWinner()).To(Equal(i == 2))
			}
		})
	})
})
