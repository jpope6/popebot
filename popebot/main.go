package main

import (
	"fmt"
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "4k3/8/8/8/8/8/8/R3K2R w KQ - 0 1"
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState
	bs.InitBoardState(testFen)
	moves := engine.GenerateAllMoves(&bs)
	engine.PrintBoard(&bs)

	move := moves.ParseMoveString("e1g1")

	if move != 0 {
		bs.MakeMove(move, engine.AllMoves)
	} else {
		fmt.Println("\nIllegal move")
	}

	engine.PrintBoard(&bs)
}
