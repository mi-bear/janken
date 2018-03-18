package main

// Result is result after Janken.
type Result int

const (
	// Invalid is error.
	Invalid Result = iota
	// Even is result after Janken.
	Even
	// Win is result after Janken.
	Win
	// Lose is result after Janken.
	Lose
)
