package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "4k3/1Q6/5K2/8/8/8/8/8 w - - 0 1"
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
