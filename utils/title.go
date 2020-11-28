package utils

import (
	"fmt"
	"strings"
)

// Title print title
func Title(s string) {
	width := 80
	if len(s) > width-2 {
		width = len(s) + len(s)%2 + 10
	}
	padding := (width-2)/2 - len(s)/2
	fmt.Printf("%s %s %s\n", strings.Repeat("=", padding), s, strings.Repeat("=", padding-len(s)%2))
}

// Chapter print chapter
func Chapter(s string) {
	width := 80
	if len(s) > width-2 {
		width = len(s) + len(s)%2 + 10
	}
	left := (width-2)/2 + len(s)/2
	right := (width - 2) - left
	fmt.Printf("%s\n", strings.Repeat("#", width))
	fmt.Printf(fmt.Sprintf("#%%%ds%%%ds#\n", left, right), s, "")
	fmt.Printf("%s\n", strings.Repeat("#", width))
}
