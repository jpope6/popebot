package engine

// Returns true if the square is currently attacked, else False
// NOTE: we do not check Queens because bishop and rook essentially does it already
func isSquareAttacked(square uint8, bs *BoardState) bool {
	bb := &bs.Position.Pieces

	switch bs.Turn {
	case White:
		// Attacked by White Pawns
		if pawnAttacks[Black][square]&bb[White][Pawn] != 0 {
			return true
		}

		// Attacked by White Knights
		if knightAttacks[square]&bb[White][Knight] != 0 {
			return true
		}

		// Attacked by White Bishop
		if getBishopAttacks(bs.Position.AllPieces, square)&bb[White][Bishop] != 0 {
			return true
		}

		// Attacked by White Rook
		if getRookAttacks(bs.Position.AllPieces, square)&bb[White][Rook] != 0 {
			return true
		}

		// Attacked by White King
		if kingAttacks[square]&bb[White][King] != 0 {
			return true
		}

	case Black:
		if pawnAttacks[White][square]&bb[Black][Pawn] != 0 {
			return true
		}

		// Attacked by Black Knight
		if knightAttacks[square]&bb[Black][Knight] != 0 {
			return true
		}

		// Attacked by Black Bishop
		if getBishopAttacks(bs.Position.AllPieces, square)&bb[Black][Bishop] != 0 {
			return true
		}

		// Attacked by Black Rook
		if getRookAttacks(bs.Position.AllPieces, square)&bb[Black][Rook] != 0 {
			return true
		}

		// Attacked by Black King
		if kingAttacks[square]&bb[Black][King] != 0 {
			return true
		}
	}

	return false
}
