package engine

import (
	"strconv"
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
	Turn         uint8
	CastleRights uint8
	EpSquare     uint8
	halfMove     uint8
	fullMove     uint16
}

// Initialize board state with a FEN string
func (bs *BoardState) InitBoardState(FEN string) {
	// Split up the FEN string
	FENfields := strings.Fields(FEN)
	squares := FENfields[0]
	turn := FENfields[1]
	castleRights := FENfields[2]
	epSquare := FENfields[3]
	halfmove := FENfields[4]
	fullmove := FENfields[5]

	// Initialize Position
	bs.Position.SetPositionWithFEN(squares)

	// Initialize Turn
	if turn == "w" {
		bs.Turn = White
	} else {
		bs.Turn = Black
	}

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
	if epSquare == "-" {
		bs.EpSquare = NoEpSquare // No en passant square
	} else {
		// Convert the epSquare string to a uint8 value
		epFile := epSquare[0] - 'a' // Convert file letter to index
		epRank, _ := strconv.Atoi(string(epSquare[1]))
		bs.EpSquare = uint8(epRank-1)*8 + uint8(epFile)
	}

	// Initialize half move counter
	halfMoveInt, _ := strconv.ParseUint(halfmove, 10, 16)
	bs.halfMove = uint8(halfMoveInt)

	// Initialize full move counter
	fullMoveInt, _ := strconv.ParseUint(fullmove, 10, 16)
	bs.fullMove = uint16(fullMoveInt)
}
