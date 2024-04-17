package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "r3k2r/8/8/8/8/8/4r3/2RNK1R1 b KQkq - 0 1"
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState
	bs.InitBoardState(testFen)
	engine.PrintBoard(&bs)

	engine.GenerateMoves(&bs)
}
