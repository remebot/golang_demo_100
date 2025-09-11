package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main_str() {
	str1 := "Hello World"
	str2 := `
反引号中的文本
会原封不动的输出
    aaaaaa`

	fmt.Println(str1)
	fmt.Println(str2)

	str3 := 'A'
	fmt.Println(str3)

	s := "hello世界"
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%c ", s[i])
	//}
	for _, r := range s {
		fmt.Print(string(r) + " ")
	}

	p1 := strings.HasPrefix(s, "he")
	p2 := strings.HasSuffix(s, "界")

	_ = p1
	println(p2)

	i := strings.Index(s, "世界")
	fmt.Println(i)

	strings.Replace(s, "世界", "world", 1)
	fmt.Println(s)

	runes := []rune(s)
	fmt.Println(len(runes))

	length := utf8.RuneCountInString(s)
	fmt.Println(length)

	//var a int = 10

	s = "hello世界"
	fmt.Println(string(s[0]))
	fmt.Println(string(s[5]))

}
