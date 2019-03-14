package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	c3 := []byte("X")
	c4 := [...]byte{1, 2, 3}
	fmt.Printf("len(c3)=%d\n", len(c3))
	fmt.Printf("disp_by_x = %x\n%T\n", c3, c3) // 58
	fmt.Println(c3)                            //[88]
	fmt.Printf("len(c4)=%d\n", len(c4))
	fmt.Printf("%T\n", c4)

	zero32_1(&c1)
	zero32_1(&c2)
	fmt.Printf("%x\n%x\n", c1, c2)
}

func zero32_1(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero32_2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
