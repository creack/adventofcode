package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID int

	Blue  int
	Red   int
	Green int
}

type Data struct {
	Lines []string

	Games []Game
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

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line, wrong part count")
		}
		gameNumberStr := strings.TrimPrefix(parts[0], "Game ")
		gameNumber, err := strconv.Atoi(gameNumberStr)
		if err != nil {
			return nil, fmt.Errorf("atoi %q: %w", gameNumberStr, err)
		}

		g := Game{}
		g.ID = gameNumber

		sets := strings.Split(parts[1], ";")
		for _, set := range sets {
			for _, elem := range strings.Split(set, ",") {
				tmp := strings.Split(strings.TrimSpace(elem), " ")

				n, err := strconv.Atoi(tmp[0])
				if err != nil {
					return nil, fmt.Errorf("atoi2 %q: %w", elem, err)
				}

				switch tmp[1] {
				case "red":
					if g.Red < n {
						g.Red = n
					}
				case "green":
					if g.Green < n {
						g.Green = n
					}
				case "blue":
					if g.Blue < n {
						g.Blue = n
					}
				}

			}
		}

		d.Games = append(d.Games, g)

	}
	return d, nil
}

func run(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	red := 12
	green := 13
	blue := 14

	total := 0
	for _, g := range d.Games {
		possible := g.Red <= red && g.Green <= green && g.Blue <= blue
		fmt.Printf("%d: %d, %d, %d -- %t\n", g.ID, g.Red, g.Green, g.Blue, possible)

		total += g.Red * g.Green * g.Blue
	}
	fmt.Printf("total: %d\n", total)
	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
