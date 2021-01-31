package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for ; sum < 200; {
		sum += sum
	}
	fmt.Println(sum)

    for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

    fmt.Println(opera(2), opera(-4))

    fmt.Println(
		"multiplica1: ",multiplica(3, 2, 10),
		"multiplica2: ",multiplica(3, 3, 20),
	)

    fmt.Println("sqrt: ",Sqrt(2))


}

func opera(x float64) string {
	if x < 0 {
		return opera(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func multiplica(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}


func Sqrt(x float64) float64 {
}
