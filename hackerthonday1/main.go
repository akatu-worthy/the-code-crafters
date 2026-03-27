package main

import (
	"fmt"
)

func main() {
	for {
		var first float64
		var second float64
		
		fmt.Println("\033[36m...WELCOME TO THE GOPHERS CALCULATOR...\033[0m")
		fmt.Println("\033[31m...YOU MAY START YOUR CALCULATION PROCESS...\033[0m")
		fmt.Println("\033[32mEnter number to be calculated:\033[0m")

		if _, err := fmt.Scan(&first); err != nil {
			fmt.Println("...Invalid input for first number...")
			continue
		}

		fmt.Println("Select Arithmetic Operator: ")
		var Operators string = ("1: +| 2: -| 3: *| 4: /| 5: Exit| 6: Help :")
		fmt.Println(Operators)
		fmt.Scan(&Operators)

		if Operators == "6" {
			fmt.Println("|add <a> + <b> → addition")
			fmt.Println("|sub <a> - <b> → subtraction")
			fmt.Println("|mul <a> * <b> → multiplication")
			fmt.Println("|div <a> / <b> → division")
			continue
		}
			
		if Operators == "5" {
			fmt.Println("Goodbye...")
			break
		}
		if Operators != "+" &&  Operators != "-" && Operators != "*" && Operators != "/" {
			fmt.Println("...Not a valid operator...")
			fmt.Println("...Try using Help operation...")
			continue
		}

		fmt.Println("Enter second number: ")
		if _, err := fmt.Scan(&second); err != nil {
			fmt.Println("...Invalid input for second number...")
			continue
		}

		fmt.Println("Result: ")

		switch Operators {

		case "1" :
			fmt.Println(first + second)

		case "2":
			fmt.Println(first - second)

		case "3":
			fmt.Println(first * second)

		case "4":
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
