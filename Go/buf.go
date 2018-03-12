// bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，
// 且同时还提供了缓冲和一些文本I/O的帮助函数的对象。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 往标准输出写, 一个能写的对象, 文件等
	// NewWriter创建一个具有默认大小缓冲, 写入w的*Writer
	buf1 := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(buf1, "Hi")
	fmt.Fprintln(buf1, "Good")
	// 数据都写入后, 调用者有义务调用Flush方法以保证所有的数据都交给了下层的io.Writer
	buf1.Flush()

	// 写入[]byte类型
	buf1.Write([]byte("My Name Is HiHi\n"))
	buf1.Flush()

	// 写入string类型
	buf1.WriteString("Nice to meet you\n")
	buf1.Flush()

	// Scanner类型提供了方便的读取数据的接口，如从换行符分隔的文本里读取每一行。
	//
	// 成功调用的Scan方法会逐步提供文件的token，跳过token之间的字节。token由SplitFunc类型的分割函数指定;
	// 默认的分割函数会将输入分割为多个行，并去掉行尾的换行标志。本包预定义的分割函数可以将文件分割为行、
	// 字节、unicode码值、空白分隔的word。调用者可以定制自己的分割函数。
	//
	// 扫描会在抵达输入流结尾、遇到的第一个I/O错误、token过大不能保存进缓冲时，不可恢复的停止。当扫描停止后，
	// 当前读取位置可能会远在最后一个获得的token后面。需要更多对错误管理的控制或token很大，或必须从reader连
	// 续扫描的程序，应使用bufio.Reader代替。
	//
	// 自定义分隔函数
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 本函数会将空白（参见unicode.IsSpace）分隔的片段（去掉前后空白后）作为一个token返回
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			// 最后一个数32位不够, 64位可以
			_, err = strconv.ParseInt(string(token), 10, 64)
		}
		return
	}
	// 调用自定义的分隔函数
	// Split设置该Scanner的分割函数。本方法必须在Scan之前调用。
	scanner.Split(split)
	// 成功转化的部分
	// Scan方法获取当前位置的token（该token可以通过Bytes或Text方法获得），并让Scanner的扫描位置移动到下一个token。
	// 当扫描因为抵达输入流结尾或者遇到错误而停止时，本方法会返回false。
	// 在Scan方法返回false后，Err方法将返回扫描时遇到的任何错误；除非是io.EOF，此时Err会返回nil。
	for scanner.Scan() {
		// Bytes方法返回最近一次Scan调用生成的token。底层数组指向的数据可能会被下一次Scan的调用重写。
		// Text方法返回最近一次Scan调用生成的token，会申请创建一个字符串保存token并返回该字符串。
		fmt.Printf("%s\n", scanner.Text())
	}
	// 存储的错误信息
	// Err返回Scanner遇到的第一个非EOF的错误。
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
	}

	// 使用默认按行分隔, 从标准输入读取, 测试时打开, 否则会阻塞
	//scanner = bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	//if err := scanner.Err(); err != nil {
	//	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	//}

	// 按单词分隔
	const input2 = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner = bufio.NewScanner(strings.NewReader(input2))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}
