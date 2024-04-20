package engine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

type PerftTest struct {
	FEN         string
	DepthValues [7]Nodes
}

func loadPerftData() (tests []PerftTest, err error) {
	filePath := "../perft_data/perftsuite.epd"

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		test := PerftTest{FEN: strings.TrimSpace(parts[0])}
		test.DepthValues = [7]Nodes{}

		for _, count := range parts[1:] {
			depth, _ := strconv.Atoi(string(count[1]))
			expected, _ := strconv.Atoi(strings.TrimSpace(count[3:]))

			test.DepthValues[depth-1] = Nodes(expected)
		}

		tests = append(tests, test)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tests, nil
}

func TestPerftDriver(t *testing.T) {
	fmt.Printf("Starting test...\n")
	perftData, err := loadPerftData()
	if err != nil {
		t.Errorf("Error loading perft data: %v", err)
		return
	}

	InitTables()
	var bs BoardState

	for _, tests := range perftData {
		for depth, expectedValue := range tests.DepthValues {
			if expectedValue == 0 {
				continue
			}

			bs.InitBoardState(tests.FEN)
			var nodes Nodes = 0

			PerftDriver(&bs, depth+1, &nodes)

			if nodes != expectedValue {
				fmt.Printf("Position: %s\n Depth: %d, Expected Perft Value: %d, Actual Perft Value: %d\n",
					tests.FEN, depth+1, expectedValue, nodes)
			}
		}

	}

	fmt.Println("Test finished.")
}
