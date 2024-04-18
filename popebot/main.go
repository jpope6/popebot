package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "8/8/2P1r3/8/4n3/1b3B2/8/8 b - - 0 1"
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState
	bs.InitBoardState(testFen)
	engine.PrintBoard(&bs)

	engine.GenerateAllMoves(&bs)
}
