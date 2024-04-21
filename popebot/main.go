package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "4k3/8/8/8/8/8/8/R3K2R w KQ - 0 1"
	command  = "position startpos moves e2e4 e7e5 g1f3 b8c6"
)

func init() {
	engine.InitTables()
}

func main() {
	var bs engine.BoardState

	engine.UciLoop(&bs)
}
