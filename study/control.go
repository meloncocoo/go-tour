package study

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func title(title string) {
	fmt.Println("==========", title)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func sqrtEx(x float64) float64 {
	z := x / 2
	for count := 1; count < 10; count++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// Control for study statement
//
func Control() {
	fmt.Println("==========For")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println("==========For is go's while")
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println("==========Forever")
	// for {}

	fmt.Println("==========If")
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println("==========If with a short statement")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	title("If and else")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	title("Exercise loops and functions")
	fmt.Println(sqrtEx(2))

	title("Switch")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	title("Switch evaluation order")
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too for away")
	}

	title("Switch with no condition")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	title("Defer")
	defer fmt.Println("world")
	fmt.Println("hello")

	title("Defer multi")
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
