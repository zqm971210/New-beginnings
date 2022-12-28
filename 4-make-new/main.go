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
	fmt.Println(v)
	fmt.Println(*v)
}

//https://www.php.cn/be/go/465883.html#:~:text=go%E8%AF%AD%E8%A8%80%E4%B8%ADmake%E5%92%8Cnew%E7%9A%84%E5%8C%BA%E5%88%AB%E6%98%AF%E4%BB%80%E4%B9%88%EF%BC%9F,%E5%8C%BA%E5%88%AB%EF%BC%9A%E5%9C%A8go%E8%AF%AD%E8%A8%80%E4%B8%AD%EF%BC%8Cmake%E5%92%8Cnew%E9%83%BD%E6%98%AF%E5%86%85%E5%AD%98%E7%9A%84%E5%88%86%E9%85%8D%EF%BC%88%E5%A0%86%E4%B8%8A%EF%BC%89%EF%BC%8C%E4%BD%86%E6%98%AFmake%E5%8F%AA%E7%94%A8%E4%BA%8Eslice%E3%80%81map%E4%BB%A5%E5%8F%8Achannel%E7%9A%84%E5%88%9D%E5%A7%8B%E5%8C%96%EF%BC%88%E9%9D%9E%E9%9B%B6%E5%80%BC%EF%BC%89%EF%BC%9B%E8%80%8Cnew%E7%94%A8%E4%BA%8E%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%86%85%E5%AD%98%E5%88%86%E9%85%8D%EF%BC%8C%E5%B9%B6%E4%B8%94%E5%86%85%E5%AD%98%E7%BD%AE%E4%B8%BA%E9%9B%B6%E3%80%82%20make%E8%BF%94%E5%9B%9E%E7%9A%84%E6%98%AF%E5%BC%95%E7%94%A8%E7%B1%BB%E5%9E%8B%E6%9C%AC%E8%BA%AB%EF%BC%9B%E8%80%8Cnew%E8%BF%94%E5%9B%9E%E7%9A%84%E6%98%AF%E6%8C%87%E5%90%91%E7%B1%BB%E5%9E%8B%E7%9A%84%E6%8C%87%E9%92%88%E3%80%82
