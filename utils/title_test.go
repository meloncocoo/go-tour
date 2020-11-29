package utils

import (
	"testing"
)

func TestTitle(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello World", "==================================Hello World================================="},
		{"", "=============================================================================="},
	}
	for _, c := range cases {
		got := Title(c.in)
		if got != c.want {
			t.Errorf("Title(%q) == %q\nwant %q", c.in, got, c.want)
		}
	}
}
