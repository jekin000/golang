//For book uncollectd example
package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unicode/utf8"
)

func main() {
	StringUTF8()
	StrAndInt()
	Const()
}

func Const() {
	fmt.Println("==>3.6 Const")
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)

	fmt.Println("====>3.6.1 iota")
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Println("====>3.6.2 untype const")
	const Pi64 float64 = math.Pi
	var x float32 = float32(Pi64)
	var y float64 = Pi64
	var z complex128 = complex128(Pi64)
	fmt.Printf("%T %[1]v\n", Pi64)
	fmt.Printf("%T %[1]v\n", x)
	fmt.Printf("%T %[1]v\n", y)
	fmt.Printf("%T %[1]v\n", z)

	var f float64 = 212
	fmt.Printf("%T %[1]v\n", (f-32)*5/9)
	fmt.Printf("%T %[1]v\n", 5/9*(f-32))     // 5/9 is an untyped integer
	fmt.Printf("%T %[1]v\n", 5.0/9.0*(f-32)) // 5.0/9.0 is an untyped float

	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')

}

func StrAndInt() {
	fmt.Println("==>3.5.5 StrToInt,IntToStr")
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))

	x, err := strconv.Atoi("123")               // x is an int
	y64, err := strconv.ParseInt("123", 10, 64) //base 10, up to 64bits
	fmt.Println(x, y64, err)
}

func StringUTF8() {
	fmt.Println("==>3.5.3 String UTF8")
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println("")

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	fmt.Println()

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println()

	n := 0
	for _, _ = range s {
		n++
	}
	fmt.Println("n =", n)
	fmt.Println("")

	n = 0
	for range s {
		n++
	}
	fmt.Println("n =", n)
	fmt.Println("")

	fmt.Printf("Unknown = \uFFFD\n")
	fmt.Println("")

	s = "ナタツスツ"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))
	fmt.Println("")
}
