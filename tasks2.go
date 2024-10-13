package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func convertBase(num string, fromBase int, toBase int) (string, error) {
	decimal, err := strconv.ParseInt(num, fromBase, 64)
	if err != nil {
		return "", err
	}

	if toBase == 10 {
		return strconv.FormatInt(decimal, 10), nil
	}

	return strconv.FormatInt(decimal, toBase), nil
}

func T1() {
	var num string
	var fromBase, toBase int

	fmt.Print("Введите число: ")
	fmt.Scanln(&num)

	fmt.Print("Введите исходную систему счисления: ")
	fmt.Scanln(&fromBase)

	fmt.Print("Введите конечную систему счисления: ")
	fmt.Scanln(&toBase)

	result, err := convertBase(num, fromBase, toBase)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Число %s в системе счисления %d равно %s в системе счисления %d\n", num, fromBase, result, toBase)
}

func solveQuadratic(a, b, c float64) (complex128, complex128) {
	discriminant := b*b - 4*a*c

	if discriminant >= 0 {
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return complex(root1, 0), complex(root2, 0)
	}

	realPart := -b / (2 * a)
	imagPart := math.Sqrt(-discriminant) / (2 * a)
	return complex(realPart, imagPart), complex(realPart, -imagPart)
}

func T2() {
	var a, b, c float64

	fmt.Print("Введите коэффициент a: ")
	fmt.Scanln(&a)
	fmt.Print("Введите коэффициент b: ")
	fmt.Scanln(&b)
	fmt.Print("Введите коэффициент c: ")
	fmt.Scanln(&c)

	root1, root2 := solveQuadratic(a, b, c)

	fmt.Printf("Корни уравнения: %v и %v\n", root1, root2)
}

func sortByAbsoluteValue(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})
	return arr
}

func T3() {
	arr := []int{-10, 2, -5, 3, -1, 0, 8}
	fmt.Println("Исходный массив:", arr)

	sortedArr := sortByAbsoluteValue(arr)
	fmt.Println("Отсортированный массив по модулю:", sortedArr)
}

func mergeSortedArrays(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}

func T4() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8}

	fmt.Println("Исходный массив 1:", arr1)
	fmt.Println("Исходный массив 2:", arr2)
	mergedArray := mergeSortedArrays(arr1, arr2)

	fmt.Println("Слияние массивов:", mergedArray)
}

func findSubstring(str, substr string) int {
	strLen := len(str)
	substrLen := len(substr)

	if substrLen > strLen {
		return -1
	}

	for i := 0; i <= strLen-substrLen; i++ {
		match := true
		for j := 0; j < substrLen; j++ {
			if str[i+j] != substr[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}

func T5() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]

	fmt.Print("Введите подстроку: ")
	substr, _ := reader.ReadString('\n')
	substr = substr[:len(substr)-1]

	index := findSubstring(str, substr)

	if index != -1 {
		fmt.Printf("Подстрока найдена на позиции %d\n", index+1)
	} else {
		fmt.Println("Подстрока не найдена")
	}
}

func T6() {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Print("Введите оператор (+, -, *, /, ^, %): ")
	fmt.Scan(&operator)

	var result float64
	var err error

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			err = fmt.Errorf("ошибка: деление на ноль")
		} else {
			result = a / b
		}
	case "^":
		result = math.Pow(a, b)
	case "%":
		if b == 0 {
			err = fmt.Errorf("ошибка: деление на ноль")
		} else {
			result = math.Mod(a, b)
		}
	default:
		err = fmt.Errorf("ошибка: недопустимая операция '%s'", operator)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Результат: %.2f\n", result)
	}
}

func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleaned := re.ReplaceAllString(s, "")

	cleaned = strings.ToLower(cleaned)

	reversed := reverse(cleaned)
	return cleaned == reversed
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func T7() {
	var input string

	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}

