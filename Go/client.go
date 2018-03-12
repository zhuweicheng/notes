// 首先程序将用户的输入作为参数service传入net.ResolveTCPAddr获取一个tcpAddr,
// 然后把tcpAddr传入DialTCP后创建了一个TCP连接conn, 通过conn来发送请求信息,
// 最后通过ioutil.ReadAll从conn中读取全部的文本, 也就是服务端响应反馈的信息
package main

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

func main() {
	// 获取提供服务的服务器ip地址
	service := os.Args[1]
	// func ResolveTCPAddr(net, addr string) (*TCPAddr, error)
	// ResolveTCPAddr获取一个TCPAddr, 一个TCPAddr类型, 他表示一个TCP的地址信息
	// ResolveTCPAddr将addr作为TCP地址解析并返回
	// 参数addr格式为"host:port"或"[ipv6-host%zone]:port", 解析得到网络名和端口名
	// net必须是"tcp", "tcp4"或"tcp6"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	// 通过net包中的DialTCP函数来建立一个TCP连接, 并返回一个TCPConn类型的对象
	// 当连接建立时服务器端也创建一个同类型的对象, 此时客户端和服务器端通过各自拥
	// 有的TCPConn对象来进行数据交换
	//
	// 一般而言, 客户端通过TCPConn对象将请求信息, 发送到服务器端, 读取服务器端响应的信息
	// 服务器端读取并解析来自客户端的请求并返回应答信息, 这个连接只有当任一端关闭了连接之后才失效,
	// 不然这连接可以一直在使用
	//
	// func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error)
	// DialTCP在网络协议net上连接本地地址laddr和远端地址raddr
	// net必须是"tcp", "tcp4", "tcp6"
	// 如果laddr不是nil, 将使用它作为本地地址, 否则自动选择一个本地地址
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	// 向连接的服务器端写入信息
	// net包中有一个类型TCPConn, 这个类型可以用来作为客户端和服务器端交互的通道, 他有两个主要的函数:
	// func (c *TCPConn) Write(b []byte) (n int, err os.Error)
	// func (c *TCPConn) Read(b []byte) (n int, err os.Error)
	// TCPConn可以用在客户端和服务器端来读写数据
	//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	_, err = conn.Write([]byte("timestamp"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write([]byte("a~~~"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write([]byte("b~~~~~"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write([]byte("c~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	// 当客户端停止写入时, 需要告诉服务端, 信息发送终止, 服务端就返回全部的数据
	if err = conn.CloseWrite(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	// 读取服务器端响应的全部内容
	// ReadAll从r读取数据直到EOF或遇到error, 返回读取的数据和遇到的错误
	// 成功的调用返回的err为nil而非EOF, 因为本函数定义为读取r直到EOF, 它不会将读取返回的EOF视为应报告的错误
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	// 输出获取到的内容
	fmt.Println(string(result))
	os.Exit(0)
}
