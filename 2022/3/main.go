package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func test(ctx context.Context) error {
	buf, err := os.ReadFile("input")
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}

	var lines []string

	for _, line := range strings.Split(string(buf), "\n") {
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	if len(lines)%3 != 0 {
		return fmt.Errorf("invalid input, not multiple of 3 (%d)", len(lines))
	}

	const adventPart = 2
	var groupSize int
	if adventPart == 1 {
		groupSize = 1
	} else if adventPart == 2 {
		groupSize = 3
	}

	var groupParts []string
	total := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line)%2 != 0 {
			return fmt.Errorf("invalid odd line: %q (%d)", line, len(line))
		}

		// Split the line in half as we have 2 items per line.
		groupParts = append(groupParts, line[0:len(line)/2], line[len(line)/2:])
		if len(groupParts)%(groupSize*2) != 0 {
			continue
		}
		// Sanity check.
		if len(groupParts) == 0 {
			return fmt.Errorf("invalid group: 0 size")
		}
		// Reset the groupParts variable.
		gp := groupParts
		groupParts = nil

		// Go over each part and keep track of the unique items.
		var items []map[rune]int
		for _, elem := range gp {
			m := map[rune]int{}
			for _, c := range elem {
				m[c] = 1
			}
			items = append(items, m)
		}

		// Go over each unique items and count how many of each we have.
		seen := map[rune]int{}
		for _, item := range items {
			for elem := range item {
				seen[elem]++
			}
		}

		// Go over our seen items and check the size to see if it is a common one.
		commonItems := map[rune]int{}
		for elem, n := range seen {
			// For part 1, we have 2 items per group.
			if adventPart == 1 && n == groupSize*2 {
				commonItems[elem]++
			}
			// For part 2, we only consider the whole.
			if adventPart == 2 && n == groupSize {
				commonItems[elem]++
			}
		}

		// Sanity check.
		if len(commonItems) != 1 {
			return fmt.Errorf("unexpected common items %q (%d, %q)", line, len(commonItems), commonItems)
		}

		// Extract the commonItem from the map.
		var commonItem rune
		for k := range commonItems {
			commonItem = k
		}

		// Convert item type to numerical value.
		itemType := byte(commonItem)
		var itemTypeValue byte
		switch {
		case itemType >= 'a' && itemType <= 'z':
			itemTypeValue = itemType - 'a' + 1
		case itemType >= 'A' && itemType <= 'Z':
			itemTypeValue = itemType - 'A' + 27
		}

		// Accumulate the total.
		total += int(itemTypeValue)
	}
	fmt.Printf("total: %d\n", total)

	return nil
}

func main() {
	if err := test(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
