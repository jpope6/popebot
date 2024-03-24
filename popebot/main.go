package main

import (
	"popebot/engine"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	testFen  = "8/3k4/1n4q1/8/3B4/8/8/4K3 w - - 0 1"
)

func main() {
	// for i := uint8(0); i < 64; i++ {
	// 	engine.PrintBitboard(engine.MaskBishopAttacks(i))
	// }

	var blockers engine.Bitboard
	blockers.SetBit(engine.D7)
	blockers.SetBit(engine.B4)
	blockers.SetBit(engine.D2)
	blockers.SetBit(engine.G4)

	engine.PrintBitboard(blockers)
	engine.PrintBitboard(engine.MaskRookAttacksWithBlockers(engine.D4, blockers))
}
