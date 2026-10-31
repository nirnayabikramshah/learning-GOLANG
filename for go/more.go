package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	var IntNum uint = 1099
	IntNum = IntNum + 1
	fmt.Println(IntNum)

	var floatNum float64 = 23.4455
	floatNum = floatNum + 21.23
	fmt.Println(floatNum)

	name, age, class := "nirnaya", 15, 10
	fmt.Println(name, age, class)
	fmt.Println(len(name))

	const company string = "Google"
	fmt.Print(company)
	new()
}

func new() {
	const pi float32 = 3.14159265359
	fmt.Println(pi)
	var x int = 4
	var y int = 2
	var result int = intergerDivistion(x, y)
	fmt.Println(result)
}

func intergerDivistion(x int, y int) int {
	var result int = x / y
	return result

}
