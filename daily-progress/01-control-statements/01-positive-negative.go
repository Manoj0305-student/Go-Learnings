package controlstatementproblems
import(
	"fmt"
)

func PrintPostiveNegativeNumbers(num int) {
	if num > 0 {
		fmt.Println("Given Number is Positive:",num)
	} else if num < 0 {
		fmt.Println("Given Number is Negative:",num)
	} else {
		fmt.Println("Given Number is Zero:",num)
	}
}