package engine

import (
	"fmt"
	"math"
)

func (bs *BoardState) Search(depth int) {
	var nodes Nodes = 0
	ply := 0
	_, bestMove := bs.negamax(depth, math.MinInt32, math.MaxInt32, &ply, &nodes)

	fmt.Printf("bestmove %s\n", bestMove.toUciMove())
}

func (bs *BoardState) negamax(
	depth, alpha, beta int, ply *int, nodes *Nodes,
) (int, EncodedMove) {
	if depth == 0 {
		return bs.quiescense(alpha, beta, ply, nodes), NoMove
	}

	*nodes++
	var bestMove EncodedMove
	originalAlpha := alpha

	legalMoves := 0

	moves := GenerateAllMoves(bs)
	moves.sortMoves(bs)

	isInCheck := bs.isKingInCheck()

	if isInCheck {
		depth++
	}

	for _, move := range moves.MoveList {
		if move == NoMove {
			continue
		}

		copyBs := bs.copy()
		*ply++

		// Ensure move is legal
		if !bs.MakeMove(move, AllMoves) {
			*ply--
			continue
		}

		legalMoves++

		score, _ := bs.negamax(depth-1, -beta, -alpha, ply, nodes)
		score = -score
		*ply--
		bs.restore(copyBs)

		if score >= beta {
			return beta, NoMove
		}

		if score > alpha {
			alpha = score

			if *ply == 0 {
				bestMove = move
			}
		}
	}

	// Checkmate or Stalemate
	if legalMoves == 0 {
		if isInCheck {
			// Checkmate score
			return math.MinInt32 + 1000 + *ply, NoMove
		} else {
			// Stalemate score
			return 0, NoMove
		}
	}

	if originalAlpha != alpha {
		return alpha, bestMove
	}

	return alpha, NoMove
}

func (bs *BoardState) quiescense(alpha, beta int, ply *int, nodes *Nodes) int {
	*nodes++
	evaluation := bs.Evaluate()

	if evaluation >= beta {
		return beta
	}

	if evaluation > alpha {
		alpha = evaluation
	}

	moves := GenerateAllMoves(bs)
	moves.sortMoves(bs)

	for _, move := range moves.MoveList {
		if move == NoMove {
			continue
		}

		copyBs := bs.copy()
		*ply++

		// Ensure move is legal
		if !bs.MakeMove(move, CaptureMoves) {
			*ply--
			continue
		}

		score := -bs.quiescense(-beta, -alpha, ply, nodes)
		bs.restore(copyBs)
		*ply--

		if score >= beta {
			return beta
		}

		if score > alpha {
			alpha = score
		}
	}

	return alpha
}

func (bs *BoardState) Evaluate() int {
	score := 0

	for piece := P; piece <= k; piece++ {
		pieceColor := piece / 6
		pieceType := piece % 6
		bb := bs.Position.Pieces[pieceColor][pieceType]

		for bb != 0 {
			square := bb.GetLsbIndex()

			score += mgMaterialScore[piece]
			score += getMgPsqtScore(piece, square)

			bb.PopBit(square)
		}
	}

	multiplier := 1
	if bs.Turn == Black {
		multiplier = -1
	}

	return score * multiplier
}

func getMgPsqtScore(piece, square uint8) int {
	switch piece {
	case P:
		return mgPawnTable[mirrorScore[square]]
	case N:
		return mgKnightTable[mirrorScore[square]]
	case B:
		return mgBishopTable[mirrorScore[square]]
	case R:
		return mgRookTable[mirrorScore[square]]
	case Q:
		return mgQueenTable[mirrorScore[square]]
	case K:
		return mgKingTable[mirrorScore[square]]
	case p:
		return -mgPawnTable[square]
	case n:
		return -mgKnightTable[square]
	case b:
		return -mgBishopTable[square]
	case r:
		return -mgRookTable[square]
	case q:
		return -mgQueenTable[square]
	case k:
		return -mgKingTable[square]
	}

	return 0
}
