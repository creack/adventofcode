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

type Node struct {
	x, y      int
	Number    int
	Neighbors []*Node
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

func isSymbol(c byte) bool {
	if c == 0 {
		return false
	}
	if c >= '0' && c <= '9' {
		return false
	}
	if c == '.' {
		return false
	}
	return true
}

func (d *Data) get(x, y int) byte {
	if y < 0 || x < 0 {
		return 0
	}
	if y >= len(d.Lines) {
		return 0
	}
	if x >= len(d.Lines[y]) {
		return 0
	}
	return d.Lines[y][x]
}

type symbol struct {
	val  byte
	x, y int
}

func (d *Data) checkAround(i, j, k int, numStr string) []symbol {
	var out []symbol

	if val := d.get(k, j-1); isSymbol(val) {
		out = append(out, symbol{val, k, j - 1})
	}
	if val := d.get(k, j+1); isSymbol(val) {
		out = append(out, symbol{val, k, j + 1})
	}
	if k == i-len(numStr) {
		if val := d.get(k-1, j-1); isSymbol(val) {
			out = append(out, symbol{val, k - 1, j - 1})
		}
		if val := d.get(k-1, j+1); isSymbol(val) {
			out = append(out, symbol{val, k - 1, j + 1})
		}
		if val := d.get(k-1, j); isSymbol(val) {
			out = append(out, symbol{val, k - 1, j})
		}
	}
	if k == i-1 {
		if val := d.get(k+1, j-1); isSymbol(val) {
			out = append(out, symbol{val, k + 1, j - 1})
		}
		if val := d.get(k+1, j+1); isSymbol(val) {
			out = append(out, symbol{val, k + 1, j + 1})
		}
		if val := d.get(k+1, j); isSymbol(val) {
			out = append(out, symbol{val, k + 1, j})
		}
	}
	return out
}

func run(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	starNodeIndex := map[string]*Node{}

	part1, part2 := 0, 0
	for j, line := range d.Lines {
		numStr := ""
		for i := 0; i < len(line); i++ {
			elem := line[i]
			if elem >= '0' && elem <= '9' {
				numStr += string(elem)

				// If we are not at the last char of a line, continue the loop.
				if i+1 != len(line) {
					continue
				}
				// If we are the last char, process numStr.
				i++
			}
			// If we don't have any digits buffered, continue.
			if numStr == "" {
				continue
			}

			// Parse the full number.
			n, err := strconv.Atoi(numStr)
			if err != nil {
				return fmt.Errorf("atoi %q: %w", elem, err)
			}
			// For over each digit of the number.
			for k := i - len(numStr); k < i; k++ {
				// Check around each digit to see if we have symbols.
				if val := d.checkAround(i, j, k, numStr); len(val) > 0 {
					// If we have at least one symbol, it is a part number.
					part1 += n

					// Part 2: Go over the symbols to handle '*'.
					for _, v := range val {
						if v.val != '*' {
							continue
						}
						// Lookup the node in the index (in case we have a common symbol in between parts).
						starNode := starNodeIndex[fmt.Sprintf("%d/%d", v.x, v.y)]
						if starNode == nil {
							// If missing, create it and store it in the index.
							starNode = &Node{
								Number: -1,
								x:      v.x,
								y:      v.y,
							}
							starNodeIndex[fmt.Sprintf("%d/%d", v.x, v.y)] = starNode
						}

						// Add the number as neighbor.
						starNode.Neighbors = append(starNode.Neighbors, &Node{Number: n, x: i, y: j})
					}
					break
				}
			}
			numStr = ""
		}
	}

	for _, elem := range starNodeIndex {
		if len(elem.Neighbors) == 2 {
			part2 += elem.Neighbors[0].Number * elem.Neighbors[1].Number
		}
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
