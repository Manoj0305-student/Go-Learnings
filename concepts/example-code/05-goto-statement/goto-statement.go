package gotostatement
import ("fmt")

func GotoStatement() {

	counter := 0

	loop: 
		counter++
		fmt.Print(counter," ")

	if counter <= 5 {
		goto loop
	}
}