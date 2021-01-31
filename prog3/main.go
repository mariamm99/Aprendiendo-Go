package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var c, python, java bool
var i, j int = 1, 2
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	k := 3
	fmt.Println(i, j, c, python, java, k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var numero1 int
	var numero2 float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", numero1, numero2, b, s)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	v := "hola" // al cambiar el valor, cambio string/int...
	fmt.Printf("v is of type %T\n", v)
	const World = "María"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	 return x*10 + 1 
	}
func needFloat(x float64) float64 {
	return x * 0.1
}
