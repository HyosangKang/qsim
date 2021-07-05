package qsim

import "strconv"

func Binary(n, d int) string {
	b := ""
	for i := 0; i < d; i++ {
		b += strconv.Itoa(n - (n/2)*2)
		n /= 2
	}
	bb := ""
	for i := 0; i < len(b); i++ {
		bb += string(b[len(b)-i-1])
	}
	return bb
}
