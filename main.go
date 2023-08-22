package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func isRoman(input string) bool {
	romanNumerals := map[string]struct{}{
		"I": {}, "II": {}, "III": {}, "IV": {}, "V": {},
		"VI": {}, "VII": {}, "VIII": {}, "IX": {}, "X": {},
	}
	_, isRoman := romanNumerals[input]
	return isRoman
}

func main() {
	for {
		fmt.Println("Введите выражение (или введите 'exit' для выхода):")
		reader := bufio.NewReader(os.Stdin)
		expression, _ := reader.ReadString('\n')
		expression = strings.TrimSpace(expression)

		if expression == "exit" {
			fmt.Println("Программа завершена.")
			break
		}

		// Проверяем наличие римских цифр
		if isRoman(expression) && strings.ContainsAny(expression, "1234567890") {
			fmt.Println("Ошибка: Используются одновременно разные системы счисления.")
			continue
		}

		// Проверяем, что строка не пуста и содержит пробелы
		if len(expression) == 0 || !strings.Contains(expression, " ") {
			fmt.Println("Ошибка: Строка не является математической операцией.")
			continue
		}

		if !isValidExpression(expression) {
			fmt.Println("Ошибка: Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			continue
		}
		fmt.Println("Введено выражение:", expression)

		numbers, operator, err := parseInput(expression)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		// Проверяем, что используется одна система счисления (или римская, или арабская)
		isRomanNumeral := isRoman(expression)
		if isRomanNumeral && strings.ContainsAny(expression, "1234567890") {
			fmt.Println("Ошибка: Используются одновременно разные системы счисления.")
			continue
		}
		if (isRomanNumeral && numbers[0] <= 0) || (!isRomanNumeral && (numbers[0] < 1 || numbers[0] > 10)) {
			fmt.Println("Ошибка: Число вне диапазона от 1 до 10:", numbers[0])
			continue
		}

		if isRomanNumeral && operator == "+" {
			fmt.Println("Ошибка: Используются одновременно разные системы счисления.")
			continue
		}

		if operator == "-" && numbers[0] <= numbers[1] {
			fmt.Println("Ошибка: В римской системе нет отрицательных чисел.")
			continue
		}

		result, err := calculate(numbers[0], numbers[1], operator)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		printResult(result, isRomanNumeral)
	}
}

func isValidExpression(expression string) bool {
	validOperators := map[string]struct{}{
		"+": {}, "-": {}, "*": {}, "/": {},
	}

	split := strings.Fields(expression)
	if len(split) != 3 {
		return false
	}

	// Проверяем, что второй элемент - оператор, и что он действительно допустимый
	_, validOperator := validOperators[split[1]]
	if !validOperator {
		return false
	}

	return true
}

func parseInput(input string) ([]int, string, error) {
	split := strings.Fields(input)
	if len(split) != 3 {
		return nil, "", fmt.Errorf("неверное выражение")
	}

	num1, err := stringToInt(split[0])
	if err != nil {
		return nil, "", err
	}

	num2, err := stringToInt(split[2])
	if err != nil {
		return nil, "", err
	}

	return []int{num1, num2}, split[len(split)-2], nil
}

func stringToInt(s string) (int, error) {
	if val, ok := romanMap[s]; ok {
		return val, nil
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("неверное число: %s", s)
	}

	if n < 1 || n > 10 {
		return 0, fmt.Errorf("число вне диапазона от 1 до 10: %s", s)
	}

	return n, nil
}

func calculate(a, b int, op string) (int, error) {
	if op == "/" {
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	}

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	default:
		return 0, fmt.Errorf("неверный оператор: %s", op)
	}
}

func intToRoman(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("Римские числа могут быть только положительными")
	}

	var result strings.Builder

	for n > 0 {
		for _, numeral := range []struct {
			Value  int
			Symbol string
		}{
			{1000, "M"},
			{900, "CM"},
			{500, "D"},
			{400, "CD"},
			{100, "C"},
			{90, "XC"},
			{50, "L"},
			{40, "XL"},
			{10, "X"},
			{9, "IX"},
			{5, "V"},
			{4, "IV"},
			{1, "I"},
		} {
			if n >= numeral.Value {
				result.WriteString(numeral.Symbol)
				n -= numeral.Value
				break
			}
		}
	}

	return result.String(), nil
}

func printResult(result int, isRomanNumeral bool) {
	if isRomanNumeral {
		romanResult, err := intToRoman(result)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(romanResult)
	} else {
		fmt.Println(result)
	}
}
