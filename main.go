// sorry i just know about this submit too late, the code below not from mine, i just quick copy from 1 guy in Group chat pull request, and fix some code of him to make it work better, sorry about this, later i will update it faster


package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("")
	fmt.Print(">")

	for scanner.Scan() {
		text := scanner.Text()
		var result int64

		// Format Input
		expressionString := formatExpression(text)

		// Parse Argument
		num1, num2, operator, err := parseArgs(expressionString)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Caculation
		result, err = caculator(num1, num2, operator)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(expressionString+" =", result)
		fmt.Print(">")
	}
}

func formatExpression(text string) string {
	// Trim Space
	formatedText := strings.TrimSpace(text)
	// Remove multiple space
	regex := regexp.MustCompile(" +")
	formatedText = regex.ReplaceAllString(formatedText, "")
	// Add Space before Operator [Ex: 9*5 => 9 *5]
	regex = regexp.MustCompile("(\\d)([*+-/])")
	formatedText = regex.ReplaceAllString(formatedText, "$1 $2")
	// Add Space after Operator [Ex: 9*5 => 9* 5]
	regex = regexp.MustCompile("([*+-/])(\\d)")
	formatedText = regex.ReplaceAllString(formatedText, "$1 $2")

	return formatedText
}

func parseArgs(text string) (a, b int64, opt string, err error) {
	args := strings.Split(text, " ")
	if len(args) != 3 {
		err =  errors.New("Invalid arguments!")
		return
	}
	opt = args[1]
	a, err = strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return
	}
	b, err = strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return
	}
	return
}

func caculator(a, b int64, opt string) (result int64, err error) {

	switch opt {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0.0, errors.New("Err : Divided by zero!")
		}
		result = a / b
	default:
		return 0.0, errors.New("Err : Invalid Operator!")
	}
	return
}