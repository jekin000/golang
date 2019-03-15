package main

import "fmt"

//reverse slice
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//equal array can compare, but slice not
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5} //array
	b := []int{0, 1, 2, 3, 4, 5}    //slice
	fmt.Printf("a=%T,b=%T\n", a, b)
	reverse(a[:])
	fmt.Println(a)

	fmt.Println()
	fmt.Println(b)
	reverse(b[:2])
	reverse(b[2:])
	reverse(b)
	fmt.Println(b)

	s1 := []string{"123", "456", "790"}
	s2 := []string{"abc", "456", "790"}
	s3 := []string{"123", "456", "790"}
	fmt.Println(equal(s1, s2))
	fmt.Println(equal(s1, s3))

	s4 := make([]string, 10)
	s5 := make([]string, 10, 20)
	fmt.Printf("s4=%T, s5=%T\n", s4, s5)

}
