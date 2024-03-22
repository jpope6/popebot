package engine

import (
	"strings"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	// Constants for mapping board squares to numbers.
	A1, B1, C1, D1, E1, F1, G1, H1 = 0, 1, 2, 3, 4, 5, 6, 7
	A2, B2, C2, D2, E2, F2, G2, H2 = 8, 9, 10, 11, 12, 13, 14, 15
	A3, B3, C3, D3, E3, F3, G3, H3 = 16, 17, 18, 19, 20, 21, 22, 23
	A4, B4, C4, D4, E4, F4, G4, H4 = 24, 25, 26, 27, 28, 29, 30, 31
	A5, B5, C5, D5, E5, F5, G5, H5 = 32, 33, 34, 35, 36, 37, 38, 39
	A6, B6, C6, D6, E6, F6, G6, H6 = 40, 41, 42, 43, 44, 45, 46, 47
	A7, B7, C7, D7, E7, F7, G7, H7 = 48, 49, 50, 51, 52, 53, 54, 55
	A8, B8, C8, D8, E8, F8, G8, H8 = 56, 57, 58, 59, 60, 61, 62, 63

	// Hexadecimal number that corresponds to casteling rights
	WhiteKingSide  = 0xF // 1000
	WhiteQueenSide = 0x4 // 0100
	BlackKingSide  = 0x2 // 0010
	BlackQueenSide = 0x1 // 0001
)

// To represent the board we typically need one bitboard for each
// piece-type and color - encapsulated inside a position structure.
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
