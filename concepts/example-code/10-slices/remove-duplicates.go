package slices

func RemoveDuplicates(arr []int) []int{

	if len(arr) == 0 {
		return []int{}
	}

	slice := make([]int,len(arr))
	
	for i:=0 ; i < len(arr)-1 ; i++ {
		if arr[i] == arr[i+1] {
			slice[i] = arr[i]
		}
		slice[i] = arr[i]
	}
	return slice

}