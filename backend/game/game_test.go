package game_test

import (
	. "github.com/manyminds/tictacnine/backend/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type move struct {
	fx, fy, x, y int
	m            Move
}

var baseOrder = [...]move{
	move{1, 1, 1, 1, MoveCross},
	move{1, 1, 0, 2, MoveCircle},
	move{0, 2, 1, 1, MoveCross},
	move{1, 1, 1, 2, MoveCircle},
	move{1, 2, 1, 1, MoveCross},
	move{1, 1, 2, 2, MoveCircle},
	move{2, 2, 1, 1, MoveCross},
	move{1, 1, 0, 0, MoveCircle},
	move{0, 0, 1, 1, MoveCross},
	move{1, 1, 0, 1, MoveCircle},
	move{0, 1, 1, 1, MoveCross},
	move{1, 1, 1, 0, MoveCircle},
	move{1, 0, 1, 1, MoveCross},
	move{1, 1, 2, 0, MoveCircle},
	move{2, 0, 1, 1, MoveCross},
	move{1, 1, 2, 1, MoveCircle},
	move{2, 1, 1, 1, MoveCross},
}

var outerFields = [...]move{
	move{0, 0, 0, 2, MoveCircle},
	move{0, 1, 0, 2, MoveCircle},
	move{0, 2, 0, 2, MoveCircle},
	move{1, 0, 0, 2, MoveCircle},
	move{1, 2, 0, 2, MoveCircle},
	move{2, 0, 0, 2, MoveCircle},
	move{2, 1, 0, 2, MoveCircle},
	move{2, 2, 0, 2, MoveCircle},
}

var winCircle = [...]move{
	move{2, 1, 0, 1, MoveCircle},
	move{0, 1, 2, 1, MoveCross},
	move{2, 1, 0, 0, MoveCircle},
	move{0, 0, 2, 1, MoveCross},
	move{2, 1, 0, 2, MoveCircle},
	move{0, 2, 0, 1, MoveCross},
	move{0, 1, 0, 0, MoveCircle},
	move{0, 0, 0, 1, MoveCross},
	move{0, 1, 0, 1, MoveCircle},
	move{0, 1, 2, 2, MoveCross},
	move{2, 2, 1, 2, MoveCircle},
	move{1, 2, 0, 1, MoveCross},
	move{0, 1, 0, 2, MoveCircle},
}

var botSequence = [...]move{
	move{1, 1, 1, 2, MoveCircle},
	move{1, 2, 1, 0, MoveCross},
	move{1, 0, 1, 2, MoveCircle},
	move{1, 2, 0, 1, MoveCross},
	move{0, 1, 0, 0, MoveCircle},
	move{0, 0, 2, 2, MoveCross},
	move{2, 2, 0, 0, MoveCircle},
	move{0, 0, 1, 1, MoveCross},
	move{1, 1, 1, 1, MoveCircle},
}

var _ = Describe("Game", func() {
	Context("Test Game Code", func() {
		var (
			game Game
		)

		BeforeEach(func() {
			game = NewTicTacNineGame()
		})

		It("won't have winner on an empty board", func() {
			Expect(game.Board().HasWinner()).To(BeFalse())
		})

		It("can undo moves", func() {
			board := game.Board()

			for _, m := range baseOrder {
				err := board.PutStone(m.fx, m.fy, m.x, m.y, m.m)
				Expect(err).ToNot(HaveOccurred())
			}

			for range baseOrder {
				board.Undo()
			}

			for _, m := range winCircle {
				err := board.PutStone(m.fx, m.fy, m.x, m.y, m.m)
				Expect(err).ToNot(HaveOccurred())
			}

			Expect(board.GetWinner()).To(Equal(MoveCircle))
		})

		It("will be won", func() {
			board := game.Board()

			for _, m := range baseOrder {
				err := board.PutStone(m.fx, m.fy, m.x, m.y, m.m)
				Expect(err).ToNot(HaveOccurred())
			}

			for _, m := range winCircle {
				err := board.PutStone(m.fx, m.fy, m.x, m.y, m.m)
				Expect(err).ToNot(HaveOccurred())
			}

			Expect(board.GetWinner()).To(Equal(MoveCircle))
		})

		It("replay bot sequence to reproduce error", func() {
			board := game.Board()
			for _, m := range botSequence {
				err := board.PutStone(m.fx, m.fy, m.x, m.y, m.m)
				Expect(err).ToNot(HaveOccurred())
			}

			err := board.PutStone(1, 1, 2, 1, MoveCross)

			Expect(err).ToNot(HaveOccurred())
		})

		It("cant move into full field", func() {
			board := game.Board()
			for _, m := range baseOrder {
				err := board.PutStone(m.fx, m.fy, m.x, m.y, m.m)
				Expect(err).ToNot(HaveOccurred())
			}

			err := board.PutStone(1, 1, 1, 1, MoveCircle)

			Expect(err).To(HaveOccurred())
		})

		It("can move anywhere after a full field", func() {
			for _, lastMove := range outerFields {
				game = NewTicTacNineGame()
				board := game.Board()
				for _, m := range baseOrder {
					err := board.PutStone(m.fx, m.fy, m.x, m.y, m.m)
					Expect(err).ToNot(HaveOccurred())
				}

				err := board.PutStone(
					lastMove.fx,
					lastMove.fy,
					lastMove.x,
					lastMove.y,
					lastMove.m,
				)

				Expect(err).ToNot(HaveOccurred())
			}
		})

		It("will fail if same color twice", func() {
			board := game.Board()
			err := board.PutStone(1, 1, 2, 1, MoveCircle)
			Expect(err).ToNot(HaveOccurred())
			err = board.PutStone(2, 1, 2, 0, MoveCircle)
			Expect(err).To(HaveOccurred())
		})

		It("will fail if play is invalid", func() {
			board := game.Board()
			err := board.PutStone(1, 1, 2, 1, MoveCircle)
			Expect(err).ToNot(HaveOccurred())
			err = board.PutStone(0, 1, 2, 0, MoveCross)
			Expect(err).To(HaveOccurred())
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
