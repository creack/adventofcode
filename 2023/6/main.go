package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func (r Race) candidates() int {
	n := 0
	for i := 0; i < r.Time; i++ {
		if r.Distance < i*(r.Time-i) {
			n++
		}
	}
	return n
}

type Data struct {
	Races []Race
}

func parse(fileName string, squash bool) (*Data, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}
	d := &Data{}
	lines := strings.Split(string(buf), "\n")
	var times, distances []int
	for i, line := range lines {
		if line == "" {
			continue
		}
		// Part 2: squash the line into one number.
		if squash {
			line = strings.ReplaceAll(line, " ", "")
		}
		var numbers []int
		for _, elem := range strings.Split(strings.Split(line, ":")[1], " ") {
			if elem == "" {
				continue
			}
			n, _ := strconv.Atoi(elem)
			numbers = append(numbers, n)
		}
		if i == 0 {
			times = numbers
			continue
		}
		distances = numbers
	}

	for i, t := range times {
		d.Races = append(d.Races, Race{Time: t, Distance: distances[i]})
	}
	return d, nil
}

func run(fileName string, squash bool) (int, error) {
	d, err := parse(fileName, squash)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}

	n := 1
	for _, elem := range d.Races {
		n *= elem.candidates()
	}
	return n, nil
}

func runAll(_ context.Context) error {
	for _, elem := range []string{"input.base", "input"} {
		part1, err := run(elem, false)
		if err != nil {
			return fmt.Errorf("run %q part 1: %w", elem, err)
		}
		part2, err := run(elem, true)
		if err != nil {
			return fmt.Errorf("run %q part 1: %w", elem, err)
		}

		fmt.Printf("% 10s: part1: % 10d, part2: % 10d\n", elem, part1, part2)
	}
	return nil
}

func main() {
	if err := runAll(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
