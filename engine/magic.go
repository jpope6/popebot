package engine

// Magic Bitboards is a multiply-right-shift hashing algorithm
// to index an attack bitboard look-up table.
//
// A magic move-bitboard generation consists of four steps:
//
// 1. Mask the occupancy bits to form a key
// 2. Multiply the key by a "magic number" to obtain an index mapping
// 3. Right shift the index mapping by (64 - n) bits to create an
//      index, when n is the number of bits in the index
// 4. Use the index to map the pre-initialized move
//
//
//                                         any consecutive
// relevant occupancy                      combination of
// rook a1, 12 bits                        the masked bits
// . . . . . . . .     . . . . . . . .     5 6 B C D E F G]
// 6 . . . . . . .     . . .some . . .     . . . .[1 2 3 4
// 5 . . . . . . .     . . . . . . . .     . . . . . . . .
// 4 . . . . . . .     . . .magic. . .     . . . . . . . .
// 3 . . . . . . .  *  . . . . . . . .  =  . . garbage . .    >> (64-12)
// 2 . . . . . . .     . . .bits . . .     . . . . . . . .
// 1 . . . . . . .     . . . . . . . .     . . . . . . . .
// . B C D E F G .     . . . . . . . .     . . . . . . . .

// getBishopMagicIndex computes the magic index for a bishop given the
// occupancy bitboard and the square.
func getBishopMagicIndex(occupancy Bitboard, square uint8) Bitboard {
	// Calculate the key by multiplying the occupancy with the
	// precomputed magic number for the given square
	key := (occupancy * bishopMagics[square])

	// Calculate the index by shifting the key to the
	// right by the difference between 64 and the number
	// of relevant bits for the square
	index := key >> (64 - bishopRelevantBits[square])

	// Return the computed index
	return index
}

// getBishopAttacks computes the attack bitboard for a bishop on
// a given square with a specific occupancy pattern.
func getBishopAttacks(occupancy Bitboard, square uint8) Bitboard {
	// Apply a bit mask to the occupancy bitboard to isolate
	// the relevant bits for the given square
	occupancy &= bishopMasks[square]

	// Multiply the masked occupancy with the precomputed
	// magic number for the square
	occupancy *= bishopMagics[square]

	// Extract the resulting attack bitboard by shifting
	// the occupancy to the right by the difference between
	// 64 and the number of relevant bits for the square
	occupancy >>= 64 - bishopRelevantBits[square]

	// Return the computed attack bitboard from the precomputed table
	return bishopAttacks[square][occupancy]
}

// getRookMagicIndex computes the magic index for a rook given the
// occupancy bitboard and the square.
func getRookMagicIndex(occupancy Bitboard, square uint8) Bitboard {
	// Calculate the key by multiplying the occupancy with the
	// precomputed magic number for the given square
	key := (occupancy * rookMagics[square])

	// Calculate the index by shifting the key to the
	// right by the difference between 64 and the number
	// of relevant bits for the square
	index := key >> (64 - rookRelevantBits[square])

	// Return the computed index
	return index
}

// getRookAttacks computes the attack bitboard for a rook on a
// given square with a specific occupancy pattern.
func getRookAttacks(occupancy Bitboard, square uint8) Bitboard {
	// Apply a bit mask to the occupancy bitboard to isolate
	// the relevant bits for the given square
	occupancy &= rookMasks[square]

	// Multiply the masked occupancy with the precomputed
	// magic number for the square
	occupancy *= rookMagics[square]

	// Extract the resulting attack bitboard by shifting
	// the occupancy to the right by the difference between
	// 64 and the number of relevant bits for the square
	occupancy >>= 64 - rookRelevantBits[square]

	// Return the computed attack bitboard from the precomputed table
	return rookAttacks[square][occupancy]
}
