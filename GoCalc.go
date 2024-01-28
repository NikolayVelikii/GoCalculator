package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func readInput() string {
	fmt.Print("Введите алгебраическое выражение используя числа от 1 до 10(арабские или римские): ")

	// Считываем строку с помощью метода ReadString
	// Делаем предположение, что введенный текст заканчивается символом новой строки
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// Разбиваем строку на слова
	text = strings.TrimSpace(text)
	// Выводим результат
	if text[0] == '-' {
		panic("Некорректный ввод: Отрицательное число")
	}
	return text
}

func splitByOperator(str string) (leftStr, rightStr string, operator rune) {
	operators := "+-*/"

	for _, op := range operators {
		index := strings.IndexRune(str, op)
		if index != -1 {
			leftStr := strings.TrimSpace(str[:index])
			rightStr := strings.TrimSpace(str[index+1:])

			return leftStr, rightStr, op
		}
	}

	panic("Некорректный ввод:  так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
}

func calculate(a, b int, operator rune) (result int) {
	switch operator {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			panic("деление на ноль!")
		}
		result = a / b
	default:
		panic("неподдерживаемый оператор")
	}

	return result
}

func isArabic(str string) bool {
	for _, char := range str {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func isRoman(str string) bool {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	for _, char := range str {
		if _, exists := romanNumerals[char]; !exists {
			return false
		}
	}
	return true
}

func parseRoman(str string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	upperStr := strings.ToUpper(str)
	total := 0
	prevValue := 0

	for _, char := range upperStr {
		value := romanNumerals[char]

		total += value

		if prevValue < value {
			total -= 2 * prevValue // Вычитаем предыдущее значение, так как оно уже было добавлено
		}

		prevValue = value
	}
	if total <= 10 {
		return total
	}

	panic("Некорректный ввод: введите число от I до X")

}

func parseArabic(str string) int {
	var num int
	_, err := fmt.Sscanf(str, "%d", &num)
	if err != nil {
		panic("Неверный формат арабского числа")
	}
	if num <= 10 && num >= 0 {
		return num
	}
	panic("Некорректный ввод: введите число от 1 до 10")
}

func intToRoman(num int) string {
	if num <= 0 {
		return "Недопустимое значение"
	}

	// Массивы символов римских цифр и их соответствующих значений
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	romanNumeral := ""
	i := 0

	// Пока число не станет равным нулю
	for num > 0 {
		// Если текущее число больше значения символа, добавляем символ и вычитаем его значение из числа
		if num >= values[i] {
			romanNumeral += symbols[i]
			num -= values[i]
		} else {
			// Если текущее число меньше значения символа, переходим к следующему символу
			i++
		}
	}

	return romanNumeral
}

func main() {
	// Вводим строку и убираем пробелы
	expression := readInput()
	// Разделяем строку на числа и знак
	left, right, operator := splitByOperator(expression)

	var result int
	var resultRoman string

	if isArabic(left) && isArabic(right) {
		// Арабские цифры
		result = calculate(parseArabic(left), parseArabic(right), operator)
		fmt.Println("Результат:", result)
	} else if isRoman(left) && isRoman(right) {
		// Римские цифры
		result = calculate(parseRoman(left), parseRoman(right), operator)
		resultRoman = intToRoman(result)
		fmt.Println("Результат:", resultRoman)
	} else {
		panic("Неправильный ввод. Используйте либо арабские, либо римские цифры.")

	}

}
