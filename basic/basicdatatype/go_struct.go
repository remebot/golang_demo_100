package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Gender  int
	Married bool
}

// 定义一个“成员函数”，方法接收者是 Person
func (p Person) SayHello() {
	fmt.Printf("大家好，我叫%s，我今年%d岁。\n", p.Name, p.Age)
}

// 定义一个带返回值的方法
func (p *Person) IsAdult() bool {
	return p.Age >= 18
}

// 定义一个指针接收者的方法，可以修改结构体内容
func (p *Person) GrowUp() {
	p.Age += 1
}

func main() {
	p := Person{
		Name:    "Bob",
		Age:     20,
		Gender:  1,
		Married: true,
	}
	// 访问 Name 字段
	fmt.Println(p.Name)
	// 修改 Age 字段
	p.Age = 21
	fmt.Println(p.Name + "今年" + strconv.Itoa(p.Age) + "岁了")

	// 使用成员函数
	p.SayHello()
	isAdult := p.IsAdult()
	if isAdult {
		fmt.Println(p.Name + "成年了")
	} else {
		fmt.Println(p.Name + "未成年")
	}
	p.GrowUp()
	fmt.Printf("生日过后，我 %d 岁啦\n", p.Age)

	var student struct {
		Name string
		Age  int
	}
	student.Name = "tom"
	student.Age = 18
	fmt.Println("我是匿名结构体对象 student，我的两个属性值为：", student.Name, student.Age)

}