func isIntersecting(segment1, segment2, segment3 [2]int) bool {
	maxStart := max(segment1[0], max(segment2[0], segment3[0]))
	minEnd := min(segment1[1], min(segment2[1], segment3[1]))

	return maxStart <= minEnd
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func T8() {
	var segment1, segment2, segment3 [2]int

	fmt.Print("Введите начальную и конечную точки первого отрезка: ")
	fmt.Scan(&segment1[0], &segment1[1])
	fmt.Print("Введите начальную и конечную точки второго отрезка: ")
	fmt.Scan(&segment2[0], &segment2[1])
	fmt.Print("Введите начальную и конечную точки третьего отрезка: ")
	fmt.Scan(&segment3[0], &segment3[1])

	if isIntersecting(segment1, segment2, segment3) {
		fmt.Println("Существует область пересечения всех трех отрезков.")
	} else {
		fmt.Println("Область пересечения не существует.")
	}
}

func findLongestWord(sentence string) string {
	re := regexp.MustCompile(`[^\w\s]`)
	cleaned := re.ReplaceAllString(sentence, "")

	words := strings.Fields(cleaned)

	var longestWord string

	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}

	return longestWord
}

func T9() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите предложение: ")

	input, _ := reader.ReadString('\n')

	longestWord := findLongestWord(input)

	fmt.Printf("Самое длинное слово: %s\n", longestWord)
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func T10() {
	var year int

	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Println("Год является високосным.")
	} else {
		fmt.Println("Год не является високосным.")
	}
}

func fibonacciUpTo(n int) {
	a, b := 0, 1
	fmt.Print("Числа Фибоначчи до ", n, ": ")

	for a <= n {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func T11() {
	var n int

	fmt.Print("Введите целое число n: ")
	fmt.Scan(&n)

	fibonacciUpTo(n)
}

func primeNumbersInRange(start, end int) {
	if start > end {
		fmt.Println("Начальное число больше конечного. Пожалуйста, введите корректный диапазон.")
		return
	}
	fmt.Printf("Простые числа в диапазоне от %d до %d: ", start, end)
	for num := start; num <= end; num++ {
		if isPrime(num) {
			fmt.Print(num, " ")
		}
	}
	fmt.Println()
}

func T12() {
	var start, end int

	fmt.Print("Введите начальное число: ")
	fmt.Scan(&start)
	fmt.Print("Введите конечное число: ")
	fmt.Scan(&end)

	primeNumbersInRange(start, end)
}

func isArmstrong(num int) bool {
	sum := 0
	digits := 0
	temp := num

	for temp > 0 {
		temp /= 10
		digits++
	}

	temp = num
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}

	return sum == num
}

func armstrongNumbersInRange(start, end int) {
	if start > end {
		fmt.Println("Начальное число больше конечного. Пожалуйста, введите корректный диапазон.")
		return
	}

	fmt.Printf("Числа Армстронга в диапазоне от %d до %d: ", start, end)
	for num := start; num <= end; num++ {
		if isArmstrong(num) {
			fmt.Print(num, " ")
		}
	}
	fmt.Println()
}

func T13() {
	var start, end int

	fmt.Print("Введите начальное число: ")
	_, err := fmt.Scan(&start)
	if err != nil {
		fmt.Println("Ошибка при вводе начального числа:", err)
		return
	}

	fmt.Print("Введите конечное число: ")
	_, err = fmt.Scan(&end)
	if err != nil {
		fmt.Println("Ошибка при вводе конечного числа:", err)
		return
	}

	armstrongNumbersInRange(start, end)
}

func reverseString(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}

func T14() {
	var input string

	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)
	reversed := reverseString(input)
	fmt.Println("Реверс строки:", reversed)
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func T15() {
	var num1, num2 int

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Ошибка при вводе первого числа:", err)
		return
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Ошибка при вводе второго числа:", err)
		return
	}

	result := gcd(num1, num2)
	fmt.Printf("Наибольший общий делитель (%d, %d) = %d\n", num1, num2, result)
}
