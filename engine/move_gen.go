package engine

import (
	// "bufio"
	"fmt"
	// "os"
)

type Moves struct {
	MoveList [256]EncodedMove
	Count    int
}

type Nodes uint64

func PerftDriver(bs *BoardState, depth int, nodes *Nodes) {
	if depth == 0 {
		*nodes++
		return
	}

	moves := GenerateAllMoves(bs)

	// Loop over generated moves
	for moveCount := 0; moveCount < moves.Count; moveCount++ {
		// Preserve board state
		boardStateCopy := bs.copy()

		// Make move
		if !bs.makeMove(moves.MoveList[moveCount], AllMoves) {
			continue
		}
		// PrintBoard(bs)
		// fmt.Println("Press Enter to continue...")
		// bufio.NewReader(os.Stdin).ReadBytes('\n')

		PerftDriver(bs, depth-1, nodes)

		// Take back
		bs.restore(boardStateCopy)
		// PrintBoard(bs)
		// fmt.Println("Press Enter to continue...")
		// bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func (moves *Moves) addMove(
	source, target, piece, promotedPiece, capture, dp, ep, castle uint8,
) {
	// Create a Move object with provided parameters.
	move := Move{
		Source:         source,
		Target:         target,
		Piece:          piece,
		PromotedPiece:  promotedPiece,
		CaptureFlag:    capture,
		DoublePushFlag: dp,
		EnPassantFlag:  ep,
		CastleFlag:     castle,
	}

	// Encode the move.
	encodedMove := move.encodeMove()

	// Add the encoded move to the MoveList.
	moves.MoveList[moves.Count] = encodedMove

	// Increment the count of moves.
	moves.Count++
}

func (moves *Moves) printMoveList() {
	// for i := 0; i < moves.Count; i++ {
	// 	moves.MoveList[i].printUciMove()
	// }

	fmt.Printf("Total number of moves: %d\n", moves.Count)
}

func GenerateAllMoves(bs *BoardState) *Moves {
	var moves Moves = Moves{}

	moves.generatePawnMoves(bs)
	moves.generateMoves(bs, Knight)
	moves.generateMoves(bs, Bishop)
	moves.generateMoves(bs, Rook)
	moves.generateMoves(bs, Queen)
	moves.generateMoves(bs, King)
	moves.generateCastlingMoves(bs)

	return &moves
}

func (moves *Moves) generatePawnMoves(bs *BoardState) {
	var source uint8
	var target uint8
	var bb Bitboard

	switch bs.Turn {
	case White:
		bb = bs.Position.Pieces[White][Pawn]
		for bb != 0 {
			source = bb.GetLsbIndex()
			target = source + 8
			handlePawnMoves(bs, moves, P, source, target, target+8)
			handlePawnCaptures(bs, moves, P, source, target)
			bb.PopBit(source)
		}

	case Black:
		bb = bs.Position.Pieces[Black][Pawn]
		for bb != 0 {
			source = bb.GetLsbIndex()
			target = source - 8
			handlePawnMoves(bs, moves, p, source, target, target-8)
			handlePawnCaptures(bs, moves, p, source, target)
			bb.PopBit(source)
		}
	}
}

func handlePawnMoves(
	bs *BoardState, moves *Moves, piece, source, target, doubleTarget uint8,
) {
	if target > H8 || bs.Position.AllPieces.GetBit(target) {
		return // Square is off board or occupied
	}

	// Pawn promotion
	if isPromotionSquare(bs.Turn, source) {
		promotionPiece := piece + 1
		stopPiece := piece + 5

		// Loop from knight to queen
		for promotionPiece < stopPiece {
			moves.addMove(
				source, target, piece, promotionPiece, NoFlag, NoFlag, NoFlag, NoFlag,
			)
			promotionPiece++
		}
	} else {
		// Single pawn push
		moves.addMove(
			source, target, piece, NoPiece, NoFlag, NoFlag, NoFlag, NoFlag,
		)

		// Double pawn push
		if isDoublePushSquare(bs.Turn, source) && !bs.Position.AllPieces.GetBit(doubleTarget) {
			moves.addMove(
				source, doubleTarget, piece, NoPiece, NoFlag, DoublePushFlag, NoFlag, NoFlag,
			)
		}
	}
}

func handlePawnCaptures(
	bs *BoardState, moves *Moves, piece, source, target uint8,
) {
	var otherPieces Bitboard

	if bs.Turn == White {
		otherPieces = bs.Position.AllBlackPieces
	} else {
		otherPieces = bs.Position.AllWhitePieces
	}

	attacks := pawnAttacks[bs.Turn][source] & otherPieces

	for attacks != 0 {
		target = attacks.GetLsbIndex()

		if isPromotionSquare(bs.Turn, source) {
			promotionPiece := piece + 1
			stopPiece := piece + 5

			// Loop from knight to queen
			for promotionPiece < stopPiece {
				moves.addMove(
					source, target, piece, promotionPiece, CaptureFlag, NoFlag, NoFlag, NoFlag,
				)
				promotionPiece++
			}
		} else {
			moves.addMove(
				source, target, piece, NoPiece, CaptureFlag, NoFlag, NoFlag, NoFlag,
			)
		}

		attacks.PopBit(target)
	}

	// TODO: Might be able to move to to outside of loop in GeneratePawnMoves
	if bs.EpSquare != NoSquare {
		epAttacks := pawnAttacks[bs.Turn][source] & (1 << bs.EpSquare)

		if epAttacks != 0 {
			target := epAttacks.GetLsbIndex()
			moves.addMove(
				source, target, piece, NoPiece, CaptureFlag, NoFlag, EnPassantFlag, NoFlag,
			)
		}
	}
}

func (moves *Moves) generateMoves(bs *BoardState, pieceType uint8) {
	var source uint8
	var target uint8
	var piece uint8
	var bb Bitboard
	var attacks Bitboard
	var availableMoves Bitboard
	var otherPieces Bitboard

	if bs.Turn == White {
		piece = (White * NumPieces) + pieceType
		bb = bs.Position.Pieces[White][pieceType]

		// NOT White Pieces
		// availableMoves = Empty sqaures and squares with Black piece
		availableMoves = ^bs.Position.AllWhitePieces

		otherPieces = bs.Position.AllBlackPieces
	} else { // Black
		piece = (Black * NumPieces) + pieceType
		bb = bs.Position.Pieces[Black][pieceType]

		// NOT Black Pieces
		// availableMoves = Empty sqaures and squares with White piece
		availableMoves = ^bs.Position.AllBlackPieces

		otherPieces = bs.Position.AllWhitePieces
	}

	for bb != 0 {
		source = bb.GetLsbIndex()

		// Get the available moves of the piece type at the source square
		attacks = getMoves(bs, pieceType, source) & availableMoves

		for attacks != 0 {
			target = attacks.GetLsbIndex()

			// Capture moves
			if otherPieces.GetBit(target) {
				moves.addMove(
					source, target, piece, NoPiece, CaptureFlag, NoFlag, NoFlag, NoFlag,
				)
			} else {
				moves.addMove(
					source, target, piece, NoPiece, NoFlag, NoFlag, NoFlag, NoFlag,
				)
			}

			attacks.PopBit(target)
		}

		bb.PopBit(source)
	}
}

func (moves *Moves) generateCastlingMoves(bs *BoardState) {
	switch bs.Turn {
	case White:
		// King side
		if canCastle(bs, WhiteKingSide) {
			moves.addMove(
				E1, G1, K, NoPiece, NoFlag, NoFlag, NoFlag, CastleFlag,
			)
		}

		// Queen side
		if canCastle(bs, WhiteQueenSide) {
			moves.addMove(
				E1, C1, K, NoPiece, NoFlag, NoFlag, NoFlag, CastleFlag,
			)
		}

	case Black:
		// King side
		if canCastle(bs, BlackKingSide) {
			moves.addMove(
				E8, G8, k, NoPiece, NoFlag, NoFlag, NoFlag, CastleFlag,
			)
		}

		// Queen side
		if canCastle(bs, BlackQueenSide) {
			moves.addMove(
				E8, C8, k, NoPiece, NoFlag, NoFlag, NoFlag, CastleFlag,
			)
		}
	}
}

// Returns true if the square is currently attacked, else False
// NOTE: we do not check Queens because bishop and rook essentially does it already
func isSquareAttacked(bs *BoardState, square uint8) bool {
	bb := &bs.Position.Pieces

	switch bs.Turn {
	case White:
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

		// Attacked by Black Queen
		if getQueenAttacks(bs.Position.AllPieces, square)&bb[Black][Queen] != 0 {
			return true
		}

		// Attacked by Black King
		if kingAttacks[square]&bb[Black][King] != 0 {
			return true
		}

	case Black:
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

		// Attacked by White Queen
		if getQueenAttacks(bs.Position.AllPieces, square)&bb[White][Queen] != 0 {
			return true
		}

		// Attacked by White King
		if kingAttacks[square]&bb[White][King] != 0 {
			return true
		}

	}

	return false
}
