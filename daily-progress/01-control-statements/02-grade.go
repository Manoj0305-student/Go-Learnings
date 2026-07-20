package controlstatementproblems

import(
	"fmt"
)

func gradesUsingSwitchCase(grade rune) {

	switch grade {
		
	  case 'A':
			fmt.Println("Excellent")

		case 'B':
			fmt.Println("Good")
	  
		case 'C':
			fmt.Println("Average")
		
		case 'D':
			fmt.Println("Poor")

		case 'F':
			fmt.Println("Fail")

		default:
			fmt.Println("Some other grade")

	}
}