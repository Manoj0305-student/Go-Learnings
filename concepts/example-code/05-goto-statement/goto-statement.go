package gotostatement
import ("fmt")

func gotoStatement() {

	counter := 0

	loop: 
		counter++
		fmt.Println(counter)

	if counter <= 5 {
		goto loop
	}
}