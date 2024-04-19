package engine

import "fmt"

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

func (move EncodedMove) printUciMove() {
	source := squareToString(move.getSourceSquare())
	target := squareToString(move.getTargetSquare())

	promotedPiece := move.getPromotedPiece()
	var promotedChar byte
	if promotedPiece != NoPiece {
		promotedChar = pieceToChar(promotedPiece)
	}

	fmt.Printf("%s%s%c\n", source, target, promotedChar)
}
