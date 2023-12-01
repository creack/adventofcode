package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Expected []int
	Lines    []string
}

func parse(fileName string) (*Data, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}
	d := &Data{}
	for _, line := range strings.Split(string(buf), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line %q", line)
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", line[1], err)
		}
		d.Expected = append(d.Expected, n)
		d.Lines = append(d.Lines, parts[0])
	}
	return d, nil
}

func run(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	// Part 1.
	// for i, line := range d.Lines {
	// 	buf := [4]byte{}
	// 	res := 0
	// 	for j := 0; j < len(line); j++ {
	// 		if j < 4 {
	// 			buf[j] = line[j]
	// 			continue
	// 		}
	// 		buf[0], buf[1], buf[2], buf[3] = buf[1], buf[2], buf[3], line[j]
	// 		m := map[byte]int{}
	// 		for _, elem := range buf {
	// 			m[elem]++
	// 		}
	// 		if len(m) == 4 {
	// 			res = j + 1
	// 			break
	// 		}
	// 	}
	// 	fmt.Printf("- %d\n", res)
	// 	if res != d.Expected[i] {
	// 		return fmt.Errorf("unexpected result %d != %d (%q)", res, d.Expected[i], line)
	// 	}
	// }

	// Part 2
	for i, line := range d.Lines {
		buf := [14]byte{}
		res := 0
		for j := 0; j < len(line); j++ {
			if j < len(buf) {
				buf[j] = line[j]
				continue
			}
			for k := 0; k+1 < len(buf); k++ {
				buf[k] = buf[k+1]
			}
			buf[len(buf)-1] = line[j]

			m := map[byte]int{}
			for _, elem := range buf {
				m[elem]++
			}
			if len(m) == len(buf) {
				res = j + 1
				break
			}
		}
		fmt.Printf("- %d\n", res)
		if res != d.Expected[i] {
			return fmt.Errorf("unexpected result %d != %d (%q)", res, d.Expected[i], line)
		}
	}

	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
