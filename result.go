package main

// Result is result after doJanken.
type Result int

const (
	// Invalid is error.
	Invalid Result = iota
	// Even is result.
	Even
	// Win is result.
	Win
	// Lose is result.
	Lose
)
