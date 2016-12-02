package stringutil

import "unicode/utf8"

func Reverse(s string) string {
	r := []rune(s)
	for i, j :=  0, len(r) -1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseUtf8(s string) string {
	totalLength := len(s)
	buffer := make([]byte, totalLength)
	for i := 0; i < totalLength; {
		r, size := utf8.DecodeRuneInString(s[i:])
		i += size
		utf8.EncodeRune(buffer[totalLength-i:], r)
	}
	return string(buffer)
}
