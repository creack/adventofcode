package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	From, To int
}

func parse() (any, error) {
	buf, err := os.ReadFile("input.base")
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	total := 0
	for _, line := range strings.Split(string(buf), "\n") {
		if line == "" {
			continue
		}
		elvesData := strings.Split(line, ",")
		if len(elvesData) != 2 {
			return nil, fmt.Errorf("invalid line %q", line)
		}
		var elfA Elf
		partsA := strings.Split(elvesData[0], "-")
		if len(partsA) != 2 {
			return nil, fmt.Errorf("invalid partA %q", partsA)
		}
		{
			from, err := strconv.Atoi(partsA[0])
			if err != nil {
				return nil, fmt.Errorf("invalid number %q partA from: %w", partsA[0], err)
			}
			to, err := strconv.Atoi(partsA[1])
			if err != nil {
				return nil, fmt.Errorf("invalid number %q partA to: %w", partsA[1], err)
			}
			elfA = Elf{From: from, To: to}
		}

		var elfB Elf
		partsB := strings.Split(elvesData[1], "-")
		if len(partsB) != 2 {
			return nil, fmt.Errorf("invalid partB %q", partsA)
		}
		{
			from, err := strconv.Atoi(partsB[0])
			if err != nil {
				return nil, fmt.Errorf("invalid number %q partA from: %w", partsA[0], err)
			}
			to, err := strconv.Atoi(partsB[1])
			if err != nil {
				return nil, fmt.Errorf("invalid number %q partA to: %w", partsA[1], err)
			}
			elfB = Elf{From: from, To: to}
		}

		// Part 1.
		// if (elfA.From >= elfB.From && elfA.To <= elfB.To) || (elfB.From >= elfA.From && elfB.To <= elfA.To) {
		// 	total++
		// }
		// Part 2.
		if (elfA.From >= elfB.From && elfA.From <= elfB.To) || (elfB.From >= elfA.From && elfB.From <= elfA.To) {
			total++
		}
	}

	return total, nil
}

func test(ctx context.Context) error {
	total, err := parse()
	if err != nil {
		return fmt.Errorf("parse: %w", err)
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
