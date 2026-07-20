package arrays

func MaxSumOfSubArray(nums []int) int {

	maxSum := nums[0]
	sum := nums[0]

	for i := 1 ; i < len(nums) ; i++ {

		if sum >= 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > maxSum {
			maxSum = sum
		}
		
	}

	return maxSum
	
}