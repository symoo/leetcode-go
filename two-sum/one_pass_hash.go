package twosum
/*
onePassHash 通过一次循环
*/
func onePassHash(nums []int, target int) []int {
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if v, ok := table[complement]; ok != false {
			return []int{v, i}
		}
		table[nums[i]] = i
	}

	panic("No two sum solution")
}