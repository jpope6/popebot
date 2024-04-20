package engine

const (
	// Constants for mapping board squares to numbers.
	A1, B1, C1, D1, E1, F1, G1, H1 = 0, 1, 2, 3, 4, 5, 6, 7
	A2, B2, C2, D2, E2, F2, G2, H2 = 8, 9, 10, 11, 12, 13, 14, 15
	A3, B3, C3, D3, E3, F3, G3, H3 = 16, 17, 18, 19, 20, 21, 22, 23
	A4, B4, C4, D4, E4, F4, G4, H4 = 24, 25, 26, 27, 28, 29, 30, 31
	A5, B5, C5, D5, E5, F5, G5, H5 = 32, 33, 34, 35, 36, 37, 38, 39
	A6, B6, C6, D6, E6, F6, G6, H6 = 40, 41, 42, 43, 44, 45, 46, 47
	A7, B7, C7, D7, E7, F7, G7, H7 = 48, 49, 50, 51, 52, 53, 54, 55
	A8, B8, C8, D8, E8, F8, G8, H8 = 56, 57, 58, 59, 60, 61, 62, 63

	// File constants
	FileA    = 0
	FileB    = 1
	FileC    = 2
	FileD    = 3
	FileE    = 4
	FileF    = 5
	FileG    = 6
	FileH    = 7
	NumFiles // Total number of files

	// Rank constants
	Rank1    = 0
	Rank2    = 1
	Rank3    = 2
	Rank4    = 3
	Rank5    = 4
	Rank6    = 5
	Rank7    = 6
	Rank8    = 7
	NumRanks // Total number of ranks

	// Edge of board constants
	NotFileA  Bitboard = 0xFEFEFEFEFEFEFEFE
	NotFileH  Bitboard = 0x7F7F7F7F7F7F7F7F
	NotFileHG Bitboard = 0x3F3F3F3F3F3F3F3F
	NotFileAB Bitboard = 0xFCFCFCFCFCFCFCFC
	NotRank1  Bitboard = 0xFFFFFFFFFFFFFF00
	NotRank8  Bitboard = 0x00FFFFFFFFFFFFFF

	// Piece Types
	Pawn      uint8 = 0
	Knight    uint8 = 1
	Bishop    uint8 = 2
	Rook      uint8 = 3
	Queen     uint8 = 4
	King      uint8 = 5
	NumPieces uint8 = 6

	// Color Types
	White uint8 = 0
	Black uint8 = 1

	// More Piece Types
	P uint8 = 0
	N uint8 = 1
	B uint8 = 2
	R uint8 = 3
	Q uint8 = 4
	K uint8 = 5
	p uint8 = 6
	n uint8 = 7
	b uint8 = 8
	r uint8 = 9
	q uint8 = 10
	k uint8 = 11

	NoColor uint8 = 2
	NoPiece uint8 = 12

	// Hexadecimal number that corresponds to castling rights
	WhiteKingSide  = 0x8 // 1000
	WhiteQueenSide = 0x4 // 0100
	BlackKingSide  = 0x2 // 0010
	BlackQueenSide = 0x1 // 0001

	// NoEpSquare will be num squares + 1
	NoSquare = 64

	// Encoding move constants
	SourceSquareHex = 0x3F
	TargetSquareHex = 0xFC0
	PieceHex        = 0xF000
	PromotedHex     = 0xF0000
	CaptureHex      = 0x100000
	DoublePushHex   = 0x200000
	EnPassantHex    = 0x400000
	CastleHex       = 0x800000

	// Flags
	NoFlag         = 0x0 // 0000
	CaptureFlag    = 0x1 // 0001
	DoublePushFlag = 0x2 // 0010
	EnPassantFlag  = 0x4 // 0100
	CastleFlag     = 0x8 // 1000

	// For the make move function
	AllMoves     uint8 = 0
	CaptureMoves uint8 = 1
)

var pieceSymbols = [2][6]rune{
	{'\u265F', '\u265E', '\u265D', '\u265C', '\u265B', '\u265A'}, // White
	{'\u2659', '\u2658', '\u2657', '\u2656', '\u2655', '\u2654'}, // Black
}

