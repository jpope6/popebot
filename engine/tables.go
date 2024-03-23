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

	for r, f := rank+1, file+1; r < Rank8 && f < FileH; r, f = r+1, f+1 {
		attacks.SetBit(GetSquare(r, f))
	}

	for r, f := rank+1, file-1; r < Rank8 && f > FileA; r, f = r+1, f-1 {
		attacks.SetBit(GetSquare(r, f))
	}

	for r, f := rank-1, file+1; r > Rank1 && f < FileH; r, f = r-1, f+1 {
		attacks.SetBit(GetSquare(r, f))
	}

	for r, f := rank-1, file-1; r > Rank1 && f > FileA; r, f = r-1, f-1 {
		attacks.SetBit(GetSquare(r, f))
	}

	return attacks
}
