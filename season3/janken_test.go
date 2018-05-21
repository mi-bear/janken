package main

import "testing"

func Test_janken(t *testing.T) {
	tests := []struct {
		user, computer Hand
		result         Result
		err            bool
	}{

		{Rock, Rock, Even, false},
		{Rock, Paper, Lose, false},
		{Rock, Scissors, Win, false},
		{Scissors, Rock, Lose, false},
		{Scissors, Paper, Win, false},
		{Scissors, Scissors, Even, false},
		{Paper, Rock, Win, false},
		{Paper, Paper, Even, false},
		{Paper, Scissors, Lose, false},

		{4, Rock, Invalid, true},
		{Paper, 0, Invalid, true},
	}

	for _, test := range tests {
		result, err := janken(test.user, test.computer)
		e := err != nil

		if e != test.err {
			t.Errorf("should have error %+v", test)
			continue
		}

		if e {
			continue
		}

		if result != test.result {
			t.Errorf("janken(%v, %v) = %v, want %v", test.user, test.computer, result, test.result)
		}
	}
}
