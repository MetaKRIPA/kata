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
	defer func() { // Обработчик паники для изящного завершения программы
		if r := recover(); r != nil {
			fmt.Println("Паника:", r)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите операцию:")
	if scanner.Scan() {
		input := scanner.Text()
		num1Str, op, num2Str := parseInput(input) // Изменено: убран возврат ошибки

		result := calculate(num1Str, op, num2Str) // Изменено: убран возврат ошибки

		fmt.Println("Результат:", result)
	}
}

// Разделяет ввод на части. В случае ошибки вызывает panic.
func parseInput(input string) (string, string, string) {
	parts := strings.Split(input, " ")
	if len(parts) == 3 {
		return parts[0], parts[1], parts[2]
	}

	for _, op := range []string{"+", "-", "*", "/"} {
		if strings.Contains(input, op) {
			parts := strings.Split(input, op)
			if len(parts) == 2 {
				return parts[0], op, parts[1]
			}
		}
	}

	panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
}

// Выполняет арифметическую операцию. В случае ошибки вызывает panic.
func calculate(num1Str, op, num2Str string) string {
	num1, isRoman1 := parseNumber(num1Str)
	num2, isRoman2 := parseNumber(num2Str)

	if isRoman1 != isRoman2 {
		panic("Используются одновременно разные системы счисления.")
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

	if isRoman1 && result < 1 {
		panic("в римской системе нет отрицательных чисел")
	} else if isRoman1 {
		for k, v := range romanToArabic {
			if v == result {
				return k
			}
		}
	}

	return strconv.Itoa(result)
}

// Парсит число из строки. В случае ошибки вызывает panic.
func parseNumber(str string) (int, bool) {
	if num, ok := romanToArabic[str]; ok {
		return num, true
	}
	num, err := strconv.Atoi(str)
	if err != nil || num < 1 || num > 10 {
		panic("некорректное число или число вне диапазона 1-10")
	}
	return num, false
}
