package main

import (
	"fmt"
	"popebot/engine"
	"time"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
)

func init() {
	engine.InitTables()
}

func main() {
	var nodes engine.Nodes = 0
	var bs engine.BoardState
	bs.InitBoardState(testFen)
	engine.PrintBoard(&bs)

	start := time.Now()

	engine.PerftDriver(&bs, 3, &nodes)

	duration := time.Since(start)
	fmt.Printf("Time: %d ms\n", duration.Milliseconds())
	fmt.Printf("Nodes: %d\n", nodes)
}
