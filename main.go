package main

import "fmt"

func main() {
	// var name string
	// name = "Ashabul Elaph Hilful"

	// // println creates a new line after execution
	// fmt.Println(name)
	// fmt.Println("Welcome !!")

	// // Simply prints without  going to new line
	// fmt.Printf("Text in same line!!!")
	// fmt.Printf("SEE!!!")

	// var message, message2 = "Wanna know a Secret??", "Let's do it"
	// fmt.Printf("Let me tell you: ow wait %v,%v ,Still trying ?", message, message2)

	// fmt.Println()

	// var value int

	// fmt.Scan(&value)
	// fmt.Println(`Your typed number is :`, value)

	// fmt.Scanln(&value)
	// fmt.Println("Your second value is :", value)

	// fmt.Scanf("%d", &value) //formated scan
	// fmt.Println("Your last value is :", value)

	var first, second int = 10, 20

	var fltNumber = float64(first + second + 30)
	fmt.Println("Sum of the numbers:", fltNumber)

}
