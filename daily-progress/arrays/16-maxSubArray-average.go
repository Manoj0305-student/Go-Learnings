package arrays
//nums = {1, 12, -5, -6, 50, 3}
//k = 4
func MaxSubArrayWithAverage(nums []int,k int) int {

	currentSum := 0

	for i := 0 ; i < k ; i++ {
		currentSum += nums[i]
	}

	maxSum := currentSum

	for i := k ; i < len(nums) ; i++ {
		currentSum = currentSum + nums[i] - nums[i-k]
	}

	if currentSum > maxSum {
		maxSum = currentSum
	}

	
	return maxSum/k

}