// Bishop number of attacked squares for every square on board
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

// Rook number of attacked squares for every square on board
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

// Decimal numbers to update castle rights based on the square
// of the peice that has moved
var castleRights = [64]uint8{
	11, 15, 15, 15, 3, 15, 15, 7,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	14, 15, 15, 15, 12, 15, 15, 13,
}

var bishopMagics [64]Bitboard = [64]Bitboard{
	0x40040844404084,
	0x2004208a004208,
	0x10190041080202,
	0x108060845042010,
	0x581104180800210,
	0x2112080446200010,
	0x1080820820060210,
	0x3c0808410220200,
	0x4050404440404,
	0x21001420088,
	0x24d0080801082102,
	0x1020a0a020400,
	0x40308200402,
	0x4011002100800,
	0x401484104104005,
	0x801010402020200,
	0x400210c3880100,
	0x404022024108200,
	0x810018200204102,
	0x4002801a02003,
	0x85040820080400,
	0x810102c808880400,
	0xe900410884800,
	0x8002020480840102,
	0x220200865090201,
	0x2010100a02021202,
	0x152048408022401,
	0x20080002081110,
	0x4001001021004000,
	0x800040400a011002,
	0xe4004081011002,
	0x1c004001012080,
	0x8004200962a00220,
	0x8422100208500202,
	0x2000402200300c08,
	0x8646020080080080,
	0x80020a0200100808,
	0x2010004880111000,
	0x623000a080011400,
	0x42008c0340209202,
	0x209188240001000,
	0x400408a884001800,
	0x110400a6080400,
	0x1840060a44020800,
	0x90080104000041,
	0x201011000808101,
	0x1a2208080504f080,
	0x8012020600211212,
	0x500861011240000,
	0x180806108200800,
	0x4000020e01040044,
	0x300000261044000a,
	0x802241102020002,
	0x20906061210001,
	0x5a84841004010310,
	0x4010801011c04,
	0xa010109502200,
	0x4a02012000,
	0x500201010098b028,
	0x8040002811040900,
	0x28000010020204,
	0x6000020202d0240,
	0x8918844842082200,
	0x4010011029020020,
}

var rookMagics [64]Bitboard = [64]Bitboard{
	0x8a80104000800020,
	0x140002000100040,
	0x2801880a0017001,
	0x100081001000420,
	0x200020010080420,
	0x3001c0002010008,
	0x8480008002000100,
	0x2080088004402900,
	0x800098204000,
	0x2024401000200040,
	0x100802000801000,
	0x120800800801000,
	0x208808088000400,
	0x2802200800400,
	0x2200800100020080,
	0x801000060821100,
	0x80044006422000,
	0x100808020004000,
	0x12108a0010204200,
	0x140848010000802,
	0x481828014002800,
	0x8094004002004100,
	0x4010040010010802,
	0x20008806104,
	0x100400080208000,
	0x2040002120081000,
	0x21200680100081,
	0x20100080080080,
	0x2000a00200410,
	0x20080800400,
	0x80088400100102,
	0x80004600042881,
	0x4040008040800020,
	0x440003000200801,
	0x4200011004500,
	0x188020010100100,
	0x14800401802800,
	0x2080040080800200,
	0x124080204001001,
	0x200046502000484,
	0x480400080088020,
	0x1000422010034000,
	0x30200100110040,
	0x100021010009,
	0x2002080100110004,
	0x202008004008002,
	0x20020004010100,
	0x2048440040820001,
	0x101002200408200,
	0x40802000401080,
	0x4008142004410100,
	0x2060820c0120200,
	0x1001004080100,
	0x20c020080040080,
	0x2935610830022400,
	0x44440041009200,
	0x280001040802101,
	0x2100190040002085,
	0x80c0084100102001,
	0x4024081001000421,
	0x20030a0244872,
	0x12001008414402,
	0x2006104900a0804,
	0x1004081002402,
}
