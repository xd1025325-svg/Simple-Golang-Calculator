package main

import  (
	"fmt"
	"strconv"
)
func main() {
    var num1 string
	var num2 string
	var operator string
	var result float64
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	var num1conv, err = strconv.Atoi(num1)
	if err != nil {
		fmt.Println("Error converting first number")
		return
	}
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
	var num2conv, err2 = strconv.Atoi(num2)
	if err2 != nil {
		fmt.Println("Error converting second number")
		return
	}
	fmt.Print("Enter operator (add, rest, multiply, divide): ")
	fmt.Scanln(&operator)
	switch operator {
	case "add":
		result = float64(num1conv) + float64(num2conv)
	case "rest":
		result = float64(num1conv) - float64(num2conv)
	case "multiply":
		result = float64(num1conv) * float64(num2conv)
	case "divide":
		result = float64(num1conv) / float64(num2conv)
	}
	fmt.Printf("Result: %v\n", result)
}