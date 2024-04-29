package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "1k6/8/2n1n2P/8/8/8/8/2K5 w - - 0 2"
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState
	debug := false

	if debug {
		bs.InitBoardState(testFen)
		engine.PrintBoard(&bs)
		bs.Search(6)
	} else {
		engine.UciLoop(&bs)
	}

}
