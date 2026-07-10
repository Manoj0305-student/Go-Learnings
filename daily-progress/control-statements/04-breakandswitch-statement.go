package controlstatementproblems

import(
	"fmt"
)

func printingNumbersUsingSwitchAndBreak() {

	number := 10
	for i := 1 ; i < number ; i++ {

		if i == 5 {
			continue
		} else if i == 8 {
			break
		}

		fmt.Println(i)
	}
} 