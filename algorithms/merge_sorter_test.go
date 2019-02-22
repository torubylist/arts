package algorithms

import (
	"testing"
	"reflect"
	"fmt"
)

func TestMergeSort(t *testing.T) {
	tests := []struct{
		nums []int
		exp []int
	} {
		{[]int{2,5,1,9,10}, []int{1,2,5,9,10}},
		{[]int{3,5,10,7}, []int{3,5,7,10}},

	}
	for i, test := range tests {
		MergeSort(test.nums)
		if reflect.DeepEqual(test.nums, test.exp) {
			fmt.Printf("test%d pass!", i)
		}else {
			fmt.Printf("test%d failed, %v expected, but got %v", i, test.exp, test.nums)
		}

	}
}
