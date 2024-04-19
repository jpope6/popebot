package engine

import (
	"fmt"
)

func GetRank(square uint8) uint8 {
	return square >> 3
}

func GetFile(square uint8) uint8 {
	return square & 7
}

func GetSquare(rank uint8, file uint8) uint8 {
	return (rank << 3) + file
}

func isPromotionSquare(turn uint8, square uint8) bool {
	if turn == White {
		return square >= A7 && square <= H7
	} else {
		return square >= A2 && square <= H2
	}
}

func isDoublePushSquare(turn uint8, square uint8) bool {
	if turn == White {
		return square >= A2 && square <= H2
	} else {
		return square >= A7 && square <= H7
	}
}

func canCastle(bs *BoardState, side uint8) bool {
	switch side {

	case WhiteKingSide:
		return bs.CastleRights&WhiteKingSide != 0 &&
			!bs.Position.AllPieces.GetBit(F1) &&
			!bs.Position.AllPieces.GetBit(G1) &&
			!isSquareAttacked(bs, E1) &&
			!isSquareAttacked(bs, F1) &&
			!isSquareAttacked(bs, G1)

	case WhiteQueenSide:
		return bs.CastleRights&WhiteQueenSide != 0 &&
			!bs.Position.AllPieces.GetBit(D1) &&
			!bs.Position.AllPieces.GetBit(C1) &&
			!bs.Position.AllPieces.GetBit(B1) &&
			!isSquareAttacked(bs, E1) &&
			!isSquareAttacked(bs, D1) &&
			!isSquareAttacked(bs, C1)

	case BlackKingSide:
		return bs.CastleRights&BlackKingSide != 0 &&
			!bs.Position.AllPieces.GetBit(F8) &&
			!bs.Position.AllPieces.GetBit(G8) &&
			!isSquareAttacked(bs, E8) &&
			!isSquareAttacked(bs, F8) &&
			!isSquareAttacked(bs, G8)

	case BlackQueenSide:
		return bs.CastleRights&WhiteQueenSide != 0 &&
			!bs.Position.AllPieces.GetBit(D8) &&
			!bs.Position.AllPieces.GetBit(C8) &&
			!bs.Position.AllPieces.GetBit(B8) &&
			!isSquareAttacked(bs, E8) &&
			!isSquareAttacked(bs, D8) &&
			!isSquareAttacked(bs, C8)

	}

	return false
}

func isKingInCheck(bs *BoardState, kingSquare uint8) bool {
	return isSquareAttacked(bs, kingSquare)
}

// Print the bitboard in a chess board
func PrintBitboard(bb Bitboard) {
	fmt.Printf("\n")

	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			square := r*8 + f

			if f == 0 {
				fmt.Printf("  %d ", r+1)
			}

			if bb&(1<<square) != 0 {
				fmt.Printf("X ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n    a b c d e f g h \n \n")
}

// Prints the full board state of the chess game
func PrintBoard(boardState *BoardState) {
	fmt.Printf("\n")

	for r := 7; r >= 0; r-- {
		fmt.Printf("%d ", r+1)

		for f := 0; f < 8; f++ {
			square := GetSquare(uint8(r), uint8(f))
			var pieceFound bool

			for color, bbColor := range boardState.Position.Pieces {
				for piece, bbPiece := range bbColor {
					if bbPiece.GetBit(square) {
						fmt.Printf("%c ", pieceSymbols[color][piece])
						pieceFound = true
						break
					}
				}
				if pieceFound {
					break
				}
			}

			if !pieceFound {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n  a b c d e f g h\n\n")
	fmt.Printf("Turn:              %s\n", colorToString(boardState.Turn))
	fmt.Printf("Castle Rights:     %s\n", castleRightsToString(boardState.CastleRights))
	fmt.Printf("En Passant Square: %s\n", EnPassantToString(boardState.EpSquare))
}

func PrintAttackedSquares(bs *BoardState) {

	fmt.Printf("\n")

	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			square := uint8(r*8 + f)

			if f == 0 {
				fmt.Printf("  %d ", r+1)
			}

			if isSquareAttacked(bs, square) {
				fmt.Printf("X ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n    a b c d e f g h \n \n")
}

// Function to convert square index to file and rank notation
func squareToString(square uint8) string {
	file := 'a' + rune(square%8)
	rank := '1' + rune(square/8)
	return string(file) + string(rank)
}

// Function to convert color constant to string representation
func colorToString(color uint8) string {
	switch color {
	case White:
		return "White"
	case Black:
		return "Black"
	}

	return ""
}

func castleRightsToString(castleRights uint8) string {
	var castleRightsStr string

	if (castleRights & WhiteKingSide) != 0 {
		castleRightsStr += "K"
	} else {
		castleRightsStr += "-"
	}
	if (castleRights & WhiteQueenSide) != 0 {
		castleRightsStr += "Q"
	} else {
		castleRightsStr += "-"
	}
	if (castleRights & BlackKingSide) != 0 {
		castleRightsStr += "k"
	} else {
		castleRightsStr += "-"
	}
	if (castleRights & BlackQueenSide) != 0 {
		castleRightsStr += "q"
	} else {
		castleRightsStr += "-"
	}

	return castleRightsStr
}

func EnPassantToString(epSquare uint8) string {
	var epSquareStr string

	if epSquare != NoSquare {
		epSquareStr = squareToString(epSquare)
	} else {
		epSquareStr = "None"
	}

	return epSquareStr
}

func pieceToUint8(piece Piece) uint8 {
	if piece.Type == NoPiece {
		return NoPiece
	}

	return (piece.Color * NumPieces) + piece.Type
}

func pieceUint8ToString(piece uint8) string {
	if piece > 11 {
		return "None"
	}

	pieceColor := piece / 6
	pieceType := piece % 6

	return colorToString(pieceColor) + " " + pieceToString(pieceType)
}

func pieceToChar(piece uint8) byte {
	switch piece {
	case P, p:
		return 'p'
	case N, n:
		return 'n'
	case B, b:
		return 'b'
	case R, r:
		return 'r'
	case Q, q:
		return 'q'
	case K, k:
		return 'k'
	}

	return ' '
}

func pieceToString(pieceType uint8) string {
	switch pieceType {
	case Pawn:
		return "Pawn"
	case Knight:
		return "Knight"
	case Bishop:
		return "Bishop"
	case Rook:
		return "Rook"
	case Queen:
		return "Queen"
	case King:
		return "King"
	}

	return ""
}
