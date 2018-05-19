// http rpc 客户端
// go run httpclient.go 127.0.0.1
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// 两个整数参数
type Args struct {
	A, B int
}

// 两个整数值 Quo除法 Rem取余
type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	// 连接提供服务的地址
	serverAddress := os.Args[1]

	// 连接服务器
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing error:", err)
	}
	// 提供参数
	args := Args{20, 6}
	var reply int
	
	fmt.Println("Call...")
	// 同步调用rpc服务
	// Arith.Multiply 服务器上提供的对应执行方法
	// args 提供符合格式需要的参数
	// reply 接收计算结果
	// 返回错误信息
	// Call调用指定的方法, 等待调用返回, 将结果写入reply, 然后返回执行的错误状态
	// serviceMethod 方法为服务端注册的已有方法, 然后是参数传递和结果返回值指针
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Call Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Call Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

	fmt.Println("Go...")
	// Go异步的调用函数
	// 本方法Call结构体类型指针的返回值代表该次远程调用
	// 通道类型的参数done会在本次调用完成时发出信号(通过返回本次Go方法的返回值)
	// 如果done为nil, Go会申请一个新的通道(写入返回值的Done字段)
	// 如果done非nil, done必须有缓冲, 否则Go方法会故意崩溃
	// new 本身返回指针类型
	args = Args{25, 4}
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	// divCall 与 replyCall 是相同的
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		fmt.Println(replyCall.Error.Error())
	}

	// 保存了所有环境值, 但是都属于interface 无法读取
	// {17 8}
	// &{2 1}
	// <nil>
	// Arith.Divide
	//
	fmt.Printf("Go Arith: %d/%d=%d remainder %d\n", args.A, args.B, quotient.Quo, quotient.Rem)
	fmt.Println(divCall.Args)
	fmt.Println(divCall.Reply)
	fmt.Println(divCall.Error)
	fmt.Println(divCall.ServiceMethod)
	
	// 返回值依然保存在指定的接收结构体中
	fmt.Println(quotient)
	
	fmt.Println(quotient.Quo)
	fmt.Println(quotient.Rem)

	// 附带的返回值, 都是接口类型
	fmt.Println(replyCall.Args)
	fmt.Println(replyCall.Reply)
	fmt.Println(replyCall.Error)
	fmt.Println(replyCall.ServiceMethod)
}
