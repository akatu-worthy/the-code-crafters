package main

import (
	"fmt"
)

func main() {
	for {
		var first float64
		var second float64
		
		fmt.Println("...WELCOME TO THE GOPHERS CALCULATOR...")
		fmt.Println("...YOU MAY START YOUR CALCULATION PROCESS...")
		fmt.Println("Enter number to be calculated: ")

		if _, err := fmt.Scan(&first); err != nil {
			fmt.Println("...Invalid input for first number...")
			continue
		}

		fmt.Println("Select Arithmetic Operator: ")
		var Operators string = "(+, -, *, /, exit): "
		fmt.Println(Operators)
		fmt.Scan(&Operators)
		
		if Operators == "exit" {
			fmt.Println("Goodbye...")
			break
		}

		fmt.Println("Enter second number: ")
		if _, err := fmt.Scan(&second); err != nil {
			fmt.Println("...Invalid input for second number...")
			continue
		}

		fmt.Println("Result: ")

		switch Operators {

		case "+":
			fmt.Println(first + second)

		case "-":
			fmt.Println(first - second)

		case "*":
			fmt.Println(first * second)

		case "/":
			if second == 0 {
				fmt.Println("...Cannot divide by zero...")
			} else {
				fmt.Println(first / second)
			}

		default:
			fmt.Println("...Invalid operator...")
		}

		continue
	}
}
