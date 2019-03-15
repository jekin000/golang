package main

import (
	"fmt"
	"sort"
)

func main() {
	array()
	Map()
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		//we can not use xv != y[k]
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func Map() {
	fmt.Println("=====4.3 map=====")
	ages1 := make(map[string]int) //mapping from strings to ints
	ages1["alice"] = 31
	ages1["charlie"] = 34
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	//empty map
	ages3 := map[string]int{}
	fmt.Println(ages3)
	fmt.Println(ages1["alice"])
	delete(ages1, "alice")

	ages1["bob"] = ages1["bob"] + 1
	fmt.Println(ages1["bob"])

	//The order is unsure
	for name, age := range ages1 {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//sored & sure order
	var names []string
	//also names := make([]string,0,len(ages1)) this is more effective
	for name := range ages1 {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages1[name])
	}

	var ages4 map[string]int
	fmt.Println(ages4 == nil)
	//for nil map , delete , query is ok, not insert will cause compile error
	//so please make/create a map before you insert it.

	if age, ok := ages4["bob"]; !ok {
		fmt.Println("pass")
	} else {
		fmt.Println(age)
	}

	fmt.Println(equal(ages1, ages2))
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))

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
