/*
给定一个数据流，数据流长度N很大，且N直到处理完所有数据之前都不可知，请问如何在只遍历一遍数据（O(N)）的情况下，能够随机选取出m个不重复的数据。

1.长度很大且未可知(N)．
2.时间复杂度O(n).
3.随机选取ｍ个数字，被选中的概率为m/N.

解题思路
看代码，很巧妙．

*/
package algorithms

import "math/rand"

func Reservior(dataStream []string, m int) []string {
	var res []string
	size := len(dataStream)
	if size < m {
		return dataStream
	}
	for i:=0;i<m;i++ {
		res = append(res, dataStream[i])
	}
	for i:=m;i<size;i++ {
		d := rand.Int63n(int64(i))
		if d < int64(m) {
			res[uint(d)] = dataStream[i]
		}
	}
	return res
}
