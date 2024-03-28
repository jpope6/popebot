package engine

// Look up tables for piece attacks
var pawnAttacks [2][64]Bitboard // [color][square]

var knightAttacks [64]Bitboard // [square]

var bishopMasks [64]Bitboard        // [square]
var bishopAttacks [64][512]Bitboard // [square][occupancy]

var rookMasks [64]Bitboard         // [square]
var rookAttacks [64][4096]Bitboard // [square][occupancy]

func InitTables() {
	initPawnTable()
	initKnightTable()
	initBishopTable()
	initRookTable()
}

// Initialize pawnAttacks
func initPawnTable() {
	// Loop through each square on the board
	for square := uint8(0); square < 64; square++ {
		// Set the square for White and Black attacks
		pawnAttacks[White][square] = maskPawnAttacks(White, square)
		pawnAttacks[Black][square] = maskPawnAttacks(Black, square)
	}
}

func maskPawnAttacks(side uint8, square uint8) Bitboard {
	var attacks Bitboard

	var board Bitboard = 1 << square

	if side == White { // White Pawn
		// Check if pawn is not on File A
		if board&NotFileA != 0 {
			// North West
			attacks.SetBit(square + 7)
		}

		// Check if pawn is not on File H
		if board&NotFileH != 0 {
			// North East
			attacks.SetBit(square + 9)
		}
	} else { // Black pawn
		// Check if pawn is not on File H
		if board&NotFileH != 0 {
			// South West
			attacks.SetBit(square - 7)
		}

		// Check if pawn is not on File A
		if board&NotFileA != 0 {
			// South East
			attacks.SetBit(square - 9)
		}
	}

	return attacks
}

// Initialize knightAttacks
func initKnightTable() {
	// Loop through each square on the board
	for square := uint8(0); square < 64; square++ {
		// Set the square for Knight attacks
		knightAttacks[square] = maskKnightAttacks(square)
	}
}

func maskKnightAttacks(square uint8) Bitboard {
	var attacks Bitboard

	var board Bitboard = 1 << square

	if board&NotFileA != 0 {
		attacks.SetBit(square + 15)
		attacks.SetBit(square - 17)
	}

	if board&NotFileH != 0 {
		attacks.SetBit(square + 17)
		attacks.SetBit(square - 15)
	}

	if board&NotFileAB != 0 {
		attacks.SetBit(square + 6)
		attacks.SetBit(square - 10)
	}

	if board&NotFileHG != 0 {
		attacks.SetBit(square + 10)
		attacks.SetBit(square - 6)
	}

	return attacks
}

// Initialize bishopsMasks and bishopAttacks
func initBishopTable() {
	// Loop through each square on the board
	for square := uint8(0); square < 64; square++ {
		// Get the attack mask for the current square
		attackMask := maskBishopAttacks(square)

		// Set the attack mask for the current square
		bishopMasks[square] = attackMask

		// Calculate the number of posssible occupancy varitiations
		// based on the relevant bits
		occupancyIndices := (1 << bishopRelevantBits[square])

		// Loop through each possible occupancy variation
		for i := 0; i < occupancyIndices; i++ {
			occupancy := setOccupancy(i, attackMask)
			magicIndex := getBishopMagicIndex(occupancy, square)

			// Store the attack mask for the current occupancy and square
			bishopAttacks[square][magicIndex] = maskBishopAttacksWithBlockers(square, occupancy)
		}
	}
}

func maskBishopAttacks(square uint8) Bitboard {
	var attacks Bitboard

	var rank uint8 = GetRank(square)
	var file uint8 = GetFile(square)

	// North East
	for r, f := rank, file; r < Rank7 && f < FileG; r, f = r+1, f+1 {
		attacks.SetBit(GetSquare(r+1, f+1))
	}

	// North West
	for r, f := rank, file; r < Rank7 && f > FileB; r, f = r+1, f-1 {
		attacks.SetBit(GetSquare(r+1, f-1))
	}

	// South East
	for r, f := rank, file; r > Rank2 && f < FileG; r, f = r-1, f+1 {
		attacks.SetBit(GetSquare(r-1, f+1))
	}

	// South West
	for r, f := rank, file; r > Rank2 && f > FileB; r, f = r-1, f-1 {
		attacks.SetBit(GetSquare(r-1, f-1))
	}

	return attacks
}

