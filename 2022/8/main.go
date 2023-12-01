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

func get(tab [][]int, x, y int) int {
	for i := 0; i < len(tab); i++ {
		line := tab[i]
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

	cases := map[string]int{}

	width := len(d.Lines[0])
	height := len(d.Lines)
	_, _ = width, height

	// Looking from the top.
	for i := 0; i < width; i++ {
	out:
		for j := 0; j < height; j++ {
			if i == 0 || j == 0 || i == width-1 || j == height-1 {
				continue
			}

			cur := get(d.Lines, j, i)
			for k := j - 1; k >= 0; k-- {
				if get(d.Lines, k, i) >= cur {
					continue out
				}
			}

			cases[fmt.Sprintf("%d/%d", i, j)] = get(d.Lines, i, j)
		}
	}

	// Looking from the right.
	for i := 0; i < width; i++ {
	out2:
		for j := 0; j < height; j++ {
			if i == 0 || j == 0 || i == width-1 || j == height-1 {
				continue
			}

			cur := get(d.Lines, j, i)
			for k := i + 1; k < width; k++ {
				if get(d.Lines, j, k) >= cur {
					continue out2
				}
			}

			cases[fmt.Sprintf("%d/%d", i, j)] = get(d.Lines, i, j)
		}
	}

	// Looking from the bottom.
	for i := 0; i < width; i++ {
	out3:
		for j := 0; j < height; j++ {
			if i == 0 || j == 0 || i == width-1 || j == height-1 {
				continue
			}

			cur := get(d.Lines, j, i)
			for k := j + 1; k < height; k++ {
				if get(d.Lines, k, i) >= cur {
					continue out3
				}
			}

			cases[fmt.Sprintf("%d/%d", i, j)] = get(d.Lines, i, j)
		}
	}

	// Looking from the left.
	for i := 0; i < width; i++ {
	out4:
		for j := 0; j < height; j++ {
			if i == 0 || j == 0 || i == width-1 || j == height-1 {
				continue
			}

			cur := get(d.Lines, j, i)
			for k := i - 1; k >= 0; k-- {
				if get(d.Lines, j, k) >= cur {
					continue out4
				}
			}

			cases[fmt.Sprintf("%d/%d", i, j)] = get(d.Lines, i, j)
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
		base := get(d.Lines, x, y)
		// Looking up.
		up := 0
		for j := y - 1; j >= 0; j-- {
			cur := get(d.Lines, x, j)
			up++
			if cur >= base {
				break
			}
		}
		// Looking down.
		down := 0
		for j := y + 1; j < height; j++ {
			cur := get(d.Lines, x, j)
			down++
			if cur >= base {
				break
			}
		}
		// Looking left.
		left := 0
		for j := x - 1; j >= 0; j-- {
			cur := get(d.Lines, j, y)
			left++
			if cur >= base {
				break
			}
		}
		// Looking right.
		right := 0
		for j := x + 1; j < width; j++ {
			cur := get(d.Lines, j, y)
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
