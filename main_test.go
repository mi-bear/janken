package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoJanken(t *testing.T) {

	assert := assert.New(t)

	assert.Equal(Rock, Rock, "引分")
	assert.Equal(Rock, Scissors, "勝利")
	assert.Equal(Rock, Paper, "敗北")

	assert.Equal(Scissors, Scissors, "引分")
	assert.Equal(Scissors, Paper, "勝利")
	assert.Equal(Scissors, Rock, "敗北")

	assert.Equal(Paper, Paper, "引分")
	assert.Equal(Paper, Rock, "勝利")
	assert.Equal(Paper, Scissors, "敗北")
}
