package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//printLn
	fmt.Println("Hello, María Moreno muñoz")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(20))
	//printf
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("numero pi", math.Pi)
}
