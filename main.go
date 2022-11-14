package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	for {
		var equation string
		fmt.Scanln(&equation)

		if equation == "exit" || equation == "Exit" || equation == "EXIT" {
			break
		}

		if (strings.Count(equation, "+") + strings.Count(equation, "-") +
			strings.Count(equation, "*") + strings.Count(equation, "/")) != 1 {
			fmt.Println("Invalid Input")
			continue
		}

		var operator string
		if strings.Contains(equation, "+") {
			operator = "+"
		} else if strings.Contains(equation, "-") {
			operator = "-"
		} else if strings.Contains(equation, "*") {
			operator = "*"
		} else {
			operator = "/"
		}

		var operatorIndex int = strings.Index(equation, operator)
		var left string = equation[:operatorIndex]
		var right string = equation[operatorIndex+1:]

		var err error
		var leftNumber int
		var rightNumber int
		leftNumber, err = strconv.Atoi(left)
		rightNumber, err = strconv.Atoi(right)
		fmt.Println(leftNumber+rightNumber, err)
	}
}
