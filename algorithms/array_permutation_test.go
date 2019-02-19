package algorithms

import (
	"testing"
	"fmt"
)

func TestPermutationNums(t *testing.T)  {
	tests := []struct{
		nums []int
		n int
		exp [][]int
	}{
		{[]int{2,3,5,7}, 7, [][]int{{2,2,3}, {2,5}, {7},},},
	}
	for i, test := range tests {
		ret := permutationNums(test.nums, test.n)
		fmt.Println(ret)
		t.Logf("test%d pass", i)
	}
}
