package main

// Result is result after Janken.
type Result int

const (
	// Invalid is error.
	Invalid Result = iota - 1
	// Even is result after Janken.
	Even
	// Lose is result after Janken.
	Lose
	// Win is result after Janken.
	Win
)
