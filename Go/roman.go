package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCVI"))
}

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	prev := 0
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		cur := roman[string(s[i])]
		if cur < prev {
			result -= cur
		} else {
			result += cur
		}
		prev = cur
	}
	return result
}
