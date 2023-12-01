package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Lines []string
}

func parse(fileName string) (*Data, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}
	d := &Data{}
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		d.Lines = append(d.Lines, line)
	}
	return d, nil
}

func run(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	part1Tokens := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	_ = part1Tokens
	part2Tokens := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
	_ = part2Tokens

	tokens := part2Tokens

	total := 0
	for _, line := range d.Lines {
		var left, right int
	subloop:
		for i := 0; i < len(line); i++ {
			for tok, n := range tokens {
				if i+len(tok) > len(line) {
					continue
				}
				if line[i:i+len(tok)] == tok {
					left = n
					break subloop
				}
			}
		}

	subloop2:
		for i := len(line) - 1; i >= 0; i-- {
			for tok, n := range tokens {
				if i-len(tok)+1 < 0 {
					continue
				}
				if line[i-len(tok)+1:i+1] == tok {
					right = n
					break subloop2
				}
			}
		}
		total += left*10 + right
	}
	fmt.Printf("%d\n", total)
	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
