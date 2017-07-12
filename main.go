package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	var me, com Hand = Rock, Paper

	result, err := doJanken(me, com)
	if err != nil {
		fmt.Println(errors.Wrap(err, "error").Error())
		return
	}
	fmt.Printf("I'm: %v, computer's: %v, Result: I %v...\n", me, com, result)
}

func doJanken(me, com Hand) (Result, error) {
	if err := me.validate(); err != nil {
		return Invalid, errors.Wrap(err, "me")
	}

	if err := com.validate(); err != nil {
		return Invalid, errors.Wrap(err, "com")
	}

	return me.janken(com), nil
}
