package arrays

/* 
 TwoSum finds the indices of two numbers in a slice that add up to a specific target.

 Supported Test Cases:

 Case 1: Standard positive numbers
 Input: nums = [2, 7, 11, 15], target = 9
 Expected Output: [0, 1]

 Case 2: Targets requiring the same value at different indices
 Input: nums = [3, 2, 4], target = 6
 Expected Output: [1, 2]

 Case 3: Handling negative integers
 Input: nums = [-3, 4, 3, 90], target = 0
 Expected Output: [0, 2]
*/
func TwoSum(arr []int, target int) []int {

	elements := make(map[int]int)

	for index , value := range arr {

		result := target - value

		if i , found := elements[result] ; found {

			return []int{index,i}

		}

		elements[value] = index
	}


	return []int{}

}