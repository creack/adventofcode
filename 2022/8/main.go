package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Lines [][]int
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
		row := make([]int, len(lines[0]))
		for i := 0; i < len(row); i++ {
			row[i] = int(line[i]) - '0'
		}
		d.Lines = append(d.Lines, row)
	}
	return d, nil
}

func (d *Data) get(x, y int) int {
	for i := 0; i < len(d.Lines); i++ {
		line := d.Lines[i]
		for j := 0; j < len(line); j++ {
			elem := line[j]
			if i == x && j == y {
				return elem
			}
		}
	}
	return -1
}

func runPart1(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	width := len(d.Lines[0])
	height := len(d.Lines)

	top := func(cur, i, j int) int {
		for k := j - 1; k >= 0; k-- {
			if d.get(i, k) >= cur {
				return -1
			}
		}
		return cur
	}
	bottom := func(cur, i, j int) int {
		for k := j + 1; k < height; k++ {
			if d.get(i, k) >= cur {
				return -1
			}
		}
		return cur
	}
	right := func(cur, i, j int) int {
		for k := i + 1; k < width; k++ {
			if d.get(k, j) >= cur {
				return -1
			}
		}
		return cur
	}
	left := func(cur, i, j int) int {
		for k := i - 1; k >= 0; k-- {
			if d.get(k, j) >= cur {
				return -1
			}
		}
		return cur
	}

	cases := map[string]int{}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if i == 0 || j == 0 || i == width-1 || j == height-1 {
				continue
			}

			for _, elem := range []func(cur, i, j int) int{top, bottom, right, left} {
				if v := elem(d.get(i, j), i, j); v != -1 {
					cases[fmt.Sprintf("%d/%d", i, j)] = v
				}
			}
		}
	}

	fmt.Printf("vt: %d\n", len(cases)+width*2+height*2-4)
	return nil
}

func runPart2(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	width := len(d.Lines[0])
	height := len(d.Lines)

	getCount := func(x, y int) int {
		base := d.get(x, y)
		// Looking up.
		up := 0
		for j := y - 1; j >= 0; j-- {
			cur := d.get(x, j)
			up++
			if cur >= base {
				break
			}
		}
		// Looking down.
		down := 0
		for j := y + 1; j < height; j++ {
			cur := d.get(x, j)
			down++
			if cur >= base {
				break
			}
		}
		// Looking left.
		left := 0
		for j := x - 1; j >= 0; j-- {
			cur := d.get(j, y)
			left++
			if cur >= base {
				break
			}
		}
		// Looking right.
		right := 0
		for j := x + 1; j < width; j++ {
			cur := d.get(j, y)
			right++
			if cur >= base {
				break
			}
		}

		return up * down * left * right
		// fmt.Printf("%d %d %d %d = %d (%d)\n", up, left, down, right, up*down*left*right, base)
	}

	var maxScore, maxI, maxJ int
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			v := getCount(i, j)
			if v > maxScore {
				maxScore = v
				maxI = i
				maxJ = j
			}
		}
	}

	fmt.Printf("%d/%d, %d\n", maxI, maxJ, getCount(maxI, maxJ))
	return nil
}

func main() {
	if err := runPart1(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	if err := runPart2(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
