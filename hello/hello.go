package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j = 1, 2

func main() {
	// golang := true
	// var c, python, java = true, false, "no"
	var x, y = swap("hello", "world")
	fmt.Println(x, y)
	// fmt.Println(i, golang, c, python, java)
}
