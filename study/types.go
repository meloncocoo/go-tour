package study

import (
	"fmt"
	"math"
	"strings"

	"github.com/meloncocoo/go-tour/utils"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func pointer() {
	i, j := 42, 2741

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func structs() {
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	fmt.Println(v)

	utils.Title("Struct Fields")
	v.X = 4
	v.Y = 100
	fmt.Println(v)

	utils.Title("Struct Pointer")
	p := &v
	p.X = 1e9
	fmt.Println(v)

	utils.Title("Struct Literals")
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
	)
	p = &Vertex{3, 4}
	fmt.Println(v1, p, v2, v3)
}

func array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	utils.Title("Slices Pointer")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:4]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceBounds() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func sliceLenCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("s", s)

	s = s[:0]
	printSlice("s", s)

	s = s[1:4]
	printSlice("s", s)

	s = s[2:]
	printSlice("s", s)
}

func nilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}

func makeSlices() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func slicesOfSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendSlices() {
	var s []int
	printSlice("s", s)

	s = append(s, 0)
	printSlice("s", s)

	s = append(s, 1)
	printSlice("s", s)

	s = append(s, 2, 3, 4)
	printSlice("s", s)
}

func forRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func forRangeContinued() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// Pic exercise code
func Pic(dx, dy int) [][]uint8 {
	pixel := make([][]uint8, dy)
	data := make([]uint8, dx)

	for i := range pixel {
		for j := range data {
			// data[j] = uint8((i + j) / 2)
			// data[j] = uint8(math.Pow(float64(i), float64(j)))
			// data[j] = uint8(float64(i) * math.Log(float64(j)))
			data[j] = uint8(i % (j + 1))
		}
		pixel[i] = data
	}

	return pixel
}

func exerciseSlices() {
	pic.Show(Pic)
}

func maps() {
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	m = map[string]Vertex{
		// "Bell Labs": {40.68433, -74.39967}
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		// "Google": {37.42202, -122.08408}
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// WordCount for test
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	fields := strings.Fields(s)

	for _, v := range fields {
		elem, ok := m[v]
		if ok {
			m[v] = elem + 1
		} else {
			m[v] = 1
		}
	}

	return m
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funcValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("Sqrt(5*5 + 12*12) =", hypot(5, 12))

	fmt.Println("Sqrt(3*3 + 4*4) =", compute(hypot))
	fmt.Println("3^4 =", compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func funcClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	current, next := 0, 1
	return func() (result int) {
		result = current
		current, next = next, current+next
		return
	}
}

func exerciseFibonacci() {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
}

// Types for study more types
func Types() {
	utils.Title("Pointer")
	pointer()

	utils.Title("Stucts")
	structs()

	utils.Title("Array")
	array()

	utils.Title("Slices")
	slices()

	utils.Title("Slice Literals")
	sliceLiterals()

	utils.Title("Slice Bounds")
	sliceBounds()

	utils.Title("Slice Length Capacity")
	sliceLenCapacity()

	utils.Title("Nil Slices")
	nilSlices()

	utils.Title("Make slices")
	makeSlices()

	utils.Title("Slices Of Slices")
	slicesOfSlices()

	utils.Title("Append Slices")
	appendSlices()

	utils.Title("For Range")
	forRange()

	utils.Title("For Range Continued.")
	forRangeContinued()

	utils.Title("Exercise Slices")
	exerciseSlices()

	utils.Title("Maps")
	maps()

	utils.Title("Mutating Maps")
	mutatingMaps()

	utils.Title("Exercise Maps")
	wc.Test(WordCount)

	utils.Title("Function Values")
	funcValues()

	utils.Title("Functions Closures")
	funcClosures()

	utils.Title("Exercise for Fibonacci")
	exerciseFibonacci()
}
