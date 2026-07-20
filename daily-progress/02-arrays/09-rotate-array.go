package arrays


/*
	Imagine an array of length 3: [1, 2, 3].
	If someone asks you to rotate it right by 3 positions (K=3), the array shifts all the way around and ends up exactly how it started: [1, 2, 3]. If they ask to rotate by 4 (K=4), it's the exact same as rotating by 1.
*/

/* 
 RightRotate and LeftRotate shift the elements of a slice by k positions in-place.

 Supported Test Cases:

 Case 1: Rotate Right by k positions
 Input Array: [1, 2, 3, 4, 5], k = 2
 Expected Output: [4, 5, 1, 2, 3]

 Case 2: Rotate Left by k positions
 Input Array: [1, 2, 3, 4, 5], k = 2
 Expected Output: [3, 4, 5, 1, 2]

 Case 3: k is larger than the array length
 Input Array: [1, 2, 3], k = 4 (Equivalent to k = 1)
 Expected Output (Right): [3, 1, 2]
*/
func RightRotate(arr []int,k int) {

	if len(arr) == 0 {
		return
	}

	// Handle cases where k is larger than the array length
	k = k%len(arr)
	
	// reverse whole array first
	reverse(0,len(arr)-1,arr)
	// reverse upto k-1 elements in array
	reverse(0,k-1,arr)
	// reverse then k to len(arr)-1 in array
	reverse(k,len(arr)-1,arr)
}


func LeftRotate(arr []int,k int) {

	if len(arr) == 0 {
		return
	}

	// Handle cases where k is larger than the array length
	k = k%len(arr)

	// reverse 0 to k-1 elements
	reverse(0,k-1,arr)
	// reverse k to len(arr)-1 elements
	reverse(k,len(arr)-1,arr)
	// reverse whole array elements
	reverse(0,len(arr)-1,arr)
}

func reverse(start,end int,arr []int) []int {

	if len(arr) == 0 {
		return []int{}
	}
	
	for start < end {
		arr[start] , arr[end] = arr[end] , arr[start]
		start++
		end--
	}

	return arr
}