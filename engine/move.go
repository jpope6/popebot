package engine

type Move struct {
	Source         uint8
	Target         uint8
	Piece          uint8
	PromotedPiece  uint8
	CaptureFlag    uint8
	DoublePushFlag uint8
	EnPassantFlag  uint8
	CastleFlag     uint8
}

type EncodedMove uint64

func (move Move) encodeMove() EncodedMove {
	var encoded EncodedMove

	encoded |= EncodedMove(move.Source)
	encoded |= EncodedMove(uint64(move.Target) << 6)
	encoded |= EncodedMove(uint64(move.Piece) << 12)
	encoded |= EncodedMove(uint64(move.PromotedPiece) << 16)

	// Combine the flags using bitwise OR operations
	flags := uint64(move.CaptureFlag | move.DoublePushFlag | move.EnPassantFlag | move.CastleFlag)

	// Shift the combined flags to their correct positions and add them to the encoded move
	encoded |= EncodedMove(flags << 20)

	return encoded
}

func (move *EncodedMove) ScoreMove(bs *BoardState) int {
	if move.isCapture() {
		var attacker uint8 = move.getPiece()
		var target uint8 = Pawn

		var start uint8
		var end uint8

		switch bs.Turn {
		case White:
			start = p
			end = k
		case Black:
			start = P
			end = K
		}

		for piece := start; piece <= end; piece++ {
			pieceColor := piece / 6
			pieceType := piece % 6

			// If there is a piece on target square, that is the target piece
			if bs.Position.Pieces[pieceColor][pieceType].GetBit(move.getTargetSquare()) {
				target = piece
				break
			}
		}

		return MvvLva[attacker%6][target%6]
	}

	return 0
}

func (move EncodedMove) getSourceSquare() uint8 {
	return uint8(move & SourceSquareHex)
}

func (move EncodedMove) getTargetSquare() uint8 {
	return uint8(move & TargetSquareHex >> 6)
}

func (move EncodedMove) getPiece() uint8 {
	return uint8(move & PieceHex >> 12)
}

func (move EncodedMove) getPromotedPiece() uint8 {
	return uint8(move & PromotedHex >> 16)
}

func (move EncodedMove) isCapture() bool {
	return (move & CaptureHex) != 0
}

func (move EncodedMove) isDoublePush() bool {
	return (move & DoublePushHex) != 0
}

func (move EncodedMove) isEnPassant() bool {
	return (move & EnPassantHex) != 0
}

func (move EncodedMove) isCastle() bool {
	return (move & CastleHex) != 0
}

func (move EncodedMove) toUciMove() string {
	var uciMove string

	uciMove += squareToString(move.getSourceSquare())
	uciMove += squareToString(move.getTargetSquare())

	promotedPiece := move.getPromotedPiece()
	if promotedPiece != NoPiece {
		uciMove += string(pieceToChar(promotedPiece))
	}

	return uciMove
}
