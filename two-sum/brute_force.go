package twosum


// brute 暴力解法, 双重循环测试每一个数字相加是否等于结果
// 时间 O(n²) 
// 空间 O(1)
func brute(nums []int, target int) []int {
	n := len(nums)
	for i := 0 ; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	panic("No two sum solution")
}
