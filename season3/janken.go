package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var user, computer Hand

func init() {
	user = Rock
	computer = Paper

	fmt.Printf("Rock=%v, Scissors=%v, Paper=%v\n", Rock, Scissors, Paper)
	fmt.Printf("Invalid=%v, Even=%v, Lose=%v, Win=%v\n", Invalid, Even, Lose, Win)
}

func main() {
	result, err := janken(user, computer)
	if err != nil {
		fmt.Println(errors.Wrap(err, "error").Error())
		return
	}
	fmt.Printf("user's: %v, computer's: %v, Result: user %v...\n", user, computer, result)
}

func janken(user, computer Hand) (Result, error) {
	if err := user.validate(); err != nil {
		return Invalid, errors.Wrap(err, "user")
	}
	if err := computer.validate(); err != nil {
		return Invalid, errors.Wrap(err, "computer")
	}

	r := (user - computer + 3) % 3
	return Result(r), nil
}
