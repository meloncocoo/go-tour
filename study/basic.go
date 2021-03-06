package study

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
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

// Initialize Value
func printVarInitValue() {
	fmt.Println("===========Init Value==========")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// Type convert
func typeConvert() {
	fmt.Println("===========Type Convert==========")
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// Constants
func printConstants() {
	fmt.Println("===========Constants==========")
	const Pi = 3.14
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstants() {
	fmt.Println("==========Numeric Constants==========")
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// Basic for study go basic
func Basic() {
	fmt.Println("Welcome to the palyground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 32))

	a, b := swap("hello", "word")
	fmt.Println(a, b)

	fmt.Println(split(170))

	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	// Base Types
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Initialize Value
	printVarInitValue()

	// Type convert
	typeConvert()

	// Type inference
	fmt.Println("==========Type Inference==========")
	fmt.Printf("%v: %T\n", 42, 42)
	fmt.Printf("%v: %T\n", 3.142, 3.142)
	fmt.Printf("%v: %T\n", 0.867+0.5i, 0.867+0.5i)

	// Constants
	printConstants()

	// Numeric Constants
	numericConstants()
}
