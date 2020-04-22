package mergeMiddleWare

import (
	"fmt"
	"testing"
)

func TestMergeN(t *testing.T) {
	ch1 := RandomArrs(10)
	ch2 := RandomArrs(10)
	ch3 := RandomArrs(10)
	rch1 := IntsSortsToChan(ch1)
	rch2 := IntsSortsToChan(ch2)
	rch3 := IntsSortsToChan(ch3)
	rrch := MergeN(rch1, rch2, rch3)
	for v := range rrch {
		fmt.Println(v)
	}
}
