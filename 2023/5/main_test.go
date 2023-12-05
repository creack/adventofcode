package main

import "testing"

func TestPart1(t *testing.T) {
	for _, tc := range []struct {
		filename string
		expect   int
	}{
		{filename: "input.base", expect: 35},
		{filename: "input", expect: 3374647},
		{filename: "input2", expect: 148041808},
	} {
		t.Run(tc.filename, func(t *testing.T) {
			n, err := part1(tc.filename)
			if err != nil {
				t.Fatal(err)
			}
			if n != tc.expect {
				t.Errorf("expect: %d, got: %d\n", tc.expect, n)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	for _, tc := range []struct {
		filename string
		expect   int
	}{
		{filename: "input.base", expect: 46},
		{filename: "input", expect: 6082852},
		{filename: "input2", expect: 148041808},
	} {
		t.Run(tc.filename, func(t *testing.T) {
			n, err := part2(tc.filename)
			if err != nil {
				t.Fatal(err)
			}
			if n != tc.expect {
				t.Errorf("expect: %d, got: %d\n", tc.expect, n)
			}
		})
	}
}
