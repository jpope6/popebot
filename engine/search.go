package engine

import "fmt"

func searchPosition(depth int) {
	fmt.Printf("bestmove d2d4\n")
}

func Evaluate(bs *BoardState) int {
	score := 0

	for piece := P; piece <= k; piece++ {
		pieceColor := piece / 6
		pieceType := piece % 6
		bb := bs.Position.Pieces[pieceColor][pieceType]

		for bb != 0 {
			square := bb.GetLsbIndex()

			score += mgMaterialScore[piece]
			score += getMgPsqtScore(piece, square)

			bb.PopBit(square)
		}
	}

	multiplier := 1
	if bs.Turn == Black {
		multiplier = -1
	}

	return score * multiplier
}

func getMgPsqtScore(piece, square uint8) int {
	switch piece {
	case P:
		return mgPawnTable[mirrorScore[square]]
	case N:
		return mgKnightTable[mirrorScore[square]]
	case B:
		return mgBishopTable[mirrorScore[square]]
	case R:
		return mgRookTable[mirrorScore[square]]
	case Q:
		return mgQueenTable[mirrorScore[square]]
	case K:
		return mgKingTable[mirrorScore[square]]
	case p:
		return -mgPawnTable[square]
	case n:
		return -mgKnightTable[square]
	case b:
		return -mgBishopTable[square]
	case r:
		return -mgRookTable[square]
	case q:
		return -mgQueenTable[square]
	case k:
		return -mgKingTable[square]
	}

	return 0
}
