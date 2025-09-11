package main

import "fmt"

func main() {
	// d := 05        // 前缀0表示八制进
	// e := 0o57      // 前缀0表示八制进
	// f := 0xab3     // 前缀0x表示十六进制
	// g := 5_0_123_7 // 表示数字 501237(_可以忽略)

	//fmt.Println(d, e, f, g)

	//fmt.Print("a")
	//fmt.Print("b")
	//fmt.Println("a")
	//fmt.Println("b")
	//fmt.Printf("a=%d,g=%f,n=%t\n", d, g, f)

	//var name string
	//var age int
	//
	//fmt.Scan(&name, &age)
	//fmt.Println(name, age)

	//var address string
	//fmt.Scanln(&address)
	//
	//fmt.Println(address)
	//
	//var b = 2
	//print(b)
	//
	//var name = "David"
	//fmt.Println(name)

	//var i, j, k = 1, 2, 3

	/*
		i, j, k := 1, 2, 3
		println(i, j, k)

		n := 3
		if n > 0 {
			fmt.Println("n 是正数")
		} else if n < 0 {
			fmt.Println("n 是负数")
		} else {
			fmt.Println("n 是零")
		}

		day := 2

		switch day {
		case 1:
			fmt.Println("星期一")
		case 2:
			fmt.Println("星期二")
		case 3:
			fmt.Println("星期三")
		case 4:
			fmt.Println("星期四")
		case 5:
			fmt.Println("星期五")
		case 6, 7:
			fmt.Println("周末")
		}*/

	i := 1
	sum := 0
	for i <= 100 {
		sum += i
		i++
	}
	println(sum)

	for {
		fmt.Println("死循环")
		break //不加 break 就会无限循环，相当于 while(true)
	}
}
