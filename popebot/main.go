package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "2k5/8/8/8/2pP4/8/8/2K5 b - d3 0 1"
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState
	bs.InitBoardState(testFen)
	engine.PrintBoard(&bs)

	engine.GeneratePawnMoves(&bs)
}
