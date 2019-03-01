//use popcount algorithm to calc how much bit 1 count a num have

package popcount

//8 bit count mapping table
//another style is:
/*
var pc[256] byte = func() (pc [256]byte){
    for i := range pc{
        pc[i] = pc[i/2] + byte(i&1)
    }
}()
*/
var pc [256]byte

func init() {
	// you could use for i,_ := range pc{
	for i := range pc {
		// if i is even, i/2 = i << 1, so there is no new added 1, so count(i) == count(i/2)
		// if i is odd,  i/2 = i << 1 + 1, so there is new added 1.
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount_Cycle(x uint64) (n int) {
	for i := uint(0); i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return
}
