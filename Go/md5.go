// md5包实现了MD5哈希算法
// hex包实现了16进制字符表示的编解码
package main

import (
	"fmt"
	"crypto/md5"
	"io"
	"encoding/hex"
)

func main() {
	// 返回数据data的MD5校验和
	b := []byte("This is a byte Hi")
	// [16]byte
	// [220 35 181 219 79 198 182 57 152 106 137 104 189 91 111 67]
	fmt.Println(md5.Sum(b))
	// string
	// dc23b5db4fc6b639986a8968bd5b6f43
	fmt.Println(fmt.Sprintf("%x", md5.Sum(b)))
	// 不能使用
	//fmt.Println(hex.EncodeToString(md5.Sum(b)))

	// 返回一个新的使用MD5校验的hash.Hash接口
	// 9e6d75cab4665a54ac5b483ecf1c7fae
	h := md5.New()
	io.WriteString(h, "Hi")
	io.WriteString(h, "Good")
	io.WriteString(h, "Do")
	fmt.Println(fmt.Sprintf("%x", h.Sum(nil)))
	//  通过比较可知, 第二个参数不影响生的32位字符内容, 跟使用hex.EncodeToString()后效果一致
	// 54686973206973206120627974659e6d75cab4665a54ac5b483ecf1c7fae
	fmt.Println(fmt.Sprintf("%x", h.Sum(b)))
	// 54686973206973206120627974659e6d75cab4665a54ac5b483ecf1c7fae
	fmt.Println(hex.EncodeToString(h.Sum(b)))

	fmt.Print("Hi\n")
}
