package engine

type Position struct {
	Pieces         [2][6]Bitboard
	AllWhitePieces Bitboard
	AllBlackPieces Bitboard
	AllPieces      Bitboard
}

type Piece struct {
	Type  uint8
	Color uint8
}

var byteToPieceMap = map[byte]Piece{
	'P': {Type: Pawn, Color: White},
	'N': {Type: Knight, Color: White},
	'B': {Type: Bishop, Color: White},
	'R': {Type: Rook, Color: White},
	'Q': {Type: Queen, Color: White},
	'K': {Type: King, Color: White},
	'p': {Type: Pawn, Color: Black},
	'n': {Type: Knight, Color: Black},
	'b': {Type: Bishop, Color: Black},
	'r': {Type: Rook, Color: Black},
	'q': {Type: Queen, Color: Black},
	'k': {Type: King, Color: Black},
}

// Input: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
// The function takes in PART of the FEN string
// It will set all bitboards in a Position struct to
// match the FEN string
func (position *Position) SetPositionWithFEN(FENsquares string) {
	// Initialize each piece bitboard based on the current byte
	// of the FEN string
	var index uint8 = 56
	for i := 0; i < len(FENsquares); i++ {
		byte := FENsquares[i]
		switch byte {
		case 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k':
			piece := byteToPieceMap[byte]
			position.Pieces[piece.Color][piece.Type].SetBit(index)
			index++
		case '1', '2', '3', '4', '5', '6', '7', '8':
			offset := uint8(byte - '0')
			index += offset
		case '/':
			index -= 16
		}
	}

	// Initialize AllWhitePieces bitboard
	// AllWhitePieces is the OR bitwise operator of all WhitePiece Bitboards
	for piece := 0; piece < len(position.Pieces[White]); piece++ {
		position.AllWhitePieces |= position.Pieces[White][piece]
	}

	// Initialize AllBlackPieces bitboard
	// AllBlackPieces is the OR bitwise operator of all BlackPiece Bitboards
	for piece := 0; piece < len(position.Pieces[Black]); piece++ {
		position.AllBlackPieces |= position.Pieces[Black][piece]
	}

	// Initialize AllPieces bitboard
	// AllPieces is the OR bitwise operator of AllWhitePieces and AllBlackPieces
	position.AllPieces = position.AllWhitePieces | position.AllBlackPieces
}
