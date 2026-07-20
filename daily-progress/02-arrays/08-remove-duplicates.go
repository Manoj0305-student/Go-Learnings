package arrays

/* 
 SortedArrayCount counts the number of unique elements in a sorted slice.

 Supported Test Cases:

 Case 1: Array has multiple duplicates scattered around
 Input Array: [1, 2, 2, 3, 4, 4, 4, 5]
 Expected Output: 5

 Case 2: Array has all identical elements
 Input Array: [7, 7, 7, 7]
 Expected Output: 1

 Case 3: Array already consists entirely of unique elements
 Input Array: [10, 20, 30, 40]
 Expected Output: 4
*/

func SortedArrayCount(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	uniqueCount := 0

	for i := 0 ; i < len(arr)-1 ; i++ {

		if arr[i] != arr[i+1] {
			uniqueCount++
		}
	}
	return uniqueCount+1
}

/* 
 UnsortedArrayCount counts the number of unique elements in an unsorted slice.

 Supported Test Cases:

 Case 1: Array has multiple duplicates scattered around
 Input Array: [4, 2, 8, 4, 9, 4, 1]
 Expected Output: 5 (unique: 4, 2, 8, 9, 1)

 Case 2: Array has all identical elements
 Input Array: [7, 7, 7, 7]
 Expected Output: 1

 Case 3: Array already consists entirely of unique elements
 Input Array: [10, 20, 30, 40]
 Expected Output: 4
*/

// In Coming Lesson we will see detailed expalination of maps concept 

func UnsortedArrayCount(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	uniqueCount := 0
	unique := make(map[int]bool)

	for _,value := range arr {
		if !unique[value] {
			unique[value] = true
			uniqueCount++
		}
	}

	return uniqueCount

}