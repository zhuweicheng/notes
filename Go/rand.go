package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s, err := Salt(6, 2, 2, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}

// Number 随机生成size个数字
func Number(size int) []byte {
	if size <= 0 || size > 10 {
		size = 10
	}
	warehouse := []int{48, 57}
	result := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + rand.Intn(9))
	}
	return result
}

// Lower 随机生成size个小写字母
func Lower(size int) []byte {
	if size <= 0 || size > 26 {
		size = 26
	}
	warehouse := []int{97, 122}
	result := make([]byte, 26)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + rand.Intn(26))
	}
	return result
}

// Lower 随机生成size个小写字母
func Upper(size int) []byte {
	if size <= 0 || size > 26 {
		size = 26
	}
	warehouse := []int{65, 90}
	result := make([]byte, 26)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + rand.Intn(26))
	}
	return result
}

// Salt 生成一个盐值
func Salt(size int, number int, lower int, upper int) (string, error) {
	// 参数校验
	length := number + lower + upper
	if size > length || length <= 0 {
		return "", errors.New("非法的长度")
	}
	switch {
	case number < 0, number > 10:
		return "", errors.New("允许的数字范围0-10")
	case lower < 0, upper < 0, lower > 26, upper > 26:
		return "", errors.New("允许的字母范围0-26")
	}
	
	// 按需要生成字符串
	var result string
	if lower >= 0 {
		lowers := string(Lower(lower))
		result += lowers
	}
	if number >= 0 {
		numbers := string(Number(number))
		result += numbers
	}
	if upper >= 0 {
		uppers := string(Upper(upper))
		result += uppers
	}
	
	return result, nil
}
