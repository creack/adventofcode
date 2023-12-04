package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Results []int
	Lines   []string
}

func parse(fileName string) (*Data, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}
	d := &Data{}

	// Split lines, populate data.
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		d.Lines = append(d.Lines, line)
	}

	// Parse each line.
	for _, line := range d.Lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line")
		}
		leftParts := strings.Split(parts[0], ":")
		if len(leftParts) != 2 {
			return nil, fmt.Errorf("invalid line: invalid left side")
		}
		left := leftParts[1]
		right := parts[1]

		leftNums := map[int]int{}
		for _, elem := range strings.Split(left, " ") {
			if elem == "" {
				continue
			}
			n, err := strconv.Atoi(elem)
			if err != nil {
				return nil, fmt.Errorf("atoi %q: %w", elem, err)
			}
			leftNums[n] = n
		}

		rightNums := map[int]int{}
		for _, elem := range strings.Split(right, " ") {
			if elem == "" {
				continue
			}
			n, err := strconv.Atoi(elem)
			if err != nil {
				return nil, fmt.Errorf("atoi %q: %w", elem, err)
			}
			rightNums[n] = n
		}

		winCount := 0
		for l := range leftNums {
			for r := range rightNums {
				if r != l {
					continue
				}
				winCount++
			}
		}
		d.Results = append(d.Results, winCount)
	}

	return d, nil
}

func run(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	part1, part2 := 0, 0
	copies := map[int][]int{}
	for i, winCount := range d.Results {
		// Part 1.
		if winCount > 0 {
			acc := 1
			for j := 1; j < winCount; j++ {
				acc *= 2
			}
			part1 += acc
		}

		// Part 2.
		hdlr := func(i, winCount int) {
			part2 += winCount
			for j := 1; j <= winCount; j++ {
				if i+j > len(d.Results) {
					continue
				}
				copies[i+j] = append(copies[i+j], d.Results[i+j])
			}
		}
		// Handle the original.
		hdlr(i, winCount)

		// Handle staged copies.
		for n, lines := range copies {
			if n > i {
				continue
			}
			for _, line := range lines {
				hdlr(n, line)
			}
			delete(copies, n)
		}
	}

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2+len(d.Lines))
	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
