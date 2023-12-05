package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(fileName string) (*Data, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	d := &Data{}
	for i, block := range strings.Split(string(buf), "\n\n") {
		if block == "" {
			continue
		}
		// Handle the first block, the inital seeds.
		if i == 0 {
			// Part 1. Each number is a seed.
			for _, elem := range strings.Split(strings.TrimPrefix(block, "seeds: "), " ") {
				n, err := strconv.Atoi(elem)
				if err != nil {
					return nil, fmt.Errorf("atoi %q: %w", elem, err)
				}
				d.Seeds = append(d.Seeds, n)
			}
			// Part 2. Each pair is a range.
			for i := 0; i < len(d.Seeds); i += 2 {
				d.SeedRanges = append(d.SeedRanges, Range{src: d.Seeds[i], len: d.Seeds[i+1]})
			}
			continue
		}

		// Handle the rest.
		var subEntries MapperList
		for j, line := range strings.Split(block, "\n") {
			// Skip the first line.
			if j == 0 || line == "" {
				continue
			}

			var entry Mapper
			if _, err := fmt.Sscanf(line, "%d %d %d", &entry.dst, &entry.src, &entry.len); err != nil {
				return nil, fmt.Errorf("invalid entry line %q: %w", line, err)
			}
			subEntries = append(subEntries, entry)

		}
		d.Entries = append(d.Entries, subEntries)
	}
	return d, nil
}
