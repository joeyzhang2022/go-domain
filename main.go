/**
package main


import "fmt"

func main() {
	fmt.Println("hello world!")
}
*/

package main

import (
	"fmt"
)

// Add01 实现2数相加
// 面向过程
func Add01(a, b int) int {
	return a + b
}

// 类型+方法
// 面向对象，方法：给某个类型绑定一个函数
type long int

// Add02 tmp叫接收者，接收者就是传递的一个参数
func (tmp long) Add02(other long) long {
	return tmp + other
}

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

func (TestPerson Person) getName() string {
	return TestPerson.name
}

type Org struct {
	dep    string
	name   string
	person Person
}

func (TestOrg Org) getName() string {
	return TestOrg.person.getName()
}

func main() {
	var result int
	result = Add01(1, 1) //普通函数调用方式
	fmt.Println("result = ", result)
	//定义一个变量
	var a long = 2
	//调用方法格式： 变量名.函数(所需参数)
	r := a.Add02(3)
	fmt.Println("r = ", r)
	//面向对象只是换了一种表现形式

	var testorg Org

	testorg.dep = "YFYB"
	testorg.name = "研发一部"
	testorg.person.sex = 'N'
	testorg.person.name = "zhy"

	fmt.Println("r = ", testorg.getName())

}
