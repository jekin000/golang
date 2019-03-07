//iota
package main

import (
	"fmt"
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota
	FlagBroadcast          //10
	FlagLoopback           //100
	FlagPointToPoint       //1000
	FlagMulticast          //10000
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

//https://stackoverflow.com/questions/34124294/writing-powers-of-10-as-constants-compactly/34145667
const (
	KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24
)

const (
	KB1, MB1, GB1, TB1, PB1, EB1, ZB1, YB1 = 1000, KB1 * KB1, MB1 * KB1, GB1 * KB1, TB1 * KB1, PB1 * KB1, EB1 * KB1, ZB1 * KB1
)

const (
	x, KB2, MB2, GB2, TB2, PB2, EB2, ZB2, YB2 = 1000, x, KB2 * x, MB2 * x, GB2 * x, TB2 * x, PB2 * x, EB2 * x, ZB2 * x
)

const (
	x3, KB3, MB3, GB3, TB3, PB3, EB3, ZB3, YB3 = 'Ϩ', x3, x3 * x3, MB3 * x3, GB3 * x3, TB3 * x3, PB3 * x3, EB3 * x3, ZB3 * x3
)

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))

	//ZiB fail
	fmt.Println(EiB)

	//YB
	fmt.Println(YB)

	//ZB1 fail
	fmt.Println(EB1)

	//ZB2 fail
	fmt.Println(EB2)

	//TB3 overflow
	//fmt.Println(TB3)
}
