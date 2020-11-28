package study

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"

	"melon.org/utils"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutines() {
	go say("world")
	say("Hello")
}

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}

func channels() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	fmt.Println(time.Now())
	x, y := <-c, <-c
	fmt.Println(time.Now())

	fmt.Println(x, y, x+y)
}

func bufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}

func goFibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func selectFibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func rangeAndClose() {
	c := make(chan int, 50)
	go goFibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func selectForGo() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("\r%d", <-c)
		}
		fmt.Println()
		quit <- 0
	}()
	selectFibonacci(c, quit)
}

func defaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		time.Sleep(200 * time.Millisecond)
		select {
		case <-tick:
			fmt.Printf("\rtick.")
		case <-boom:
			fmt.Printf("\nBOOM!\n")
			return
		default:
			fmt.Printf("\r..")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Walk send all values to channel c
func Walk(t *tree.Tree, c chan int) {
	walkTree(t, c)
	close(c)
}

func walkTree(t *tree.Tree, c chan int) {
	if t != nil {
		walkTree(t.Left, c)
		c <- t.Value
		walkTree(t.Right, c)
	}
}

// Same check two leaf of tree has same value
func Same(t1, t2 *tree.Tree) bool {
	values1 := make(chan int)
	values2 := make(chan int)

	go Walk(t1, values1)
	go Walk(t2, values2)

	for value := range values1 {
		if value != <-values2 {
			return false
		}
	}

	return true
}

func exerciseBinaryTrees() {
	c := make(chan int)
	go Walk(tree.New(1), c)
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

// SafeCounter a safe counter by mutex
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc of SafeCounter method
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
	fmt.Printf("\r=\t%d\t=", c.Value("somekey"))
}

// Value of SafeCounter method
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func mutexCounter() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		time.Sleep(10 * time.Millisecond)
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println()
	fmt.Println("DONE!")
}

// Web Crawler

// Fetcher interface
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// Crawl method
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}

	return
}

func exersiceWebCrawler() {
	var fetcher = fakeFetcher{
		"https://golang.org/": &fakeResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": &fakeResult{
			"Package fmt",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": &fakeResult{
			"Package os",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	}
	Crawl("https://golang.org/", 4, fetcher)
}

// Concurrency main method
func Concurrency() {
	// utils.Title("Go Routines")
	// goroutines()

	// utils.Title("Channels")
	// channels()

	// utils.Title("Buffered Channels")
	// bufferedChannels()

	// utils.Title("Range And Close")
	// rangeAndClose()

	// utils.Title("Select For Go")
	// selectForGo()

	// utils.Title("Default Selection For Go")
	// defaultSelection()

	// utils.Title("Exercise Equivalent Binary Tree")
	// exerciseBinaryTrees()

	// utils.Title("Mutex Counter")
	// mutexCounter()

	utils.Title("Exersice Web Crawler")
	exersiceWebCrawler()
}
