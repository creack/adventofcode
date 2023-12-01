package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func test(ctx context.Context) error {
	buf, err := os.ReadFile("input")
	if err != nil {
		return fmt.Errorf("readFile: %w", err)
	}

	var elves [][]int
	var cur []int
	for _, line := range strings.Split(string(buf), "\n") {
		if line == "" {
			elves = append(elves, cur)
			cur = nil
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("atoi %q: %w", line, err)
		}
		cur = append(cur, n)
	}
	if len(cur) > 0 {
		elves = append(elves, cur)
	}

	var sums []int
	for _, elf := range elves {
		sum := 0
		for _, elem := range elf {
			sum += elem
		}
		sums = append(sums, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	if len(sums) < 3 {
		return fmt.Errorf("not enough sums (%d)", len(sums))
	}
	fmt.Printf("%d\n", sums[0]+sums[1]+sums[2])

	return nil
}

func main() {
	if err := test(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
