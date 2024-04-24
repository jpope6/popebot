package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 9"
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
