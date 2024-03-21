package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите операцию:")
	if scanner.Scan() {
		input := scanner.Text()
		num1Str, op, num2Str, err := parseInput(input)
		if err != nil {
			fmt.Println("Паника:", err)
			return
		}

		result, err := calculate(num1Str, op, num2Str)
		if err != nil {
			fmt.Println("Паника:", err)
			return
		}

		fmt.Println("Результат:", result)
	}
}

// Разделяет ввод на части и возвращает их
func parseInput(input string) (string, string, string, error) {
	parts := strings.Split(input, " ")
	if len(parts) == 3 {
		return parts[0], parts[1], parts[2], nil
	}

	// Пытаемся обработать ввод без пробелов
	for _, op := range []string{"+", "-", "*", "/"} {
		if strings.Contains(input, op) {
			parts := strings.Split(input, op)
			if len(parts) == 2 {
				return parts[0], op, parts[1], nil
			}
		}
	}

	return "", "", "", fmt.Errorf("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
}

// Выполняет арифметическую операцию
func calculate(num1Str, op, num2Str string) (string, error) {
	num1, isRoman1, err := parseNumber(num1Str)
	if err != nil {
		return "", err
	}

	num2, isRoman2, err := parseNumber(num2Str)
	if err != nil {
		return "", err
	}

	if isRoman1 != isRoman2 {
		return "", fmt.Errorf("Используются одновременно разные системы счисления.")
	}

	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	if isRoman1 {
		if result < 1 {
			return "", fmt.Errorf("в римской системе нет отрицательных чисел")
		}
		for k, v := range romanToArabic {
			if v == result {
				return k, nil
			}
		}
	}

	if isRoman1 {
		for k, v := range romanToArabic {
			if v == result {
				return k, nil
			}
		}
	}

	return strconv.Itoa(result), nil
}

// Парсит число из строки
func parseNumber(str string) (int, bool, error) {
	if num, ok := romanToArabic[str]; ok {
		return num, true, nil
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, false, fmt.Errorf("некорректное число")
	}
	return num, false, nil
}
