package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "r1bqkbnr/pppp1p1p/2n3p1/4p3/2B1P3/5Q2/PPPP1PPP/RNB1K1NR w KQkq - 0 1"
	command  = "position startpos moves e2e4 e7e5 g1f3 b8c6"
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState
	debug := true

	if debug {
		bs.InitBoardState(testFen)
		engine.PrintBoard(&bs)
		bs.Search(6)
	} else {
		engine.UciLoop(&bs)
	}

}
