// tcp rpc 服务端
package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

// 参数
type Params struct {
	A, B int
}

// 除法和取余
type Division struct {
	Quo, Rem int
}

type Num int

// 求积
// 两个数相乘
func (n *Num) Multiply(args *Params, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// 除法
// 取余
func (n *Num) Divide(args *Params, quo *Division) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	num := new(Num)
	rpc.Register(num)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// 采用TCP协议, 然后需要自己控制连接, 当有客户端连接上来后, 我们需要把这个连接交给rpc来处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
