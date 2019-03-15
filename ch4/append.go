package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		//have room to grow
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}

func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	//same with appendInt2
	if zlen <= cap(x) {
		//have room to grow
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var x1 []int
	x1 = append(x1, 1)
	x1 = append(x1, 2, 3)
	x1 = append(x1, 4, 5, 6)
	x1 = append(x1, x1...)
	fmt.Println(x1)

	var x2 []int
	x2 = appendInt2(x2, 1)
	x2 = appendInt2(x2, 2, 3)
	x2 = appendInt2(x2, 4, 5, 6)
	x2 = appendInt2(x2, x2...)
	fmt.Println(x2)
}
