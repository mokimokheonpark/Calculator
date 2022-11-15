package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	for true {

		var equation string
		fmt.Scanln(&equation)

		if equation == "exit" || equation == "Exit" || equation == "EXIT" {
			break
		}

		var operator string
		var operatorIndex int

		if (strings.Count(equation, "+") + strings.Count(equation, "*") + strings.Count(equation, "/")) == 1 {
			if strings.Contains(equation, "+") {
				operator = "+"
			} else if strings.Contains(equation, "*") {
				operator = "*"
			} else {
				operator = "/"
			}
			operatorIndex = strings.Index(equation, operator)
		} else if (strings.Count(equation, "+") + strings.Count(equation, "*") + strings.Count(equation, "/")) == 0 {
			operator = "-"
			if strings.Count(equation, "-") == 1 || (strings.Count(equation, "-") == 2 && strings.Contains(equation, "--")) {
				operatorIndex = strings.Index(equation, operator)
			} else if strings.Count(equation, "-") == 3 || (strings.Count(equation, "-") == 2 && !strings.Contains(equation, "--")) {
				var occurrence int = 0
				for i := 0; i < len(equation); i++ {
					if equation[i] == 45 {
						operatorIndex = i
						occurrence++
					}
					if occurrence == 2 {
						break
					}
				}
			} else {
				fmt.Println("Error")
				continue
			}
		} else {
			fmt.Println("Error")
			continue
		}

		var left string = equation[:operatorIndex]
		var right string = equation[operatorIndex+1:]
		leftNumber, _ := strconv.ParseFloat(left, 32)
		rightNumber, _ := strconv.ParseFloat(right, 32)

		if operator == "+" {
			fmt.Println(leftNumber + rightNumber)
		} else if operator == "-" {
			fmt.Println(leftNumber - rightNumber)
		} else if operator == "*" {
			fmt.Println(leftNumber * rightNumber)
		} else {
			fmt.Println(leftNumber / rightNumber)
		}
	}
}
