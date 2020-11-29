package study

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"golang.org/x/tour/pic"

	"github.com/meloncocoo/go-tour/utils"
	"golang.org/x/tour/reader"
)

// Abser is a interface
type Abser interface {
	Abs() float64
}

// MyFloat is my float64
type MyFloat float64

// Vertex is vertex type
type Vertex struct {
	X, Y float64
}

// Abs for MyFloat implements
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Abs for Vertex implements
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func goInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	a = v

	fmt.Println(a.Abs())
}

// I interface define
type I interface {
	M()
}

// T interface define
type T struct {
	S string
}

// M for type *T
func (t *T) M() {
	fmt.Println(t.S)
}

// F type define
type F float64

// M for type F
func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceValues() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic error
	// fmt.Println(f)
}

func doTypeSwitches(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func typeSwitches() {
	doTypeSwitches(21)
	doTypeSwitches("hello")
	doTypeSwitches(true)
}

// Person struct define
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

// IPAddr type define
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func ipFormat() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// MyError struct define
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() (err error) {
	fmt.Println("Call run method.")
	err = &MyError{
		time.Now(),
		"it didn't work",
	}
	return
}

func errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// ErrNegativeSqrt struct define
type ErrNegativeSqrt struct {
	x float64
}

func (err *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", err.x)
}

// Sqrt method
func Sqrt(x float64) (value float64, err error) {
	if x < 0 {
		err = &ErrNegativeSqrt{
			-2,
		}
	}

	value = math.Sqrt(x)
	return
}

func exerciseErrors() {
	var val, err = Sqrt(2)
	fmt.Println(val, err)
	val, err = Sqrt(-2)
	fmt.Println(val, err)
}

func readerDemo() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// MyReader struct define
type MyReader struct{}

func (reader MyReader) Read(b []byte) (n int, err error) {
	b = b[:cap(b)]
	for i := range b {
		b[i] = 'A'
	}
	return cap(b), nil
}

func exerciseReader() {
	reader.Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

// Rot13 method
func Rot13(b byte) byte {
	if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
		b += 13
	} else if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
		b -= 13
	}

	return b
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	for i := 0; i < len(p); i++ {
		p[i] = Rot13(p[i])
	}

	return n, err
}

func exerciseRotReader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func imageDemo() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 10).RGBA())
}

// Image struct define
type Image struct {
	Width, Height int
	colors        uint8
}

// ColorModel implement
func (m *Image) ColorModel() color.Model {
	return color.RGBA64Model
}

// Bounds method
func (m *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Width, m.Height)
}

// At method
func (m *Image) At(x, y int) color.Color {
	return color.RGBA{m.colors + uint8(x), m.colors + uint8(y), 0, 125}
}

func exerciseImage() {
	m := Image{200, 100, 20}
	pic.ShowImage(&m)
}

// Methods for study usage of go methods
func Methods() {
	utils.Title("Show go interface usage")
	goInterface()

	utils.Title("Interface values")
	interfaceValues()

	utils.Title("Type Assertions")
	typeAssertions()

	utils.Title("Type Switches")
	typeSwitches()

	utils.Title("Stringer")
	stringer()

	utils.Title("Stringer Exercise For IP address")
	ipFormat()

	utils.Title("Errors")
	errors()

	utils.Title("Exercise Errors")
	exerciseErrors()

	utils.Title("Reader Demo")
	readerDemo()

	utils.Title("Exercise Reader")
	exerciseReader()

	utils.Title("Exercise Rot Reader")
	exerciseRotReader()

	utils.Title("Images Demo")
	imageDemo()

	utils.Title("Exercise Image")
	exerciseImage()
}
