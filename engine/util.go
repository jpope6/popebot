package engine

func GetRank(square uint8) uint8 {
	return square >> 3
}

func GetFile(square uint8) uint8 {
	return square & 7
}

func GetSquare(rank uint8, file uint8) uint8 {
	return (rank << 3) + file
}
