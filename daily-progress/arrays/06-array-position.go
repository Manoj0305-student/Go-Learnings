package arrays

/* 
 ReturnArrayPostion searches for a specific element in a slice and returns its index.

 Supported Test Cases:

 Case 1: The element exists in the array
 Input Array: [10, 25, 45, 8, 12, 6]
 Target Element to Search For: 8
 Expected Output: 3

 Case 2: The element does not exist in the array
 Input Array: [10, 25, 45, 8, 12, 6]
 Target Element to Search For: 99
 Expected Output: -1

 Case 3: The element appears multiple times
 Input Array: [5, 12, 7, 5, 3]
 Target Element to Search For: 5
 Expected Output: 0
*/
func ReturnArrayPostion(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1
}