package arrays

/* 
 MergeSortedArrays combines two pre-sorted slices into a single sorted slice.

 Supported Test Cases:

 Case 1: Standard alternating merge
 Input: arr_1 = [1, 3, 5], arr_2 = [2, 4, 6]
 Expected Output: [1, 2, 3, 4, 5, 6]

 Case 2: One array has completely smaller elements than the other
 Input: arr_1 = [1, 2, 3], arr_2 = [7, 8, 9]
 Expected Output: [1, 2, 3, 7, 8, 9]

 Case 3: Arrays are different sizes
 Input: arr_1 = [1, 5, 9, 10, 15], arr_2 = [2, 3]
 Expected Output: [1, 2, 3, 5, 9, 10, 15]
*/

func MergeSortedArrays(arr_1,arr_2 []int) []int {

	if len(arr_1) == 0 && len(arr_2) == 0 {
		return [] int {}
	}

	result := make([]int,len(arr_1)+len(arr_2))
	i,j,start := 0,0,0

	for i < len(arr_1) && j < len(arr_2) {

		if arr_1[i] < arr_2[j] {
			result[start] = arr_1[i];
			i++
		} else {
			result[start] = arr_2[j]
			j++
		}
		start++
	}

	for i < len(arr_1) {
		result[start] = arr_1[i]
		i++
		start++
	}

	for j < len(arr_2) {
		result[start] = arr_2[j]
		j++
		start++
	}

	return result

}