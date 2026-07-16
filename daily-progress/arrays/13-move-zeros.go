package arrays

/* 
 MoveZerosAtEnd moves all 0's to the end of a slice in-place while maintaining 
 the relative order of the non-zero elements.

 Supported Test Cases:

 Case 1: Standard mix of zeros and non-zeros
 Input Array: [0, 1, 0, 3, 12]
 Expected Output: [1, 3, 12, 0, 0]

 Case 2: No zeros present in the array
 Input Array: [4, 5, 6, 7]
 Expected Output: [4, 5, 6, 7]

 Case 3: Array consists entirely of zeros
 Input Array: [0, 0, 0]
 Expected Output: [0, 0, 0]
*/

func MoveZerosAtEnd(nums []int) {

	j := 0

	for i:=0 ; i < len(nums) ; i++ {
		if nums[i] != 0 {
			nums[j],nums[i] = nums[i],nums[j]
			j++
		}
	}

}