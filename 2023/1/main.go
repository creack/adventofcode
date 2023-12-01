package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
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

func runPart1(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	total := 0
	for _, line := range d.Lines {
		var left, right byte
	subloop:
		for _, elem := range line {
			if elem >= '0' && elem <= '9' {
				left = byte(elem)
				break subloop
			}
		}

	subloop2:
		for i := len(line) - 1; i >= 0; i-- {
			elem := line[i]
			if elem >= '0' && elem <= '9' {
				right = elem
				break subloop2
			}
		}
		n, err := strconv.Atoi(string([]byte{left, right}))
		if err != nil {
			return fmt.Errorf("atoi %q: %w", line, err)
		}
		total += n
	}
	fmt.Printf("%d\n", total)
	return nil
}

func runPart2(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	tokens := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	total := 0
	for _, line := range d.Lines {
		var left, right byte
	subloop:
		for i := 0; i < len(line); i++ {
			elem := line[i]

			if elem >= '0' && elem <= '9' {
				left = byte(elem)
				break subloop
			}

			for n, tok := range tokens {
				if i+len(tok) >= len(line) {
					continue
				}
				if line[i:i+len(tok)] == tok {
					left = byte(n+1) + '0'
					break subloop
				}
			}
		}

	subloop2:
		for i := len(line) - 1; i >= 0; i-- {
			elem := line[i]
			if elem >= '0' && elem <= '9' {
				right = elem
				break subloop2
			}

			for n, tok := range tokens {
				if i-len(tok) < 0 {
					continue
				}
				if line[i-len(tok)+1:i+1] == tok {
					right = byte(n+1) + '0'
					break subloop2
				}
			}
		}
		n, err := strconv.Atoi(string([]byte{left, right}))
		if err != nil {
			return fmt.Errorf("atoi %q: %w", line, err)
		}
		total += n
	}
	fmt.Printf("%d\n", total)
	return nil
}

func main() {
	if err := runPart2(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
