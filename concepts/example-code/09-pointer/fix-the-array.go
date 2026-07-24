package pointer

import "fmt"

func mainFun() {
	arr  := [...]int{1,2,3,4,5}
	modifyTheArray(arr[:])
	fmt.Println(arr)
}

func modifyTheArray(arr []int) {

	arr[2] = 100

}