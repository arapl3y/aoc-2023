package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{""}
	want := 1

	got := part1(input)

	if got != want {
		t.Errorf("part1(%q) = %q, want %q", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []string{""}
	want := 1

	got := part2(input)

	if got != want {
		t.Errorf("part1(%q) = %q, want %q", input, got, want)
	}
}
