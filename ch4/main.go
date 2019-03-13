package main

import (
	"fmt"
)

func main() {
	array()
}

func array() {
	fmt.Println("=====4.1 array=====")
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Printf("q[2]=%d\n", q[2])
	fmt.Printf("r[2]=%d\n", r[2])

	q1 := [...]int{1, 2, 3}
	fmt.Printf("[...]int{1,2,3} %T\n", q1)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "ou", GBP: "GB", RMB: "yuan"}
	fmt.Println(RMB, symbol[RMB])

	r1 := [...]int{99: -1}
	fmt.Printf("len(r1)=%d, r1[99]=%d,r1[33]=%d\n", len(r1), r1[99], r1[33])

	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1 == b1, b1 == c1, c1 == a1)
}
