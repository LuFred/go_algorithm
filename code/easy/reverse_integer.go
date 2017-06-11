package easy

import "math"

/**
	Reverse Integer
question:
	Reverse digits of an integer.
	Example1: x = 123, return 321
	Example2: x = -123, return -321

	Note:
The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.
**/
func Reverse(x int) int {
	r := 0
	i := 0
	n := false
	if x < 0 {
		x = x * -1
		n = true
	}
	for ; x > 0; x = x / 10 {
		i = x % 10
		r = r*10 + i
	}
	if r > math.MaxInt32 {
		return 0
	}
	if n {
		r = r * -1
	}
	return r
}
