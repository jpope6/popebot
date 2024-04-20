package engine

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
		if !bs.MakeMove(moves.MoveList[moveCount], AllMoves) {
			continue
		}

		PerftDriver(bs, depth-1, nodes)

		// Take back
		bs.restore(boardStateCopy)
	}
}
