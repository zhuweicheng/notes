// http rpc 服务端
//
// rpc包提供了通过网络或其他I/O连接对一个对象的导出方法的访问
// 服务端注册一个对象, 使它作为一个服务被暴露, 服务的名字是该对象的类型名
// 注册之后, 对象的导出方法就可以被远程访问
// 服务端可以注册多个不同类型的对象(服务), 但注册具有相同类型的多个对象是错误的
//
// 只有满足如下标准的方法才能用于远程访问, 其余方法会被忽略:
//
// - 方法是导出的
// - 方法有两个参数, 都是导出类型或内建类型
// - 方法的第二个参数是指针
// - 方法只有一个error接口类型的返回值
//
// 事实上, 方法必须看起来像这样:
// func (t *T) MethodName(argType T1, replyType *T2) error
// 其中T, T1和T2都能被 encoding/gob 包序列化
// 这些限制即使使用不同的编解码器也适用
// (未来, 对定制的编解码器可能会使用较宽松一点的限制)
//
// 方法的第一个参数代表调用者提供的参数;
// 第二个参数代表返回给调用者的参数
// 方法的返回值, 如果非nil, 将被作为字符串回传, 在客户端看来就和errors.New创建的一样
// 如果返回了错误, 回复的参数将不会被发送给客户端
//
// 服务端可能会单个连接上调用ServeConn管理请求
// 更典型地, 它会创建一个网络监听器然后调用Accept
// 或者, 对于HTTP监听器, 调用HandleHTTP和http.Serve
//
// 想要使用服务的客户端会创建一个连接, 然后用该连接调用NewClient
//
// 更方便的函数Dial(DialHTTP)会在一个原始的连接(或HTTP连接)上依次执行这两个步骤
//
// 生成的Client类型值有两个方法, Call和Go, 它们的参数为要调用的服务和方法, 一个包含参数的指针, 一个用于接收结果的指针
//
// Call方法会等待远端调用完成, 而Go方法异步的发送调用请求并使用返回的Call结构体类型的Done通道字段传递完成信号
//
// 除非设置了显式的编解码器, 本包默认使用encoding/gob包来传输数据
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

// 两个整数参数
type Args struct {
	A, B int
}

// 商
// Quo除法 Rem取余
type Quotient struct {
	Quo, Rem int
}

type Arith int

// 执行乘法
// Args 参数结构体
// reply 保存结果
// error 返回值
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// 执行除法
// Args 参数结构体
// Quotient 保存结算结果
// error 返回值
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	// 注册Arith的RPC服务, 然后通过rpc.HandleHTTP函数把该服务注册到了HTTP协议上,
	// 然后我们就可以利用http的方式来传递数据了
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
