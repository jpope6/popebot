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

			if isSquareAttacked(square, bs) {
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
	if color == White {
		return "White"
	}
	return "Black"
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

	if epSquare != NoEpSquare {
		epSquareStr = squareToString(epSquare)
	} else {
		epSquareStr = "None"
	}

	return epSquareStr
}
