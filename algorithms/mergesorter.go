package algorithms

func MergeSort(nums []int)  {
	size := len(nums)
	if size == 0 || size == 1 {
		return
	}
	left := nums[:size/2]
	right := nums[size/2:]
	MergeSort(left)
	MergeSort(right)
	i,j,k, tmpNums := 0,0,0, make([]int, size)
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmpNums[k] = left[i]
			i++
			k++
		}else{
			tmpNums[k] = right[j]
			j++
			k++
		}
	}
	for i < len(left) {
		tmpNums[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		tmpNums[k] = right[j]
		j++
		k++
	}
	copy(nums, tmpNums)
}
