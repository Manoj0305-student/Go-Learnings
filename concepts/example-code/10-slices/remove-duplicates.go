package slices

func RemoveDuplicates(arr []int) []int{

	if len(arr) == 0 {
		return []int{}
	}

	slice := make([]int,len(arr))

	j:=0
	//1,1,2,2,3
	slice[0] = arr[0]
	for i:=1 ; i < len(arr) ; i++ {
		if arr[i] != slice[j] {
			j++
			slice[j] = arr[i]
		}
	}
	return slice

}