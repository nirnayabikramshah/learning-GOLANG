package main

import "fmt"

func main() {
	var number_1 float64
	var number_2 float64
	var operator string
	fmt.Print("Enter the first number :")
	fmt.Scanln(&number_1)
	fmt.Print("Enter the  second number :")
	fmt.Scanln(&number_2)
	fmt.Print("Enter the operator :")
	fmt.Scanln(&operator)

	if operator == "+" {
		fmt.Println(number_1 + number_2)
	} else if operator == "-" {
		fmt.Println(number_1 - number_2)
	} else if operator == "/" {
		fmt.Println(number_1 / number_2)
	} else if operator == "*" {
		fmt.Println(number_1 * number_2)
	} else if operator == "" {
		fmt.Println("Please enter the numbers ,retry :")
	} else if operator == "/" {
		// Check for division by zero before dividing
		if number_2 == 0 {
			fmt.Println("Can't be divisible by 0")
		} else {
			fmt.Println(number_1 / number_2)
		}
	} else if operator == "End" {
		fmt.Println("Exiting program. Goodbye!")
		return
	} else {
		fmt.Println("Invalid operator entered")
	}
}
