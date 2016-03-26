package main

import (
	"fmt"
	"time"

	"github.com/manyminds/tictacnine/backend/game"
)

type ai struct {
	Elo float64
}

func main() {
	g := game.NewTicTacNineGame()
	playerCircle := game.NewRandomPlayer(game.MoveCircle)
	/*
	 *playerCircle := game.NewAIPlayer(game.MoveCircle, game.DefaultStrength)
	 */
	playerCross := game.NewAIPlayer(game.MoveCross, game.DefaultStrength+3)
	/*
	 *log.SetOutput(ioutil.Discard)
	 */
	draws := 0
	winXs := 0
	winOs := 0
	winXn := 0
	winOn := 0

	start := time.Now()
	startColor := false
	for i := 0; i < 1000; i++ {
		g = game.NewTicTacNineGame()
		b := g.Board()
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

		switch b.GetWinner() {
		case game.MoveCircle:
			if startColor {
				winOs++
			} else {
				winOn++
			}
		case game.MoveCross:
			if !startColor {
				winXs++
			} else {
				winXn++
			}
		case game.MoveNone:
			draws++
		}

		fmt.Printf(
			"Draws %d, AI (X/stronger): [%d](%d/%d), AI (O/weaker): [%d](%d/%d)\nTime Elapsed %s\n",
			draws, winXs+winXn, winXs, winXn, winOs+winOn, winOs, winOn, time.Now().Sub(start),
		)
	}
}
