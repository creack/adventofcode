package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	N, From, To int
}

type Data struct {
	Stack [][]byte
	Moves []Move
	Lines []string
}

func (d *Data) DumpStack() {
	maxHeight := 0
	for _, elem := range d.Stack {
		if len(elem) > maxHeight {
			maxHeight = len(elem)
		}
	}
	for h := maxHeight - 1; h >= 0; h-- {
		for i := 0; i < len(d.Stack); i++ {
			if len(d.Stack[i]) < h+1 {
				fmt.Printf("    ")
				continue
			}
			fmt.Printf("[%c] ", d.Stack[i][h])
		}
		fmt.Printf("\n")
	}
	for i := 0; i < len(d.Stack); i++ {
		fmt.Printf(" %d  ", i+1)
	}
	fmt.Printf("\n")
}

func parse(fileName string) (*Data, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}
	lines := strings.Split(string(buf), "\n")

	// Find how tall the stack is.
	height := 0
	for _, line := range lines {
		if line == "" {
			break
		}
		lines = append(lines, line)
		height++
	}
	height--

	// FInd how wide the stack is.
	width := 0
	for _, elem := range strings.Split(lines[height], " ") {
		if elem == "" {
			continue
		}
		n, err := strconv.Atoi(elem)
		if err != nil {
			return nil, fmt.Errorf("atoi %q: %w", elem, err)
		}
		width = n
	}

	// Create a 2D slice representing the stack.
	stack := make([][]byte, width)
	for i := 0; i < width; i++ {
		for h := 1; h <= height; h++ {
			if len(lines[height-h]) < 4*i {
				break
			}
			if lines[height-h][1+i*4] == ' ' {
				break
			}
			stack[i] = append(stack[i], lines[height-h][1+i*4])
		}
	}

	// Parse the moves.
	var moves []Move
	section := 0
	for _, elem := range lines {
		if elem == "" {
			section++
			continue
		}
		if section != 1 {
			continue
		}
		var n, from, to int
		if _, err := fmt.Sscanf(elem, "move %d from %d to %d", &n, &from, &to); err != nil {
			return nil, fmt.Errorf("invalid move %q: %w", elem, err)
		}
		moves = append(moves, Move{N: n, From: from - 1, To: to - 1})
	}

	d := &Data{
		Stack: stack,
		Moves: moves,
		Lines: lines,
	}

	return d, nil
}

func run(_ context.Context) error {
	d, err := parse("input.base")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	// d.DumpStack()

	// Part 1:
	// for _, elem := range d.Moves {
	// 	for i := 0; i < elem.N; i++ {
	// 		val := d.Stack[elem.From][len(d.Stack[elem.From])-1]
	// 		d.Stack[elem.From] = d.Stack[elem.From][:len(d.Stack[elem.From])-1]
	// 		d.Stack[elem.To] = append(d.Stack[elem.To], val)
	// 	}
	// 	// d.DumpStack()
	// }

	// Part 2:
	for _, elem := range d.Moves {
		vals := d.Stack[elem.From][len(d.Stack[elem.From])-elem.N:]
		d.Stack[elem.From] = d.Stack[elem.From][:len(d.Stack[elem.From])-elem.N]
		d.Stack[elem.To] = append(d.Stack[elem.To], vals...)
		// d.DumpStack()
	}

	d.DumpStack()
	for _, elem := range d.Stack {
		fmt.Printf("%c", elem[len(elem)-1])
	}
	fmt.Printf("\n")
	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
}
