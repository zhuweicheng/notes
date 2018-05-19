// Reverse digits of an integer.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Reverse(321))
}

// Reverse digits of an integer.
// Example1: x = 123, return 321
// Example2: x = -123, return -321
func Reverse(x int) int {
	result := 0
	for ; x != 0; x /= 10 {
		result = result*10 + x%10
	}
	// The input is assumed to be a 32-bit signed integer.
	// Your function should return 0 when the reversed integer overflows.
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