func maskBishopAttacksWithBlockers(square uint8, blockers Bitboard) Bitboard {
	var attacks Bitboard

	var rank uint8 = GetRank(square)
	var file uint8 = GetFile(square)

	// North East
	for r, f := rank, file; r < Rank8 && f < FileH; r, f = r+1, f+1 {
		square := GetSquare(r+1, f+1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	// North West
	for r, f := rank, file; r < Rank8 && f > FileA; r, f = r+1, f-1 {
		square := GetSquare(r+1, f-1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	// South East
	for r, f := rank, file; r > Rank1 && f < FileH; r, f = r-1, f+1 {
		square := GetSquare(r-1, f+1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	// South West
	for r, f := rank, file; r > Rank1 && f > FileA; r, f = r-1, f-1 {
		square := GetSquare(r-1, f-1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	return attacks
}

// Initialize rookMasks and rookAttacks
func initRookTable() {
	// Loop through each square on the board
	for square := uint8(0); square < 64; square++ {
		// Get the attack mask for the current square
		attackMask := maskRookAttacks(square)

		// Set the attack mask for the current square
		rookMasks[square] = attackMask

		// Calculate the number of posssible occupancy varitiations
		// based on the relevant bits
		occupancyIndices := (1 << rookRelevantBits[square])

		// Loop through each possible occupancy variation
		for i := 0; i < occupancyIndices; i++ {
			occupancy := setOccupancy(i, attackMask)
			magicIndex := getRookMagicIndex(occupancy, square)

			// Store the attack mask for the current occupancy and square
			rookAttacks[square][magicIndex] = MaskRookAttacksWithBlockers(square, occupancy)
		}
	}
}

func maskRookAttacks(square uint8) Bitboard {
	var attacks Bitboard

	var rank uint8 = GetRank(square)
	var file uint8 = GetFile(square)

	// North
	for r := rank; r < Rank7; r++ {
		attacks.SetBit(GetSquare(r+1, file))
	}

	// South
	for r := rank; r > Rank2; r-- {
		attacks.SetBit(GetSquare(r-1, file))
	}

	// East
	for f := file; f < FileG; f++ {
		attacks.SetBit(GetSquare(rank, f+1))
	}

	// West
	for f := file; f > FileB; f-- {
		attacks.SetBit(GetSquare(rank, f-1))
	}

	return attacks
}

func MaskRookAttacksWithBlockers(square uint8, blockers Bitboard) Bitboard {
	var attacks Bitboard

	var rank uint8 = GetRank(square)
	var file uint8 = GetFile(square)

	// North
	for r := rank; r < Rank8; r++ {
		square := GetSquare(r+1, file)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	// South
	for r := rank; r > Rank1; r-- {
		square := GetSquare(r-1, file)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	// East
	for f := file; f < FileH; f++ {
		square := GetSquare(rank, f+1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	// West
	for f := file; f > FileA; f-- {
		square := GetSquare(rank, f-1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	return attacks
}

func MaskQueenAttacks(square uint8, blockers Bitboard) Bitboard {
	return getBishopAttacks(blockers, square) | getRookAttacks(blockers, square)
}

// SetOccupancy generates an occupancy map based on the provided index and attack mask.
// It sets bits in the occupancy map corresponding to the set bits in the attack mask.
func setOccupancy(index int, attackMask Bitboard) Bitboard {
	// Initialize an empty occupancy map
	var occupancy Bitboard

	// Determine the number of set bits in the attack mask
	bitCount := int(attackMask.CountBits())

	// Iterate over the range of set bits in the attack mask
	for i := 0; i < bitCount; i++ {
		// Get the index of the least significant set bit in the attack mask
		square := attackMask.GetLsbIndex()

		// Clear the least significant set bit in the attack mask
		attackMask.PopBit(square)

		// Check if the corresponding bit is on the board
		if index&(1<<i) != 0 {
			// Set the corresponding bit in the occupancy map
			occupancy.SetBit(square)
		}
	}

	// Return the generated occupancy map
	return occupancy
}
