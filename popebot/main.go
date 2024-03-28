package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "8/3k4/1n4q1/8/3B4/8/8/4K3 w - - 0 1"
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState
	bs.InitBoardState(startFen)
	engine.PrintBoard(&bs)
}
