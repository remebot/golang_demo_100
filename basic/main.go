package main

import "fmt"

func main() {
	d := 05        // 前缀0表示八制进
	e := 0o57      // 前缀0表示八制进
	f := 0xab3     // 前缀0x表示十六进制
	g := 5_0_123_7 // 表示数字 501237(_可以忽略)

	fmt.Println(d, e, f, g)

	//fmt.Print("a")
	//fmt.Print("b")
	//fmt.Println("a")
	//fmt.Println("b")
	fmt.Printf("a=%d,g=%f,n=%t\n", d, g, f)
}
