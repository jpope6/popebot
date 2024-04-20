package main

import (
	"fmt"
	"popebot/engine"
	"time"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
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

	engine.PerftDriver(&bs, 4, &nodes)

	duration := time.Since(start)
	fmt.Printf("Time: %d ms\n", duration.Milliseconds())
	fmt.Printf("Nodes: %d\n", nodes)
}
