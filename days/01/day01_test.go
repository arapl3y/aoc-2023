package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{"1testinput1", "2hello2", "one4hhy3", "two8four4"}
	want := 70

	got := part1(input)

	if got != want {
		t.Errorf("part1(%q) = %q, want %q", input, got, want)
	}
}
