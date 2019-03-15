package main

import "fmt"

func push(stack []int, v int) []int {
	stack = append(stack, v)
	return stack
}

func top(stack []int) int {
	return stack[len(stack)-1]
}

func pop(stack []int) []int {
	stack = stack[:len(stack)-1]
	return stack
}

func removeSameOrder(slice []int, i int) []int {
	//slice = [5,6,7,8,9]
	//slice[i:] = [7,8,9]
	//slice[i+1:] = [8,9]
	copy(slice[i:], slice[i+1:])
	//slice = [5,6,8,9,9]
	return slice[:len(slice)-1]
}

func removeNonOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s1 := []int{5, 6, 7, 8, 9}
	fmt.Println(removeSameOrder(s1, 2))
	s2 := []int{5, 6, 7, 8, 9}
	fmt.Println(removeNonOrder(s2, 2))

	stack := []int{1, 2, 3}
	stack = push(stack, 4)
	fmt.Println(top(stack))
	stack = pop(stack)
	fmt.Println(top(stack))
}
