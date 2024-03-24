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

func arabicToRoman(num int) string {
	if num < 1 {
		panic("Результат в римской системе меньше I невозможен")
	}

	var result strings.Builder
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	for _, numeral := range romanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
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
	// Используем Fields для учета любого количества пробелов
	parts := strings.Fields(input)
	// Ожидается ровно 3 части: операнд оператор операнд
	if len(parts) != 3 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *) с пробелом.")
	}
	return parts[0], parts[1], parts[2]
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

	if isRoman1 {
		if result < 1 {
			panic("в римской системе нет отрицательных чисел")
		}
		return arabicToRoman(result) // Используем функцию преобразования для результата
	}

	return strconv.Itoa(result)
}

// Парсит число из строки. В случае ошибки вызывает panic.
func parseNumber(str string) (int, bool) {
	if num, ok := romanToArabic[str]; ok {
		return num, true
	}
	num, err := strconv.Atoi(str)
	// Изменено: Уточнено сообщение об ошибке для случаев, когда число вне диапазона или некорректно
	if err != nil || num < 1 || num > 10 {
		panic(fmt.Sprintf("Число '%s' некорректно или вне диапазона 1-10", str))
	}
	return num, false
}
