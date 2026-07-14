package deferpanicrecover
import "fmt"

func SafeDivision(a,b int) {

	defer func() {
	
		if r := recover(); r != nil {
			fmt.Println("Recoverd:",r)
		}

 	}()

	if b == 0 {
		panic("Any number cannot be divided by Zero.")
	}

 result := a/b

 fmt.Println("Result:",result)

}

