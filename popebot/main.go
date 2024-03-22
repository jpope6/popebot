package main

import (
	"fmt"
	"popebot/engine"
)

func main() {
	boardState := engine.InitBoardState("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fmt.Printf("%b\n", boardState.CastleRights)
}
