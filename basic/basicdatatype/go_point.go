package main

import "fmt"

func main() {
	var a int = 100
	p := &a
	*p = 200
	fmt.Println(p)

	//a := 10
	//increase(&a)
	//fmt.Println(a)

}

func increase(n *int) {
	*n++
}
