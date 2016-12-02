package stringutil

import "testing"

var cases = []struct {
	in, want string
}{
	{"Hello, world", "dlrow ,olleH"},
	{"Hello, 세계", "계세 ,olleH"},
	{"한", "한"},
	{"", ""},
}

func TestReverse(t *testing.T) {
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, instead of %q", c.in, got, c.want)
		}
	}
}

func TestReverseUtf8(t *testing.T) {
	for _, c := range cases {
		got := ReverseUtf8(c.in)
		if got != c.want {
			t.Errorf("ReverseUtf8(%q) == %q, instead of %q", c.in, got, c.want)
		}
	}
}
