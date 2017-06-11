package easy

/**
	Two Sum
question:
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.
	Example:
	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
**/
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	l := len(nums)
	i := 0
	res := []int{}
	for i < l {
		r := target - nums[i]
		if m[r] != 0 && m[r]-1 != i {
			res = append(res, m[r]-1)
			res = append(res, i)
			break
		}
		m[nums[i]] = i + 1
		i += 1
	}
	return res
}
