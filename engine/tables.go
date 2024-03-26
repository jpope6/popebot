package engine

// File constants
const (
	FileA = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
	NumFiles // Total number of files
)

// Rank constants
const (
	Rank1 = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	NumRanks // Total number of ranks
)

// Look up tables for piece attacks
var bishopMasks [64]Bitboard
var bishopAttacks [64][512]Bitboard

var rookMasks [64]Bitboard
var rookAttacks [64][4096]Bitboard

// Bishop relevant occupancy bit count for every square on board
var bishopRelevantBits = [64]int{
	6, 5, 5, 5, 5, 5, 5, 6,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	6, 5, 5, 5, 5, 5, 5, 6,
}

// Rook relevant occupancy bit count for every square on board
var rookRelevantBits = [64]int{
	12, 11, 11, 11, 11, 11, 11, 12,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	12, 11, 11, 11, 11, 11, 11, 12,
}

func InitTables() {
	initBishopTable()
	initRookTable()
}

func initBishopTable() {
	for square := uint8(0); square < 64; square++ {
		attackMask := maskBishopAttacks(square)
		bitCount := attackMask.CountBits()
		occupancyIndices := (1 << bitCount)

		for i := 0; i < occupancyIndices; i++ {
			occupancy := setOccupancy(i, attackMask)
			magicIndex := getBishopMagicIndex(&occupancy, square)
			bishopAttacks[square][magicIndex] = maskBishopAttacksWithBlockers(square, occupancy)
		}
	}
}

func initRookTable() {
	for square := uint8(0); square < 64; square++ {
		attackMask := maskRookAttacks(square)
		bitCount := attackMask.CountBits()
		occupancyIndices := (1 << bitCount)

		for i := 0; i < occupancyIndices; i++ {
			occupancy := setOccupancy(i, attackMask)
			magicIndex := getRookMagicIndex(&occupancy, square)
			rookAttacks[square][magicIndex] = maskRookAttacksWithBlockers(square, occupancy)
		}
	}
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

func maskRookAttacksWithBlockers(square uint8, blockers Bitboard) Bitboard {
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
