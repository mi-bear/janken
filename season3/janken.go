package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	var user, computer Hand = Rock, Paper

	result, err := janken(user, computer)
	if err != nil {
		fmt.Println(errors.Wrap(err, "error").Error())
		return
	}
	fmt.Printf("I'm: %v, computer's: %v, Result: I %v...\n", user, computer, result)
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
