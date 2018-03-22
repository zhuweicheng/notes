// 外部归并排序
package main

import (
	"bufio"
	"fmt"
	"notes/go/util"
	"os"
)

// 文件名
const FILENAME = "small.in"

// 写800M数据
const N = 100000000

func main() {
	//a2()
	a3()
}

// 指定输入的两个切片数组进行排序
func a1() {
	p := util.Merge(util.InMemSort(util.ArrSort(1, 4, 6, 2, 18, 3, 26, 5)), util.InMemSort(util.ArrSort(28, 2, 10, 3, 82, 5, 42, 1)))
	for m := range p {
		fmt.Println(m)
	}
}

// 生成文件
func a2() {
	file, err := os.Create(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := util.RandomSource(N)
	// 使用缓冲方式写数据，使用缓冲比不适用缓冲快了好几倍
	w := bufio.NewWriter(file)
	util.WriteSink(w, p)
	// 最后数据可能不能写缓冲, 需要强制刷新缓冲
	w.Flush()

	fd, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	// 使用有缓冲方式读数据
	r := util.ReadSource(bufio.NewReader(fd), -1)
	n := 0
	for m := range r {
		fmt.Println(m)
		n++
		// 只输出前100个二进制数字
		if n >= 100 {
			break
		}
	}
}

// 外部归并排序
func a3() {
	p := create(FILENAME, 800000000, 4)
	write(p, "large.out")
	echo("large.out")
}

// 可读入资源切片
func create(fileName string, fileSize, count int) <-chan int {
	chunkSize := fileSize / count

	var result []<-chan int
	for i := 0; i < count; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		// 设置文件偏移量, 并从偏移量开始0处读取文件内容
		file.Seek(int64(i*chunkSize), 0)

		p := util.ReadSource(bufio.NewReader(file), chunkSize)

		result = append(result, util.InMemSort(p))
	}

	return util.MergeN(result...)
}

// 将排序好的数据写入文件中
func write(in <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	util.WriteSink(w, in)
}

// 输出文件内容
func echo(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := util.ReadSource(file, -1)
	n := 0
	for m := range p {
		fmt.Println(m)
		n++
		if n >= 100 {
			break
		}
	}
}
