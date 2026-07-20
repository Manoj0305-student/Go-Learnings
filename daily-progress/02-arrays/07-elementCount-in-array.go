package arrays

/* 
 ReturnCountOfElementInArray counts how many times a target element appears in a slice.

 Supported Test Cases:

 Case 1: Element appears multiple times
 Input Array: [4, 2, 8, 4, 9, 4, 1]
 Target Element: 4
 Expected Output: 3

 Case 2: Element appears exactly once
 Input Array: [10, 20, 30, 40]
 Target Element: 30
 Expected Output: 1

 Case 3: Element is not in the array at all
 Input Array: [1, 2, 3, 5]
 Target Element: 99
 Expected Output: 0 
*/
func ReturnCountOfElementInArray(arr []int, target int) (int) {

	count := 0
	for _,value := range arr {
		if value == target {
			count++
		}
	}
	return count
}