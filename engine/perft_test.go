package engine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type PerftTest struct {
	FEN        string
	DepthNodes [7]Nodes
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
		test.DepthNodes = [7]Nodes{}

		for _, count := range parts[1:] {
			depth, _ := strconv.Atoi(string(count[1]))
			expected, _ := strconv.Atoi(strings.TrimSpace(count[3:]))

			test.DepthNodes[depth-1] = Nodes(expected)
		}

		tests = append(tests, test)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tests, nil
}

func TestPerftDriver(t *testing.T) {
	fmt.Printf("Starting test...\n\n")
	perftData, err := loadPerftData()
	if err != nil {
		t.Errorf("Error loading perft data: %v", err)
		return
	}

	InitTables()
	var bs BoardState

	for _, tests := range perftData {
		fmt.Printf("Position: %s\n", tests.FEN)
		for depth, expectedValue := range tests.DepthNodes {
			if expectedValue == 0 {
				continue
			}

			start := time.Now()

			bs.InitBoardState(tests.FEN)
			var nodes Nodes = 0

			PerftDriver(&bs, depth+1, &nodes)

			duration := time.Since(start).Milliseconds()

			if nodes != expectedValue {
				fmt.Printf("Depth: %d, Expected Perft Value: %d, Actual Perft Value: %d\n",
					depth+1, expectedValue, nodes)
				t.Errorf("Position failed: %s, Depth: %d\n", tests.FEN, depth+1)
			} else {
				fmt.Printf("Depth: %d, Time taken: %d ms\n", depth+1, duration)
			}
		}
		fmt.Printf("\n----------------------------------------------------------------\n\n")
	}

	fmt.Println("Test finished.")
}
