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

		It("validate start field", func() {
			board := game.Board()
			err := board.PutStone(1, 1, 1, 1, MoveCircle)
			Expect(err).ToNot(HaveOccurred())

			err = board.PutStone(0, 1, 1, 1, MoveCircle)
			Expect(err).To(HaveOccurred())

			err = board.PutStone(0, 0, 1, 1, MoveCircle)
			Expect(err).To(HaveOccurred())

			err = board.PutStone(2, 1, 1, 1, MoveCircle)
			Expect(err).To(HaveOccurred())

			err = board.PutStone(1, 2, 1, 1, MoveCircle)
			Expect(err).To(HaveOccurred())

			err = board.PutStone(2, 2, 1, 1, MoveCircle)
			Expect(err).To(HaveOccurred())

			err = board.PutStone(0, 2, 1, 1, MoveCircle)
			Expect(err).To(HaveOccurred())

			err = board.PutStone(2, 0, 1, 1, MoveCircle)
			Expect(err).To(HaveOccurred())
		})
	})
})
