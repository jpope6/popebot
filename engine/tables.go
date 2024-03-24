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

var bishopAttacks [64][512]Bitboard
var rookAttacks [64][4096]uint64

func MaskBishopAttacks(square uint8) Bitboard {
	var attacks Bitboard

	var rank uint8 = GetRank(square)
	var file uint8 = GetFile(square)

	for r, f := rank, file; r < Rank7 && f < FileG; r, f = r+1, f+1 {
		attacks.SetBit(GetSquare(r+1, f+1))
	}

	for r, f := rank, file; r < Rank7 && f > FileB; r, f = r+1, f-1 {
		attacks.SetBit(GetSquare(r+1, f-1))
	}

	for r, f := rank, file; r > Rank2 && f < FileG; r, f = r-1, f+1 {
		attacks.SetBit(GetSquare(r-1, f+1))
	}

	for r, f := rank, file; r > Rank2 && f > FileB; r, f = r-1, f-1 {
		attacks.SetBit(GetSquare(r-1, f-1))
	}

	return attacks
}

func MaskBishopAttacksWithBlockers(square uint8, blockers Bitboard) Bitboard {
	var attacks Bitboard

	var rank uint8 = GetRank(square)
	var file uint8 = GetFile(square)

	for r, f := rank, file; r < Rank8 && f < FileH; r, f = r+1, f+1 {
		square := GetSquare(r+1, f+1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	for r, f := rank, file; r < Rank8 && f > FileA; r, f = r+1, f-1 {
		square := GetSquare(r+1, f-1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	for r, f := rank, file; r > Rank1 && f < FileH; r, f = r-1, f+1 {
		square := GetSquare(r-1, f+1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	for r, f := rank, file; r > Rank1 && f > FileA; r, f = r-1, f-1 {
		square := GetSquare(r-1, f-1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	return attacks
}

func MaskRookAttacks(square uint8) Bitboard {
	var attacks Bitboard

	var rank uint8 = GetRank(square)
	var file uint8 = GetFile(square)

	for r := rank; r < Rank7; r++ {
		attacks.SetBit(GetSquare(r+1, file))
	}

	for r := rank; r > Rank2; r-- {
		attacks.SetBit(GetSquare(r-1, file))
	}

	for f := file; f < FileG; f++ {
		attacks.SetBit(GetSquare(rank, f+1))
	}

	for f := file; f > FileB; f-- {
		attacks.SetBit(GetSquare(rank, f-1))
	}

	return attacks
}

func MaskRookAttacksWithBlockers(square uint8, blockers Bitboard) Bitboard {
	var attacks Bitboard

	var rank uint8 = GetRank(square)
	var file uint8 = GetFile(square)

	for r := rank; r < Rank8; r++ {
		square := GetSquare(r+1, file)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	for r := rank; r > Rank1; r-- {
		square := GetSquare(r-1, file)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	for f := file; f < FileH; f++ {
		square := GetSquare(rank, f+1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	for f := file; f > FileA; f-- {
		square := GetSquare(rank, f-1)
		attacks.SetBit(square)

		if blockers.GetBit(square) {
			break
		}
	}

	return attacks
}
