package engine

import (
	"fmt"
)

// Here is an example of what a 64-bit number looks like:
// 0000000000000000000000000000000000000000000000000000000000000000
// When used to represent a chess state, this is called a bitboard
type Bitboard uint64

func (b *Bitboard) SetBit(square uint8) {
	*b |= 1 << square
}

// Returns 1 if there is a 1 on the square bit
// Returns 0 if there is a 0 on the square bit
func (b *Bitboard) GetBit(square uint8) bool {
	// Create a mask with a 1 on the square-indexed bit
	var mask Bitboard
	mask.SetBit(square)

	// Use bitwise AND to check if the bit at 'square' position is set
	return (*b & mask) != 0
}

func (b *Bitboard) PopBit(square int) {
	*b &= ^(1 << square)
}

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
