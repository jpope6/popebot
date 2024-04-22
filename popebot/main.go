package main

import (
	"fmt"
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "8/8/8/8/8/8/3p4/8 w - - 0 1"
	command  = "position startpos moves e2e4 e7e5 g1f3 b8c6"
	debug    = true
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState

	if debug {
		bs.InitBoardState(testFen)
		engine.PrintBoard(&bs)
		fmt.Printf("\nScore: %d\n", engine.Evaluate(&bs))
	} else {
		engine.UciLoop(&bs)
	}

}
