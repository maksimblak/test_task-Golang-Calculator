package main

import (
	"bufio"
	_ "errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для выполнения сложения
func add(a, b int) int {
	return a + b
}

// Функция для выполнения вычитания
func subtract(a, b int) int {
	return a - b
}

// Функция для выполнения умножения
func multiply(a, b int) int {
	return a * b
}

// Функция для выполнения деления
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на 0")
	}
	return a / b, nil
}

// Функция конвертации ввода
func convertInput(input string) (int, bool, error) {

	num, err := strconv.Atoi(input)
	if err != nil {
		return convertRomanToArabic(input), true, nil
	} else {
		return num, false, nil
	}
}

// Функция для преобразования римского числа в арабское
func convertRomanToArabic(roman string) int {
	arabic := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currentValue := getValueOfRomanDigit(roman[i])

		if currentValue < prevValue {
			arabic -= currentValue
		} else {
			arabic += currentValue
		}

		prevValue = currentValue
	}

	return arabic
}

// Функция для преобразования арабского числа в римское
func convertArabicToRoman(arabic int) string {
	if arabic < 1 || arabic > 3999 {
		return "Вне диапазона"
	}

	roman := ""
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(values); i++ {
		for arabic >= values[i] {
			arabic -= values[i]
			roman += symbols[i]
		}
	}

	return roman
}

// Функция для получения значения римской цифры
func getValueOfRomanDigit(digit byte) int {
	switch digit {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

// Основная функция программы main()
func main() {
	fmt.Println("Доступные операции: +, -, *, /")
	fmt.Println("Введите выражение: ")

	// Создание сканнера для считывания ввода
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Ошибка чтения ввода", scanner.Err())
		return
	}
	input := scanner.Text()

	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Ошибка, формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	}

	// преобразование операндов в числа
	var isRomanInput bool
	a, isRomanInput, err := convertInput(parts[0])
	if err != nil {
		fmt.Println("Ошибка конвертации первого числа:", err)
		return
	}
	b, isRomanInputB, err := convertInput(parts[2])
	if err != nil {
		fmt.Println("Ошибка конвертации второго числа")
		return
	}

	// Проверка на использование одновременно разных систем счисления

	if isRomanInput != isRomanInputB {
		panic("Ошибка, используются одновременно разные системы счисления.")

	}

	// Проверка на допустимый диапазон для арабских чисел
	if a <= 0 || a > 10 || b <= 0 || b > 10 {
		panic("Ошибка, числа должны быть в диапазоне от 1 до 10 включительно.")
	}

	// Проверка на допустимый диапазон для римских чисел
	if isRomanInput && (a < 1 || a > 10 || b < 1 || b > 10) {
		fmt.Println("Ошибка, римские числа должны быть в диапазоне от I до X включительно.")
		return
	}

	operator := parts[1]
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Ошибка, Недопустимая операция:", operator)
		return
	}

	if a <= 0 {
		if operator == "-" || operator == "*" || operator == "/" {
			fmt.Println("Ошибка, в римской системе нет отрицательных чисел или чисел меньше единицы.")
			return
		}
		result := convertArabicToRoman(add(a, b))
		if err != nil {
			fmt.Println("Ошибка конвертации результата")
			return
		}
		fmt.Println("Вывод:", result)
	} else {
		switch operator {
		case "+":
			result := add(a, b)
			if isRomanInput {
				resultRoman := convertArabicToRoman(result)
				if err != nil {
					fmt.Println("Ошибка конвертации результата")
					return
				}
				fmt.Println("Вывод:", resultRoman)
			} else {
				fmt.Println("Вывод:", result)
			}
		case "-":
			result := subtract(a, b)
			if isRomanInput {
				if result <= 0 {
					fmt.Println("Ошибка, в римской системе нет отрицательных чисел.")
					return
				}
				resultRoman := convertArabicToRoman(result)
				fmt.Println("Вывод:", resultRoman)
			} else {
				fmt.Println("Вывод:", result)
			}
		case "*":
			result := multiply(a, b)
			if isRomanInput {
				resultRoman := convertArabicToRoman(result)
				if err != nil {
					fmt.Println("Ошибка конвертации результата")
					return
				}
				fmt.Println("Вывод:", resultRoman)
			} else {
				fmt.Println("Вывод:", result)
			}
		case "/":
			if b == 0 {
				fmt.Println("Ошибка, деление на ноль")
				return
			}

			result, err := divide(a, b)
			if err != nil {
				fmt.Println("Ошибка при делении:", err)
				return
			}

			if isRomanInput {

				if result <= 0 {
					fmt.Println("в римской системе нет отрицательных чисел")
					return
				}

				if result == 0 {
					fmt.Println("в римской системе нет нуля")
					return
				}

				resultRoman := convertArabicToRoman(result)
				if err != nil {
					fmt.Println("Ошибка конвертации результата")
					return
				}
				fmt.Println("Вывод:", resultRoman)
			} else {
				fmt.Println("Вывод:", result)
			}

		}
	}
}
