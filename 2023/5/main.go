package main

import (
	"fmt"
	"math"
)

type Mapper struct {
	Range
	dst int
}

func (m Mapper) processMapping(in Range) Range {
	return Range{
		src: int(int(in.src) - int(m.src) + int(m.dst)),
		len: int(in.len),
	}
}

type MapperList []Mapper

// Go over each of the mapper block lines, check if it intersects with
// the input, if it does, process the mapping.
func (rl MapperList) processMappings(entries []Range) []Range {
	var out []Range

	for _, elem := range rl {
		var retryEntries []Range
		for _, entry := range entries {
			inter, unmatched := entry.intersect(elem.Range)
			if inter.len > 0 {
				out = append(out, elem.processMapping(inter))
			}
			retryEntries = append(retryEntries, unmatched...)
		}
		entries = retryEntries
	}

	return append(out, entries...)
}

type Data struct {
	Seeds []int // First line numbers.

	SeedRanges []Range
	Entries    []MapperList
}

// Go over each blocks and process their mappings.
func (d *Data) processMappings() []Range {
	out := d.SeedRanges
	for _, elem := range d.Entries {
		out = elem.processMappings(out)
	}
	return out
}

func part2(fileName string) (int, error) {
	d, err := parse(fileName)
	if err != nil {
		return -1, fmt.Errorf("parse: %w", err)
	}

	minLoc := math.MaxInt
	for _, elem := range d.processMappings() {
		minLoc = min(minLoc, elem.src)
	}
	return minLoc, nil
}

func part1(fileName string) (int, error) {
	d, err := parse(fileName)
	if err != nil {
		return -1, fmt.Errorf("parse: %w", err)
	}

	for _, s := range d.Seeds {
		d.SeedRanges = append(d.SeedRanges, Range{src: s, len: 1})
	}

	minLoc := math.MaxInt
	for _, elem := range d.processMappings() {
		minLoc = min(minLoc, elem.src)
	}
	return minLoc, nil
}

func main() {
	if _, err := part2("input.base"); err != nil {
		println("Fail:", err.Error())
		return
	}
}
