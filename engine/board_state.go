package engine

import (
	"strings"
)

// To represent the board, we will store multiple bitboards
// in a Position object
//
// Turn will be a true or false to represent whose turn it is
// true = white to move
// false = black to move
//
// castleRights will be 4-bit binary number such as 0110
// position in this order -> wk wq bk bq
// 0 = no castle rights
// 1 = castle rights
type BoardState struct {
	Position     Position
	Turn         bool
	CastleRights uint8
	EpSquare     uint8
}

// Initialize board state with a FEN string
func InitBoardState(FEN string) *BoardState {
	bs := BoardState{}

	// Split up the FEN string
	FENfields := strings.Fields(FEN)
	squares := FENfields[0]
	turn := FENfields[1]
	castleRights := FENfields[2]
	// epSquare := FENfields[3]
	// halfmove := fields[4]
	// fullmove := fields[5]

	// Initialize Position
	bs.Position.SetPositionWithFEN(squares)

	// Initialize Turn
	bs.Turn = turn == "w"

	// Initialize CastleRights
	var rights uint8 = 0x0 // 0000
	for _, right := range castleRights {
		switch right {
		case 'K':
			rights |= WhiteKingSide
		case 'Q':
			rights |= WhiteQueenSide
		case 'k':
			rights |= BlackKingSide
		case 'q':
			rights |= BlackQueenSide
		}
	}
	bs.CastleRights = rights

	// Initialize epSquare
	// bs.EpSquare = epSquare

	return &bs
}
