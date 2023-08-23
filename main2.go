//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//// Словарь для преобразования  римских чисел в арабские
//var romanNumerals = map[string]int{
//	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
//	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
//}
//
//// Словарь для преобразования арабских числел в римские
//var arabicNumerals = map[int]string{
//	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
//	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
//}
//
//// Функция для выполнения сложения
//func add(a, b int) int {
//	return a + b
//}
//
//// Функция для выполнения вычитания
//func subtract(a, b int) int {
//	return a - b
//}
//
//// Функция для выполнения умножения
//func multiply(a, b int) int {
//	return a * b
//}
//
//// Функция для выполнения деления
//func divide(a, b int) (int, error) {
//	if b == 0 {
//		return 0, fmt.Errorf("деление на ноль")
//	}
//	return a / b, nil
//}
//
//// Функция для преобразования римского числа в арабское
//func convertRomanToArabic(roman string) (int, error) {
//	arabic, found := romanNumerals[roman]
//	if !found {
//		return 0, fmt.Errorf("неверное римское число")
//	}
//	return arabic, nil
//}
//
//// Функция для преобразования арабского числа в римское
//func convertArabicToRoman(arabic int) (string, error) {
//	if arabic <= 0 {
//		return "", fmt.Errorf("результат не может быть отрицательным или нулевым")
//	}
//
//	roman, found := arabicNumerals[arabic]
//	if !found {
//		return "", fmt.Errorf("неверное арабское число")
//	}
//	return roman, nil
//}
//
//// Основная функция программы main()
//func main() {
//	fmt.Println("Доступные операции: +, -, *, /")
//	fmt.Println("Введите выражение: ")
//
//	// Создание сканнера для считывания ввода
//	scanner := bufio.NewScanner(os.Stdin)
//	if !scanner.Scan() {
//		fmt.Println("Ошибка чтения ввода", scanner.Err())
//		return
//	}
//	input := scanner.Text()
//
//	// Разбиение введенной строки на операнды и оператор
//	parts := strings.Fields(input)
//	if len(parts) != 3 {
//		fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
//		return
//	}
//
//	// преобразование операндов в числа
//	var isRomanInput bool
//	a, err := strconv.Atoi(parts[0]) // Преобразование строки в число
//	if err != nil {
//		// Если не удалось преобразовать в число, попытаться преобразовать в римское число
//		romanA, convErr := convertRomanToArabic(parts[0])
//		if convErr != nil {
//			fmt.Println("Ошибка", convErr)
//			return
//		}
//		a = romanA
//		isRomanInput = true
//	}
//	// Аналогично для второго операнда
//	b, err := strconv.Atoi(parts[2])
//	if err != nil {
//		romanB, convErr := convertRomanToArabic(parts[2])
//		if convErr != nil {
//			fmt.Println("Ошибка", convErr)
//			return
//		}
//		b = romanB
//		isRomanInput = true
//	}
//	// Проверка на использование одновременно разных систем счисления
//	if (a <= 0 && b > 0) || (a > 0 && b <= 0) {
//		fmt.Println("Используются одновременно разные системы счисления.")
//		return
//	}
//
//	//if (isRomanInput && b > 0) || (!isRomanInput && b <= 0) {
//	//	fmt.Println("Вывод ошибки:", "используются одновременно разные системы счисления.")
//	//	return
//	//}
//
//	// Проверка введенного оператора
//	operator := parts[1]
//	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
//		fmt.Println("Недопустимая операция:", operator)
//		return
//	}
//
//	// Обработка случая, когда хотя бы один операнд - римское число
//	if a <= 0 {
//		// Обработка операций с римскими числами
//		if operator == "-" || operator == "*" || operator == "/" {
//			fmt.Println("В римской системе нет отрицательных чисел или чисел меньше единицы.")
//			return
//		}
//		// Выполнение операции и вывод результата в римских числах
//		result, convErr := convertArabicToRoman(add(a, b))
//		if convErr != nil {
//			fmt.Println("Ошибка при конвертации результата в римские числа", convErr)
//			return
//		}
//		fmt.Println(result)
//	} else {
//		// Обработка арабских чисел
//		switch operator {
//		case "+":
//			result := add(a, b)
//
//			// Вывод результата в римских числах, если ввод был в этой системе
//			if isRomanInput {
//				resultRoman, convErr := convertArabicToRoman(result)
//				if convErr != nil {
//					fmt.Println("Ошибка при конвертации результата в римские числа:", convErr)
//					return
//				}
//				fmt.Println("Вывод:", resultRoman)
//			} else {
//				fmt.Println("Вывод:", result)
//			}
//		case "-":
//			result := subtract(a, b)
//			if isRomanInput {
//				// Обработка результата и вывод в римских числах
//				if result <= 0 {
//					fmt.Println("Ошибка", "В римской системе нет отрицательных чисел.")
//					return
//				}
//				resultRoman, convErr := convertArabicToRoman(result)
//				if convErr != nil {
//					fmt.Println("Ошибка при конвертации результата в римские числа", convErr)
//					return
//				}
//				fmt.Println("Вывод:", resultRoman)
//			} else {
//				fmt.Println("Вывод:", result)
//			}
//		case "*":
//			result := multiply(a, b)
//			if isRomanInput {
//				resultRoman, convErr := convertArabicToRoman(result)
//				if convErr != nil {
//					fmt.Println("Ошибка при конвертации результата в римские числа", convErr)
//					return
//				}
//				fmt.Println("Вывод:", resultRoman)
//			} else {
//				fmt.Println("Вывод:", result)
//			}
//		case "/":
//			result, err := divide(a, b)
//			if err != nil {
//				fmt.Println("Ошибка", err)
//				return
//			}
//			if isRomanInput {
//				resultRoman, convErr := convertArabicToRoman(result)
//				if convErr != nil {
//					fmt.Println("Ошибка при конвертации результата в римские числа", convErr)
//					return
//				}
//				fmt.Println("Вывод:", resultRoman)
//			} else {
//				fmt.Println("Вывод:", result)
//			}
//		}
//	}
//}
