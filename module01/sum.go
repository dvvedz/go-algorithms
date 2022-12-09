package module01

func Sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum
}

func RecursionSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + Sum(nums[1:])
}
