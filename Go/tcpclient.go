// tcp rpc 客户端
// go run tcpclient.go 127.0.0.1:1234
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Arg struct {
	A, B int
}

type Quot struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	// http 和 tcp 唯一的区别在这里, http 使用DialHTTP来连接, 而tcp使用Dial来连接
	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Arg{17, 8}
	var reply int
	err = client.Call("Num.Multiply", args, &reply)
	if err != nil {
		log.Fatal("num error:", err)
	}
	fmt.Printf("Num: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quot
	err = client.Call("Num.Divide", args, &quot)
	if err != nil {
		log.Fatal("num error:", err)
	}
	fmt.Printf("Num: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
