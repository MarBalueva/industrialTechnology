package main

import (
	"fmt"
	"math"
)

func triangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}

func Task1() {
	var height, base float64

	fmt.Print("Введите длину основания: ")
	fmt.Scan(&height)
	fmt.Print("Введите высоту: ")
	fmt.Scan(&base)
	square := triangleArea(base, height)
	fmt.Printf("Площадь треугольника: %f\n", square)
}

func sortArray(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// fmt.Printf("Шаг %d: %v\n", i+1, arr)

		if !swapped {
			break
		}
	}
	return arr
}

func Task2() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Исходный массив:", arr)

	sortArr := sortArray(arr)

	fmt.Println("Отсортированный массив:", sortArr)
}

func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			sum += i * i
		}
	}
	return sum
}

func Task3() {
	var num int

	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)

	fmt.Printf("Сумма квадратов всех чётных чисел от 1 до %d: %d\n", num, sumOfSquares(num))
}

func Task4() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}

func Task5() {
	var n int

	fmt.Print("Введите число:")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println("Число является простым")
	} else {
		fmt.Println("Число не является простым")
	}
}

func generatePrimes(limit int) []int {
	var arr []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			arr = append(arr, i)
		}
	}
	return arr
}

func Task6() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	result := generatePrimes(n)
	fmt.Printf("Простые числа до %d: %d\n", n, result)
}

func toBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		bit := n % 2
		binary = fmt.Sprintf("%d%s", bit, binary)
		n /= 2
	}
	return binary
}

func Task7() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	result := toBinary(n)
	fmt.Printf("Число %d в двоичной системе: %s\n", n, result)
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := math.MinInt

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max
}

func Task8() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Исходный массив:", arr)

	max := findMax(arr)

	fmt.Println("Максимальный элемент: ", max)
}

func Task9() {
	T15()
}

func sumArray(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func Task10() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Исходный массив:", arr)

	sum := sumArray(arr)

	fmt.Println("Сумма элементов массива: ", sum)
}
