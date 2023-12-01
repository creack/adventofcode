package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// Rock: A - 1 - X
// Paper: B - 2 - Y
// Scissors: C - 3 - Z
//
// Win: 6
// Draw: 3
// Lose: 0

const (
	ScoreWin  = 6
	ScoreDraw = 3
	ScoreLose = 0

	Rock     = 1
	Paper    = 2
	Scissors = 3
)

var winMap = map[int]map[int]int{
	Rock: {
		Rock:     ScoreDraw,
		Paper:    ScoreWin,
		Scissors: ScoreLose,
	},
	Paper: {
		Rock:     ScoreLose,
		Paper:    ScoreDraw,
		Scissors: ScoreWin,
	},
	Scissors: {
		Rock:     ScoreWin,
		Paper:    ScoreLose,
		Scissors: ScoreDraw,
	},
}

var stratWinMap = map[int]map[int]int{
	Rock: {
		ScoreDraw: Rock,
		ScoreWin:  Paper,
		ScoreLose: Scissors,
	},
	Paper: {
		ScoreLose: Rock,
		ScoreDraw: Paper,
		ScoreWin:  Scissors,
	},
	Scissors: {
		ScoreWin:  Rock,
		ScoreLose: Paper,
		ScoreDraw: Scissors,
	},
}

var stratMoveMap = map[string]int{
	"X": ScoreLose,
	"Y": ScoreDraw,
	"Z": ScoreWin,
}

var moveMap = map[string]int{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func play(a, b int) int {
	return winMap[a][b] + b
}

func test(ctx context.Context) error {
	buf, err := os.ReadFile("input")
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}

	total := 0
	for _, line := range strings.Split(string(buf), "\n") {
		parts := strings.Split(line, " ")
		if line == "" || len(parts) < 2 {
			continue
		}

		first, ok := moveMap[parts[0]]
		if !ok {
			return fmt.Errorf("unknown first move %q", parts[0])
		}

		// Part 1:
		// sec, ok := moveMap[parts[1]]
		// if !ok {
		// 	return fmt.Errorf("unknown sec move %q", parts[1])
		// }
		// total += play(first, sec)

		// Part 2:
		sec, ok := stratMoveMap[parts[1]]
		if !ok {
			return fmt.Errorf("unknown sec strat move %q", parts[1])
		}
		total += play(first, stratWinMap[first][sec])
	}
	fmt.Printf("total: %d\n", total)

	return nil
}

func main() {
	if err := test(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
