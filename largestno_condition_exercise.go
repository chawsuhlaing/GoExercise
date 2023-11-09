package main
import "fmt"
func main () {
	number1, number2, number3 := 15, 12, 15
	fmt.Println(number1, number2, number3)
	if number1 >= number2 && number1 >= number3 {
		fmt.Println("Number 1 is the largest number among three numbers")
	} else if number2 >= number1 && number2 >= number3 {
		fmt.Println("Number 2 is the largest number among three numbers")
	} else {fmt.Println ("Number 3 is the largest number three numbers")}
}