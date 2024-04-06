package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()

	fmt.Println("Введите выражение:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n') // читаем введенный текст
	text = strings.TrimSpace(text)

	getResult(text)

}

func getResult(text string) {
	if strings.Contains(text, "+") { // сложение
		parts := strings.Split(text, " + ")
		if len(parts) == 2 {
			sufString(parts)
			for _, part := range parts {
				if part == "" {
					panic("Выражение не должно содержать пустых строк")
				} else if len(part) > 10 {
					panic("Строка должна содержать не более 10 символов")
				}
			}
			fmt.Println("\"" + parts[0] + parts[1] + "\"")
		} else {
			panic("Выражение должно иметь вид \"строка\" оператор \"строка\"")
		}

	} else if strings.Contains(text, "-") { // вычетание
		parts := strings.Split(text, " - ")
		if len(parts) == 2 {
			sufString(parts)
			for _, part := range parts {
				if part == "" {
					panic("Выражение не должно содержать пустых строк")
				} else if len(part) > 10 {
					panic("Строка должна содержать не более 10 символов")
				}
			}
			result := strings.Replace(parts[0], parts[1], "", -1)
			fmt.Println("\"" + result + "\"")
		} else {
			panic("Выражение должно иметь вид \"строка\" оператор \"строка\"")
		}

	} else if strings.Contains(text, "*") { // умножение
		parts := strings.Split(text, " * ")
		if len(parts) != 2 {
			panic("Выражение должно иметь вид \"строка\" * число")
		}

		// Проверяем, заключена ли строка в кавычки
		if !strings.HasPrefix(parts[0], "\"") || !strings.HasSuffix(parts[0], "\"") {
			panic("Строка должна быть в двойных ковычках")
		}
		// Убираем кавычки перед проверкой длины
		parts[0] = strings.Trim(parts[0], "\"")

		if len(parts[0]) > 10 {
			panic("Строка должна содержать не более 10 символов")
		}

		val, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("Второй операнд должен быть целым числом")
		}
		if val < 1 || val > 10 {
			panic("Множитель должен быть в диапазоне 1-10")
		}

		result := repeatString(parts[0], val)
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Printf("\"%s\"\n", result)

	} else if strings.Contains(text, "/") { // деление
		parts := strings.Split(text, " / ")
		if len(parts) != 2 {
			panic("Выражение должно иметь вид \"строка\" / число")
		}
		if !strings.HasPrefix(parts[0], "\"") || !strings.HasSuffix(parts[0], "\"") {
			panic("Строка должна быть в двойных ковычках")
		}
		parts[0] = strings.Trim(parts[0], "\"")
		if len(parts[0]) > 10 {
			panic("Строка должна содержать не более 10 символов")
		}
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("Второй операнд должен быть целым числом")
		}
		if val < 1 || val > 10 {
			panic("Делитель должен быть в диапазоне 1-10")
		}
		segmentLength := len(parts[0]) / val
		result := parts[0][:segmentLength]
		fmt.Printf("\"%s\"\n", result)

	} else {
		panic("Выражение должно иметь вид операнд оператор операнд")
	}

}

func sufString(parts []string) []string {
	for index, peace := range parts {
		if strings.HasPrefix(peace, "\"") && strings.HasSuffix(peace, "\"") {
			parts[index] = strings.Trim(peace, "\"")
		} else {
			panic("Все строки должны быть заключены в кавычки")
		}
	}
	return parts
}

func repeatString(s string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += s
	}
	return result
}
