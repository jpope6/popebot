package engine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

func UciCommand() {
	fmt.Printf("id name popebot\n")
	fmt.Printf("id author jpope6\n")
	fmt.Printf("uciok\n")
}

func UciLoop(bs *BoardState) {
	// Define reader for standard input
	reader := bufio.NewReader(os.Stdin)

	// Print engine info
	UciCommand()

	// Main loop
	for {
		// Get user/GUI input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim whitespace from input
		input = strings.TrimSpace(input)

		// Parse UCI commands
		switch {
		case input == "isready":
			fmt.Printf("readyok\n")
		case strings.HasPrefix(input, "position"):
			ParsePosition(bs, input)
		case input == "ucinewgame":
			ParsePosition(bs, "position startpos")
		case strings.HasPrefix(input, "go"):
			ParseGo(bs, input)
		case input == "quit":
			return
		}
	}
}

func ParseMoveString(bs *BoardState, moveStr string) EncodedMove {
	var file uint8
	var rank uint8
	var uciMove string

	// Source square
	file = uint8(moveStr[0] - 'a')
	rank = uint8(moveStr[1] - '1')
	uciMove += squareToString(GetSquare(rank, file))

	// Target square
	file = uint8(moveStr[2] - 'a')
	rank = uint8(moveStr[3] - '1')
	uciMove += squareToString(GetSquare(rank, file))

	// Promotion piece
	if len(moveStr) > 4 {
		uciMove += string(moveStr[4])
	}

	moves := GenerateAllMoves(bs)

	for _, move := range moves.MoveList {
		if move.toUciMove() == uciMove {
			return move
		}
	}

	// Illegal move
	return 0
}

func ParsePosition(bs *BoardState, command string) {
	// Split the command string by space
	tokens := strings.Fields(command)

	// Find the index of "fen" token
	fenIndex := -1
	for i, token := range tokens {
		if token == "fen" {
			fenIndex = i
			break
		}
	}

	// Extract FEN string
	var fen string
	if fenIndex != -1 {
		fen = strings.Join(tokens[fenIndex+1:fenIndex+7], " ")
	} else {
		fen = startFen
	}

	// Initialize BoardState with FEN string
	bs.InitBoardState(fen)

	// Find and parse moves if available
	movesIndex := -1
	for i, token := range tokens {
		if token == "moves" {
			movesIndex = i
			break
		}
	}

	if movesIndex != -1 {
		// Parse moves
		for _, moveStr := range tokens[movesIndex+1:] {
			move := ParseMoveString(bs, moveStr)

			// Stop parsing moves if no more valid moves
			if move == 0 {
				break
			}

			bs.MakeMove(move, AllMoves)
		}
	}
}

func ParseGo(bs *BoardState, command string) {
	// Init depth
	depth := -1

	// Find index of "depth" substring
	depthIndex := strings.Index(command, "depth")

	// Handle fixed depth search
	if depthIndex != -1 {
		// Extract the depth value after "depth"
		depthStr := command[depthIndex+len("depth"):]

		// Convert string to integer
		depthInt, err := strconv.Atoi(strings.TrimSpace(depthStr))
		if err == nil {
			depth = depthInt
		}
	} else {
		depth = 6 // Default depth
	}

	bs.Search(depth)
}
