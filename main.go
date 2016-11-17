package main

import (
	"fmt"
	"os"
)

const (
	_        = iota
	Rock     = 1 // グー
	Scissors = 2 // チョキ
	Paper    = 3 // パー
)
const (
	_    = iota
	Even = 0
	Win  = 1
	Lose = 2
)

func main() {

	var a int     // 自分
	var b int     // 相手
	var res int   // 勝敗
	var err error // エラー

	a = Rock  // グー
	b = Paper // チョキ

	if res, err = doJanken(a, b); err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

	switch res {
	case Even:
		fmt.Println("引分")
	case Win:
		fmt.Println("勝利")
	case Lose:
		fmt.Println("敗北")
	default:
		fmt.Println("???")
	}
}

func doJanken(a, b int) (int, error) {

	if a > 3 || b > 3 {

		message := "parameter error"
		return Even, fmt.Errorf("error: %s", message)
	}

	res := (b - a + 3) % 3

	return res, nil
}
