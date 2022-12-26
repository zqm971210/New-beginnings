package main

import (
	"context"
	"fmt"
)

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	startx := 0
	starty := 0
	offset := 1
	num := 1
	for loop := 0; loop < n/2; loop++ {
		i, j := startx, starty
		for ; j < n-offset; j++ {
			ans[i][j] = num
			num++
		}
		for ; i < n-offset; i++ {
			ans[i][j] = num
			num++
		}
		for ; j > startx; j-- {
			ans[i][j] = num
			num++
		}
		for ; i > starty; i-- {
			ans[i][j] = num
			num++
		}
		offset++
		startx++
		starty++
	}
	if n%2 == 1 {
		mid := n / 2
		ans[mid][mid] = num
	}
	return ans
}

func spiralOrder1(matrix [][]int) []int {
	row, col := len(matrix), len(matrix[0])
	ans := make([]int, row*col)
	n, loop, num, startx, starty, offset := 0, 0, 0, 0, 0, 1
	if row > col {
		n = col
	} else {
		n = row
	}
	for ; loop < n/2; loop++ {
		i, j := startx, starty
		for ; j < col-offset; j++ {
			ans[num] = matrix[i][j]
			num++
		}
		for ; i < row-offset; i++ {
			ans[num] = matrix[i][j]
			num++
		}
		for ; j > startx; j-- {
			ans[num] = matrix[i][j]
			num++
		}
		for ; i > startx; i-- {
			ans[num] = matrix[i][j]
			num++
		}
		startx++
		starty++
		offset++
	}

	if n%2 == 1 {
		mid := n / 2
		if row > col {
			for i := mid; i < mid+row-col+1; i++ {
				ans[num] = matrix[i][mid]
				num++
			}
		} else {
			for i := mid; i < mid+col-row+1; i++ { //1+1+1
				ans[num] = matrix[mid][i]
				num++
			}
		}
	}
	return ans
}

func spiralOrder(matrix [][]int) []int {
	var ans []int
	row, col := len(matrix), len(matrix[0])
	startx, starty := 0, 0
	offset := 1
	loop := mix(row, col) / 2
	for ; loop > 0; loop-- {
		i, j := startx, starty
		for ; j < col-offset; j++ {
			ans = append(ans, matrix[i][j])
		}
		for ; i < row-offset; i++ {
			ans = append(ans, matrix[i][j])
		}
		for ; j > starty; j-- {
			ans = append(ans, matrix[i][j])
		}
		for ; i > startx; i-- {
			ans = append(ans, matrix[i][j])
		}
		startx++
		starty++
		offset++
	}

	n := mix(row, col) // n = 1
	if n%2 == 1 {
		mid := mix(row, col) / 2
		if row > col {
			for i := mid; i < mid+row-col+1; i++ {
				ans = append(ans, matrix[i][mid])
			}
		} else {
			for i := mid; i < mid+col-row+1; i++ {
				ans = append(ans, matrix[mid][i])
			}
		}
	}
	return ans
}

func mix(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//注意：

//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。

//示例 1：
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"

func minWindow1(s string, t string) string {
	// map 滑动窗口来做
	sMap, tMap := map[byte]int{}, map[byte]int{}
	sLen, tLen := len(s), len(t)
	//初始化t窗口
	// for _,v := range t {
	//     tMap[v] ++
	// }

	//上面初始化 v是rune 不匹配
	for i := 0; i < tLen; i++ {
		tMap[t[i]]++
	}

	left, right, lenth := 0, -1, sLen+1
	//left ，right 记录的是  最小窗口的左右指针
	for i, j := 0, 0; j < sLen; j++ {
		// i ，j 为左右滑动窗口指针
		//记录  且一直扩张
		sMap[s[j]]++

		//满足条件
		for judge(sMap, tMap) {
			//取最小窗口
			if j-i+1 < lenth {
				lenth = j - i + 1
				left = i
				right = j
			}
			//收缩窗口
			sMap[s[i]]--
			i++
		}
	}
	//因为切片左闭右开 所以 right+1 所以right初始化为-1
	return s[left : right+1]
}

func minWindow(s string, t string) string {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	sLen, tLen := len(s), len(t)
	for i := 0; i < tLen; i++ {
		tMap[t[i]]++
	}
	left, right, anslen := 0, -1, sLen+1
	for begin, end := 0, 0; end < sLen; end++ {
		sMap[s[end]]++
		for judge(sMap, tMap) {
			// 取最小窗口
			windowlen := end - begin + 1
			if windowlen < anslen {
				anslen = windowlen
				left = begin
				right = end
			}
			//收缩窗口
			sMap[s[begin]]--
			begin++
		}
	}
	return s[left : right+1]
}

func judge(sMap, tMap map[byte]int) bool {
	for k, tv := range tMap {
		if sMap[k] < tv {
			return false
		}
	}
	return true
}

//判断 s是否完全覆盖t
func judge1(sMap, tMap map[byte]int) bool {
	for k, tv := range tMap {
		if sMap[k] < tv {
			return false
		}
	}
	return true
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			cancel()
			break
		}
	}

	//strings.Contains()
}
