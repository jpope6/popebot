package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq c6 0 1"
	testFen  = "3k4/8/8/1Pp5/5pP1/8/8/3K4 b - g3 0 1"
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState
	bs.InitBoardState(testFen)
	engine.PrintBoard(&bs)

	moves := engine.GenerateAllMoves(&bs)
	moves.Test(&bs)
}
