package controlstatementproblems
import
(
	"fmt"
)

func sumOfNnaturalNumbers() {
	
	sum := 0
	num := 5
	for i := 1 ; i < num ; i++ {
		sum +=i
		fmt.Println("Sum of N Natural numnbers:",sum)
	}
}
