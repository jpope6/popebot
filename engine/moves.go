package engine

import "fmt"

func generateMoves(bs *BoardState) {
}

func GeneratePawnMoves(bs *BoardState) {
	var sourceSquare uint8
	var targetSquare uint8
	var bb Bitboard

	switch bs.Turn {
	case White:
		bb = bs.Position.Pieces[White][Pawn]
		for bb != 0 {
			sourceSquare = bb.GetLsbIndex()
			targetSquare = sourceSquare + 8
			handlePawnMoves(bs, sourceSquare, targetSquare, targetSquare+8)
			handlePawnCaptures(bs, sourceSquare, targetSquare)
			bb.PopBit(sourceSquare)
		}
	case Black:
		bb = bs.Position.Pieces[Black][Pawn]
		for bb != 0 {
			sourceSquare = bb.GetLsbIndex()
			targetSquare = sourceSquare - 8
			handlePawnMoves(bs, sourceSquare, targetSquare, targetSquare-8)
			handlePawnCaptures(bs, sourceSquare, targetSquare)
			bb.PopBit(sourceSquare)
		}
	}
}

func handlePawnMoves(bs *BoardState, sourceSquare, targetSquare, doubleTargetSquare uint8) {
	// If square is on board and there is not a piece on the targetSquare
	if targetSquare <= H8 && !bs.Position.AllPieces.GetBit(targetSquare) {
		// Pawn promotion
		if isPromotionSquare(bs.Turn, sourceSquare) {
			fmt.Printf("Pawn Promote: %s, %s\n", squareToString(sourceSquare), squareToString(targetSquare))
		} else {
			// Single pawn push
			fmt.Printf("Pawn single push: %s, %s\n", squareToString(sourceSquare), squareToString(targetSquare))

			// Double pawn push
			if isDoublePushSquare(bs.Turn, sourceSquare) &&
				!bs.Position.AllPieces.GetBit(doubleTargetSquare) {
				fmt.Printf("Pawn double push: %s, %s\n", squareToString(sourceSquare), squareToString(doubleTargetSquare))
			}
		}
	}
}

func handlePawnCaptures(bs *BoardState, sourceSquare, targetSquare uint8) {
	var otherPieces Bitboard

	if bs.Turn == White {
		otherPieces = bs.Position.AllBlackPieces
	} else {
		otherPieces = bs.Position.AllWhitePieces
	}

	attacks := pawnAttacks[bs.Turn][sourceSquare] & otherPieces

	for attacks != 0 {
		targetSquare = attacks.GetLsbIndex()

		if isPromotionSquare(bs.Turn, sourceSquare) {
			fmt.Printf("Pawn Capture Promote: %s, %s\n", squareToString(sourceSquare), squareToString(targetSquare))
		} else {
			fmt.Printf("Pawn Capture: %s, %s\n", squareToString(sourceSquare), squareToString(targetSquare))
		}

		attacks.PopBit(targetSquare)
	}

	// TODO: Might be able to move to to outside of loop in GeneratePawnMoves
	if bs.EpSquare != NoEpSquare {
		epAttacks := pawnAttacks[bs.Turn][sourceSquare] & (1 << bs.EpSquare)

		if epAttacks != 0 {
			targetEpSquare := epAttacks.GetLsbIndex()
			fmt.Printf("Pawn epCapture: %s, %s\n", squareToString(sourceSquare), squareToString(targetEpSquare))
		}
	}
}

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
