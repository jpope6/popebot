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
	// Reset the board state
	bs.reset()

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
		bs.EpSquare = NoSquare // No en passant square
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

// Make move takes an encoded move and a move flag (quiet vs capture)
// and will make the move on a copy of the board
func (bs *BoardState) makeMove(move EncodedMove, moveFlag uint8) {
	if moveFlag == AllMoves || move.isCapture() {
		// copyBs := bs.copy()

		source := move.getSourceSquare()
		target := move.getTargetSquare()
		piece := move.getPiece()
		promotedPiece := move.getPromotedPiece()
		capture := move.isCapture()
		doublePush := move.isDoublePush()
		enPassant := move.isEnPassant()
		castle := move.isCastle()

		pieceColor := piece / 6
		pieceType := piece % 6

		// Make the move
		bs.Position.Pieces[pieceColor][pieceType].PopBit(source)
		bs.Position.Pieces[pieceColor][pieceType].SetBit(target)

		if capture {
			bs.handleCapture(target)
		}

		if promotedPiece != NoPiece {
			bs.handlePromotion(piece, promotedPiece, target)
		}

		if enPassant {
			bs.handleEnPassant(target)
		}

		// Reset En Passnt square
		bs.EpSquare = NoSquare

		if doublePush {
			bs.handleDoublePush(target)
		}

		if castle {
			bs.handleCastle(target)
		}

		bs.updateCastleRights(source, target)
		bs.updateBitboards()
	}
}

// handleCapture will remove the captured piece from it's
// respective bitboard at the target square
func (bs *BoardState) handleCapture(target uint8) {
	var start uint8
	var end uint8

	switch bs.Turn {
	case White:
		start = p
		end = k
	case Black:
		start = P
		end = K
	}

	for piece := start; piece <= end; piece++ {
		pieceColor := piece / 6
		pieceType := piece % 6

		// If there is a piece on target square, get rid of it
		if bs.Position.Pieces[pieceColor][pieceType].GetBit(target) {
			bs.Position.Pieces[pieceColor][pieceType].PopBit(target)
			break
		}
	}
}

func (bs *BoardState) handlePromotion(piece, promotedPiece, target uint8) {
	pieceColor := piece / 6
	pieceType := piece % 6
	promotedColor := promotedPiece / 6
	promotedType := promotedPiece % 6

	bs.Position.Pieces[pieceColor][pieceType].PopBit(target)
	bs.Position.Pieces[promotedColor][promotedType].SetBit(target)
}

func (bs *BoardState) handleEnPassant(target uint8) {
	if bs.Turn == White {
		bs.Position.Pieces[Black][Pawn].PopBit(target - 8)
	} else {
		bs.Position.Pieces[White][Pawn].PopBit(target + 8)
	}
}

func (bs *BoardState) handleDoublePush(target uint8) {
	if bs.Turn == White {
		bs.EpSquare = target - 8
	} else {
		bs.EpSquare = target + 8
	}
}

func (bs *BoardState) handleCastle(target uint8) {
	switch target {
	// White King Side
	case G1:
		bs.Position.Pieces[White][Rook].PopBit(H1)
		bs.Position.Pieces[White][Rook].SetBit(F1)
	// White Queen Side
	case C1:
		bs.Position.Pieces[White][Rook].PopBit(A1)
		bs.Position.Pieces[White][Rook].SetBit(D1)
		// Black King Side
	case G8:
		bs.Position.Pieces[Black][Rook].PopBit(H8)
		bs.Position.Pieces[Black][Rook].SetBit(F8)
	// Black Queen Side
	case C8:
		bs.Position.Pieces[Black][Rook].PopBit(A8)
		bs.Position.Pieces[Black][Rook].SetBit(D8)
	}
}

func (bs *BoardState) updateCastleRights(source, target uint8) {
	bs.CastleRights &= castleRights[source]
	bs.CastleRights &= castleRights[target]
}

func (bs *BoardState) updateBitboards() {
	bs.Position.AllWhitePieces = 0
	bs.Position.AllBlackPieces = 0
	bs.Position.AllPieces = 0

	// White Pieces
	for piece := P; piece <= K; piece++ {
		pieceColor := piece / 6
		pieceType := piece % 6
		bs.Position.AllWhitePieces |= bs.Position.Pieces[pieceColor][pieceType]
	}

	// Black Pieces
	for piece := P; piece <= K; piece++ {
		pieceColor := piece / 6
		pieceType := piece % 6
		bs.Position.AllBlackPieces |= bs.Position.Pieces[pieceColor][pieceType]
	}

	bs.Position.AllPieces = bs.Position.AllWhitePieces | bs.Position.AllBlackPieces
}

func (bs *BoardState) copy() *BoardState {
	var copyBs *BoardState = &BoardState{
		Position:     bs.Position,
		Turn:         bs.Turn,
		CastleRights: bs.CastleRights,
		EpSquare:     bs.EpSquare,
		halfMove:     bs.halfMove,
		fullMove:     bs.fullMove,
	}

	return copyBs
}

func (bs *BoardState) restore(other *BoardState) {
	*bs = *other
}

// Reset the Board State to a blank Board State
func (bs *BoardState) reset() {
	*bs = BoardState{}
}
