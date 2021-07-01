package main

import "fmt"

// Practice, append single param to the end of a slice
func appendIntPractice(x []int, y int) []int {
	var z []int
	// 因为是要追加一个元素到切片中，所以需要 + 1
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 此时，x 切片仍旧有空间，直接追加即可
		// z 此时比 x 多了一个元素，此元素为默认值 0，起到占位置作用
		z = x[:zlen]
	} else {
		// 此时，slice 已经没有空间了，需要重新分配一个底层数组
		// 容量扩展 * 2
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		// 按照新的 zcap 和 zlen 重新创建 z 切片
		z = make([]int, zlen, zcap)
		// 将 x 切片中的元素复制到新 z 中
		copy(z, x)
	}
	// 将 y 复制到上述占位置的元素上，len(x)-1 才是未扩展切片的最后一个元素
	z[len(x)] = y
	return z
}

// Practice, append slice to the end of another slice
func appendSlicePractice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen < cap(x) {
		// x 中还有足够空间容纳 y 中元素
		// 占位元素，默认值为 0
		z = x[:zlen]
	} else {
		// x 中没有足够空间容纳 y 中元素
		// 扩展容量 * 2 倍
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	// 将 y 复制到上述占位置的元素上，len(x)-1 才是未扩展切片的最后一个元素
	copy(z[len(x):], y)
	return z
}

// Practice, remove ont param of slice
func remove(s []int, n int) []int {
	copy(s[n:], s[n+1:])
	// 因为是重用底层数组，所以此时 s 为 {5, 6, 8, 9, 9}，最后的 9 元素是原底层数组的元素
	// 所以需要返回 s[:len(s)-1] 舍弃掉后面的
	return s[:len(s)-1]
}

func main() {
	var x, y []int

	// 追加单个元素到 slice 末尾
	for i := 0; i < 10; i++ {
		y = appendIntPractice(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	// 追加多个元素到 slice 末尾
	 m := []int{1, 2, 3}
	 n := []int{4, 5, 6}
	 var h []int

	 h = appendSlicePractice(m, n...)
	 fmt.Printf("h is: %v \n", h)

	 // 使用 slice 模拟 stack, 删除 slice 中一个元素
	 s := []int{5, 6 , 7, 8 ,9}
	 remove(s, 2)

}