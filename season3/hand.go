package main

import "github.com/pkg/errors"

// Hand is sign on Janken.
type Hand int

const (
	// Rock is sign on Janken.
	Rock Hand = iota + 1
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
