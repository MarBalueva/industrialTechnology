package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sumOfDigits(n int) int {
	n = int(math.Abs(float64(n)))
	if n == 0 {
		return 0
	}
	return n%10 + sumOfDigits(n/10)
}

func task1() {
	//задача 1
	fmt.Println("Задача 1")
	fmt.Println("Введите число: ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numberStr := scanner.Text()

	number, err := strconv.Atoi(numberStr)

	if err != nil {
		fmt.Println("Ошибка: введите корректное число")
		return
	}
	result := sumOfDigits(number)

	fmt.Println("Сумма цифр: ", result)
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func task2() {
	fmt.Println("Задача 2")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите температуру с указанием единицы (например, 25 (Celsius) или 77 (Fahrenheit)): ")

	scanner.Scan()
	input := scanner.Text()

	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		fmt.Println("Ошибка: неверный формат ввода. Введите число и единицу измерения (Celsius или Fahrenheit).")
		return
	}

	valueStr := parts[0]
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Println("Ошибка: неверный формат числа.")
		return
	}

	unit := strings.ToLower(parts[1])

	switch unit {
	case "celsius":
		fahrenheit := celsiusToFahrenheit(value)
		fmt.Printf("%.2f°C = %.2f°F\n", value, fahrenheit)
	case "fahrenheit":
		celsius := fahrenheitToCelsius(value)
		fmt.Printf("%.2f°F = %.2f°C\n", value, celsius)
	default:
		fmt.Println("Ошибка: укажите корректную единицу измерения (Celsius или Fahrenheit).")
	}
}

func doubleArray(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	return append([]int{arr[0] * 2}, doubleArray(arr[1:])...)
}

func task3() {
	fmt.Println("Задача 3")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите массив чисел через пробел: ")

	scanner.Scan()
	input := scanner.Text()

	strNumbers := strings.Split(input, " ")

	var numbers []int
	for _, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Ошибка: введите корректные числа.")
			return
		}
		numbers = append(numbers, num)
	}

	result := doubleArray(numbers)

	fmt.Println("Результат удвоения элементов массива: ", result)
}

func task4() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строки (введите пустую строку для завершения ввода):")

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	result := strings.Join(lines, " ")

	fmt.Println("Объединенная строка:", result)
}

func task5() {
	fmt.Println("Задача 6")
	var x1, y1, x2, y2 float64

	fmt.Println("Введите координаты первой точки (x1 y1):")
	fmt.Scan(&x1, &y1)
	fmt.Println("Введите координаты второй точки (x2 y2):")
	fmt.Scan(&x2, &y2)

	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("Расстояние между точками: %.2f\n", distance)

}

func task6() {
	fmt.Println("Задача 6")
	var x int

	fmt.Println("Введите число: ")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}

func task7() {
	fmt.Println("Задача 7")
	var x int

	fmt.Println("Введите год: ")
	fmt.Scan(&x)

	if x%4 == 0 {
		fmt.Println("Год високосный")
	} else {
		fmt.Println("Год невисокосный")
	}
}

func task8() {
	fmt.Println("Задача 8")
	var a, b, c int

	fmt.Println("Введите три числа:")
	fmt.Scan(&a, &b, &c)

	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	fmt.Printf("Наибольшее число: %d\n", max)
}

func task9() {
	fmt.Println("Задача 9")
	var age int

	fmt.Println("Введите возраст:")
	fmt.Scan(&age)

	if age >= 0 && age <= 12 {
		fmt.Println("Ребенок")
	} else if age >= 13 && age <= 17 {
		fmt.Println("Подросток")
	} else if age >= 18 && age <= 64 {
		fmt.Println("Взрослый")
	} else if age >= 65 {
		fmt.Println("Пожилой")
	} else {
		fmt.Println("Некорректный возраст")
	}
}

func task10() {
	fmt.Println("Задача 10")
	var x int

	fmt.Println("Введите число: ")
	fmt.Scan(&x)

	if x%3 == 0 && x%5 == 0 {
		fmt.Println("Число одновременно делится на 3 и на 5")
	} else {
		fmt.Println("Число одновременно не делится на 3 и на 5")
	}
}

func task11() {
	fmt.Println("Задача 11")
	var n int

	fmt.Println("Введите число:")
	fmt.Scan(&n)

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Printf("Факториал числа %d: %d\n", n, factorial)
}

func task12() {
	fmt.Println("Задача 12")
	var n int

	fmt.Println("Введите количество чисел Фибоначчи:")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Количество чисел должно быть больше 0")
		return
	}

	fib1, fib2 := 0, 1

	fmt.Print(fib1)
	if n > 1 {
		fmt.Print(", ", fib2)
	}

	for i := 3; i <= n; i++ {
		next := fib1 + fib2
		fmt.Printf(", %d", next)
		fib1 = fib2
		fib2 = next
	}
}

func task13() {
	fmt.Println("Задача 13")
	var n int

	fmt.Println("Введите количество элементов в массиве:")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Исходный массив:", arr)

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println("Перевёрнутый массив:", arr)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func task14() {
	fmt.Println("Задача 14")
	var n int

	fmt.Println("Введите число:")
	fmt.Scan(&n)

	fmt.Printf("Простые числа до %d: ", n)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}

func task15() {
	fmt.Println("Задача 15")
	var n, sum int

	fmt.Println("Введите количество элементов в массиве:")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Исходный массив:", arr)

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println("Сумма элементов массива:", sum)
}

func main() {
	// task1()
	// task2()
	// task3()
	// task4()
	// task5()
	// task6()
	// task7()
	// task8()
	// task9()
	// task10()
	// task11()
	// task12()
	// task13()
	// task14()
	// task15()
	T15()
}
