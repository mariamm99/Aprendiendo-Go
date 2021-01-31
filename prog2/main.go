package main

import "fmt"

func main() {
	fmt.Println(suma(42, 13))
	a, b, c := swap("hello", "world", "!!!!")
	fmt.Println(a, b, c)
	fmt.Println(split(10))
}

func suma(x int, y int) int {
	return x + y
}

func swap(x, y, z string) (string, string, string) {
	return y, x, z
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return y, x
}
