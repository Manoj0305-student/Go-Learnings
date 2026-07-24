package slices

import "fmt"

func SlicesUnderstanding() {
	var arr = [5]int{1,2,3,4,5}
	s1 := arr[0:2]
	s1 = append(s1, 90)

	fmt.Println(arr)
	fmt.Println(s1)
}