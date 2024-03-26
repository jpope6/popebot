package engine

import "fmt"

func GetRank(square uint8) uint8 {
	return square >> 3
}

func GetFile(square uint8) uint8 {
	return square & 7
}

func GetSquare(rank uint8, file uint8) uint8 {
	return (rank << 3) + file
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
