package engine

// Here is an example of what a 64-bit number looks like:
// 0000000000000000000000000000000000000000000000000000000000000000
// When used to represent a chess state, this is called a bitboard
type Bitboard uint64

func (b *Bitboard) SetBit(square uint8) {
	*b |= 1 << square
}

func (b *Bitboard) ClearBit(square int) {
	*b &= ^(1 << square)
}
