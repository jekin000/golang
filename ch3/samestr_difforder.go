// exe 3.12
// Judge two str, if they have same charactor but different order, then return true
package main

import (
	"fmt"
)

//asc
func sort(s string) string {
	r := []rune(s)

	for i, _ := range r {
		k := i
		for j := i + 1; j < len(r); j++ {
			if r[j] < r[k] {
				k = j
			}
		}
		if k != i {
			r[i], r[k] = r[k], r[i]
		}
	}

	return string(r)
}

func isSameDiff(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return sort(s1) == sort(s2)
}

func main() {
	s1 := "ZxzyQAabB"
	s2 := "xzQAZabyB"
	fmt.Printf("%s vs %s, result is %v\n", s1, s2, isSameDiff(s1, s2))

	s1 = "abc"
	s2 = "abcd"
	fmt.Printf("%s vs %s, result is %v\n", s1, s2, isSameDiff(s1, s2))

}
