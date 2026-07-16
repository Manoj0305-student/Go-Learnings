package arrays

/* 
 HasDuplicate checks if any value appears at least twice in a slice.

 Supported Test Cases:

 Case 1: Slice contains duplicates
 Input Array: [1, 2, 3, 1]
 Expected Output: true

 Case 2: Slice contains completely unique numbers
 Input Array: [1, 2, 3, 4]
 Expected Output: false

 Case 3: Empty slice or single element slice
 Input Array: [] or [7]
 Expected Output: false
*/
func HasDuplicate(nums []int) bool {

	seen := make(map[int]bool)
	for _,value := range nums {
		if seen[value] {
			return true
		}
		seen[value] = true
	}
	return false
}