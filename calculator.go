package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isRomanNumeral(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return false // Арабское число
	}
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, numeral := range romanNumerals {
		if numeral == s {
			return true // Римское число
		}
	}
	return false
}

func romanToArabic(s string) int {
	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	val := 0
	for len(s) > 0 {
		for key, value := range romanNumerals {
			if strings.HasPrefix(s, key) {
				val += value
				s = strings.TrimPrefix(s, key)
			}
		}
	}
	return val
}

func arabicToRoman(num int) string {

	if num <= 0 || num > 10 {
		return "Invalid input for Roman numerals conversion"
	}

	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	return romanNumerals[num-1]
}

func calculate(op string, a int, b int) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	fmt.Println(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение в формате: число операция число (например: 3 + 5)")
	input, _ := reader.ReadString('\n')

	expression := strings.Fields(input)
	if len(expression) != 3 {
		fmt.Println("Некорректный формат выражения")
		return
	}

	a, err1 := strconv.Atoi(expression[0])
	b, err2 := strconv.Atoi(expression[2])

	if err1 != nil || err2 != nil {
		if isRomanNumeral(expression[0]) && isRomanNumeral(expression[2]) {
			a = romanToArabic(expression[0])
			b = romanToArabic(expression[2])
		} else {
			fmt.Println("Некорректный ввод чисел")
			return
		}
	}

	if isRomanNumeral(expression[0]) != isRomanNumeral(expression[2]) {
		fmt.Println("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
		return
	}

	operation := expression[1]
	if (operation == "*" || operation == "/") && (a > 10 || b > 10) {
		fmt.Println("Калькулятор должен принимать на вход числа от 1 до 10 включительно")
		return
	}

	calculate(operation, a, b)
}
