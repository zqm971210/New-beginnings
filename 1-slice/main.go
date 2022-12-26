package main

import "fmt"

func demo1() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

func demo3() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]
	fmt.Println(s1, len(s1), cap(s1)) // [2,3,4] len:3 cap:8
	fmt.Println(s2, len(s2), cap(s2)) // [4,5,6,7] len:4 cap:5

	s2 = append(s2, 100) // [4,5,6,7,100] len:5 cap:5
	s2 = append(s2, 200) // [4,5,6,7,100,200] len:6 cap:10
	fmt.Println(s2, len(s2), cap(s2))

	s1[2] = 20 // [2,3,20]

	fmt.Println(s1)    // [2,3,20]
	fmt.Println(s2)    // [4,5,6,7,100,200] len:6 cap:10
	fmt.Println(slice) // [0,1,2,3,20,5,6,7,100,9]
}

func main() {
	//demo2()
	var tag bool
	fmt.Println(tag)
}

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func demo2() {
	s := []int{1, 1, 1}
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	//myAppendPtr(&s)
	fmt.Println(s)
}
