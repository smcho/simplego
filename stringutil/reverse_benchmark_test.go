package stringutil

// go test -v -test.bench=".*" -test.timeout="60m"

import "testing"

var s = "Hello, 세계"

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Reverse(s)
	}
}

func BenchmarkReverseUtf8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReverseUtf8(s)
	}
}
