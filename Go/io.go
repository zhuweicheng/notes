// io 包的使用
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Start...")

	// 1. 从文件读取内容
	file, err := os.OpenFile("test", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for {
		f, err := ReadFrom(file, 1024)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(f))
	}

	// 2. 从字符串中读取
	d, err := ReadFrom(strings.NewReader("This is a apple."), 128)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))

	// 3. 从标准输入读取内容
	fmt.Println("3. Read From stdin, please input q to quit")
	for {
		d, err := ReadFrom(os.Stdin, 10)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(d))
		// 要捕捉到输入的字符, 需要过滤所有的非字符值, 默认会多输出一个换行
		if strings.TrimSpace(string(d)) == "q" {
			fmt.Println("Hi, Quit.")
			break
		}
	}

	fmt.Println("End")
}

// 函数将 io.Reader 作为参数, 也就是说, ReadFrom 可以从任意的地方读取数据, 只要来源实现了 io.Reader 接口
// 比如, 我们可以从标准输入, 文件, 字符串等读取数据
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
