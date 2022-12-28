package main

import "fmt"

func main() {
	//make 只能用来分配及初始化类型为slice、map、chan 的数据。
	//new 可以分配任意类型的数据；
	//make 返回的是引用类型本身
	//new 分配返回的是指针，即类型*Type

	//var p int
	//var v *int
	//v = &p
	//*v = 11
	//fmt.Println(v)

	//var v *int
	//v = new(int)
	//*v = 8
	//fmt.Println(*v)

	//初始化一个指针变量，其值为nil，nil的值是不能直接赋值的
	var v *int
	//fmt.Println(*v)
	fmt.Println(v)
	v = new(int)
	fmt.Println(*v)
	fmt.Println(v)
}
