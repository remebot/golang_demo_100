package main

import "fmt"

func main1() {

	//fmt.Print("a")
	//fmt.Print("b")
	//fmt.Println("a")
	//fmt.Println("b")

	//a := 0xab3
	//g := 3.14
	//n := true
	//
	//fmt.Printf("a=%d,g=%f,n=%t", a, g, n)

	//k := 2.171241
	//fmt.Printf("k=%.2f", k)
	//
	//k2 := 2.17999
	//fmt.Printf("k2=%.2f", k2)
	//
	//fmt.Println("")
	////const A, B, C = 1, 2, 3
	////fmt.Println(A, B, C)
	//const (
	//	A = 1
	//	B = 2
	//	C = 3
	//)
	//fmt.Println(A, B, C)

	// 定义星期枚举
	const (
		Monday    = iota + 1 // 1
		Tuesday              // 2
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
		Saturday             // 6
		Sunday               // 7
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}
