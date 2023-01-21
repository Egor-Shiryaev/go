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
	scanner.Scan()
	line := scanner.Text()
	strSlice := strings.Split(line, " ")

	if len(strSlice) != 3 {
		fmt.Println(errors.New("некорректные входные данные"))
		return
	}

	num1 := strSlice[0]
	num2 := strSlice[2]
	operator := strSlice[1]

	checkOperator, _ := regexp.MatchString(`[^-+*/]`, operator)
	if checkOperator {
		fmt.Println(errors.New("калькутор умеет выполнть только операции: + - * /")) // true
		return
	}

	if !(checkArab(num1) && checkArab(num2) || checkRim(num1) && checkRim(num2)) {
		fmt.Println(errors.New("вычисляет одновременно только с римскими или арабскими цифрами"))
		return
	}

	if (checkArab(num1) && checkArab(num2)) && (strToNum(num1) > 10 || strToNum(num2) > 10) {
		fmt.Println(errors.New("число должно быть меньше 10"))
		return
	}

	if checkRim(num1) && checkRim(num2) {
		a := rimToArab(num1)
		b := rimToArab(num2)
		res := arifm(a, b, operator)
		if res <= 0 {
			fmt.Println(errors.New("результат работы с римскими числами должен быть больше нуля"))
			return
		} else {
			fmt.Println(arabToRim(res))
			return
		}
	}

	if checkArab(num1) && checkArab(num2) {
		fmt.Println(arifm(strToNum(num1), strToNum(num2), operator))
		return
	}
}

func checkRim(numStr string) bool {
	var arr = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, elem := range arr {
		if elem == numStr {
			return true
		}
	}
	return false
}

func checkArab(numStr string) bool {
	_, err := strconv.Atoi(numStr)
	if err != nil {
		return false
	}
	return true
}

// Из строки  "1" в число 1
func strToNum(str string) int {
	x, _ := strconv.Atoi(str)
	return x
}

// Арифметическая операция
func arifm(a int, b int, c string) int {
	var res int
	switch c {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}
	return res
}

func rimToArab(numArab string) int {
	var res int
	switch numArab {
	case "I":
		res = 1
	case "II":
		res = 2
	case "III":
		res = 3
	case "IV":
		res = 4
	case "V":
		res = 5
	case "VI":
		res = 6
	case "VII":
		res = 7
	case "VIII":
		res = 8
	case "IX":
		res = 9
	case "X":
		res = 10
	}
	return res
}

func arabToRim(num int) string {
	roman := ""
	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var i = len(romans) - 1

	for num > 0 {
		for numbers[i] <= num {
			roman += romans[i]
			num -= numbers[i]
		}
		i -= 1
	}

	return roman
}
