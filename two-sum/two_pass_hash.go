package twosum

/* Two-pass hash table
通过空间换时间使查询时间从o(n)降至o(1)
哈希表可以使查询接近常数时间
之所以是接近而不是达到常数是因为如果出现哈希冲突查询时间会退化至o(n)
如果使用较好的哈希函数可以使查询时间摊销至O(1)
简单的实现是使用两次迭代 第一次把元素的值和下标添加到表中
第二次检查每个元素的补数(target-nums[i])在表中是否存在 
注意补数不能是nums[i]本身
时间: O(n) 迭代两次, 由于哈希表的查询效率为O(1), 所以时间复杂度为O(n)
空间: O(n) 额外的空间依赖于元素的数量即N
*/
func twoPass(nums []int, target int) []int {
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		table[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if v, ok := table[complement]; ok != false && v != i {
			return []int{i, v}
		}
	}
	panic("No two sum solution")
}