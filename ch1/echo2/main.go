//Echo2 prints its comand-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// _ is index
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
