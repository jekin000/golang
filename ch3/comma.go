//example of slice
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

//exe 3.10
func comma1(s string) string {
	var buf bytes.Buffer

	for i, v := range s {
		//Important
		if i > 0 && (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(v)
	}

	return buf.String()
}

func comma2(s string) string {
	var buf bytes.Buffer
	dot := strings.LastIndex(s, ".")

	var k string

	if s[0] == '+' || s[0] == '|' {
		k = s[1:]
		buf.WriteByte(s[0])
	} else {
		k = s
	}

	if dot == -1 {
		dot = len(k)
	}

	for i, v := range k {
		if i > 0 && dot-i > 0 && (dot-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(v)
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("12345678901234567890")) //12,345,678,901,234,567,890
	fmt.Println(comma1("12345678901234567890"))
	fmt.Println(comma2("12345678901234567890"))
	fmt.Println(comma2("1234.5434"))   //1,234.5434
	fmt.Println(comma1("123456"))      //123,456
	fmt.Println(comma2("123456.5434")) //123,456.5434
	fmt.Println(comma2("0.12345"))     //0.12345
	fmt.Println(comma2("+123456"))     //+123,456
}
