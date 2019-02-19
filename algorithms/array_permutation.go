/*
given array list {2,3,5,7}, given target 7, please return {{2,2,3}, {2,5},{7},}
*/
package algorithms

func permutationNums(nums []int, n int) [][]int  {
	return permutation_nums(0, nums, n, []int{})
}

func permutation_nums(pos int, nums []int, n int, ret []int) [][]int {
	var res [][]int
	if n == 0 {
		res = append(res, ret)
		return res
	}
	for i:=pos;i<len(nums);i++ {
		if n >= nums[i] {
			tmp := make([]int, len(ret))
			copy(tmp, ret)
			res = append(res, permutation_nums(i, nums, n-nums[i], append(tmp, nums[i]))...)
		}
	}
	return res
}