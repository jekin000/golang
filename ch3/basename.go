// a example of shell basename
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
// example of slice
package main

import (
	"fmt"
	"strings"
)

func basename1(s string) string {
	//Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	//Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename1("a/b.c.go"))
	fmt.Println(basename2("a/b.c.go"))
	fmt.Println(basename2("a/b"))
}
