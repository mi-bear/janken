package main

import "github.com/pkg/errors"

// Hand is sign on Janken.
type Hand int

const (
	_ Hand = iota
	// Rock is sign on Janken.
	Rock
	// Scissors is sign on Janken.
	Scissors
	// Paper is sign on Janken.
	Paper
)

func (h Hand) validate() error {
	switch h {
	case Rock, Scissors, Paper:
		return nil
	}
	return errors.Errorf("%s is invalid Hand.", h)
}

func (h Hand) janken(com Hand) Result {
	result := (h - com + 3) % 3
	switch result {
	case Even, Win, Lose:
		return result
	}
	return Invalid
}
