package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

func isPrime2(num int) (bool, int) {
	if num < 2 {
		return false, -1
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false, i
		}
	}
	return true, -1
}

func Task_1() {
	var number int

	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Ошибка при вводе числа:", err)
		return
	}

	isPrimeNumber, divisor := isPrime2(number)
	if isPrimeNumber {
		fmt.Printf("Число %d является простым.\n", number)
	} else {
		fmt.Printf("Число %d не является простым. Первый делитель: %d\n", number, divisor)
	}
}

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		fmt.Printf("Шаг %d: %v\n", i+1, arr)

		if !swapped {
			break
		}
	}
}

func Task_3() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Исходный массив:", arr)

	bubbleSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}

func Task_4() {
	size := 10

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

var memo = make(map[int]int)

func fibonacci(n int) int {
	if val, exists := memo[n]; exists {
		return val
	}

	if n <= 1 {
		return n
	}

	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

func Task_5() {
	var n int

	fmt.Print("Введите число для вычисления Фибоначчи: ")
	fmt.Scan(&n)

	fmt.Printf("Число Фибоначчи для %d: %d\n", n, fibonacci(n))
}

func reverseNumber(num int) int {
	reversed := 0

	for num != 0 {
		lastDigit := num % 10
		reversed = reversed*10 + lastDigit
		num /= 10
	}

	return reversed
}

func Task_6() {
	var num int

	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)

	fmt.Printf("Число в обратном порядке: %d\n", reverseNumber(num))
}

func pascalTriangle(levels int) {
	triangle := make([][]int, levels)

	for i := 0; i < levels; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	for i := 0; i < levels; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", triangle[i][j])
		}
		fmt.Println()
	}
}

func Task_7() {
	var levels int
	fmt.Print("Введите количество уровней треугольника Паскаля: ")
	fmt.Scan(&levels)

	pascalTriangle(levels)
}

func isPalindromeNumber(num int) bool {
	if num < 0 {
		return false
	}

	original := num
	reversed := reverseNumber(num)

	return original == reversed
}

func Task_8() {
	var num int

	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)

	if isPalindromeNumber(num) {
		fmt.Println("Число является палиндромом")
	} else {
		fmt.Println("Число не является палиндромом")
	}
}

func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}

	min := math.MaxInt
	max := math.MinInt

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func Task_9() {
	arr := []int{5, 7, -3, 8, 12, 0, -7, 4}

	min, max := findMinMax(arr)

	fmt.Println("Исходный массив:", arr)
	fmt.Printf("Минимальный элемент: %d\n", min)
	fmt.Printf("Максимальный элемент: %d\n", max)
}

func Task_10() {
	rand.Seed(time.Now().UnixNano())

	targetNumber := rand.Intn(100) + 1
	const maxAttempts = 10
	var guess int

	fmt.Println("Я загадал число от 1 до 100. Попробуйте угадать!")

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		fmt.Printf("Попытка %d/%d: Введите ваше число: ", attempts, maxAttempts)
		fmt.Scan(&guess)

		if guess < targetNumber {
			fmt.Println("Мое число больше.")
		} else if guess > targetNumber {
			fmt.Println("Мое число меньше.")
		} else {
			fmt.Printf("Поздравляю! Вы угадали число %d с %d попытки.\n", targetNumber, attempts)
			return
		}
	}

	fmt.Printf("Вы проиграли! Загаданное число было: %d\n", targetNumber)
}

func countUniqueWords(input string) int {
	words := strings.Fields(input)

	wordCount := make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	return len(wordCount)
}

func Task_12() {
	fmt.Println("Введите строку:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	uniqueCount := countUniqueWords(input)

	fmt.Printf("Количество уникальных слов: %d\n", uniqueCount)
}

const (
	rows = 10
	cols = 10
)

func displayGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("● ")
			} else {
				fmt.Print("○ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func countAliveNeighbors(grid [][]int, x, y int) int {
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	count := 0

	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]
		if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
			count += grid[newX][newY]
		}
	}

	return count
}

func updateGrid(grid [][]int) [][]int {
	newGrid := make([][]int, rows)
	for i := range newGrid {
		newGrid[i] = make([]int, cols)
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			aliveNeighbors := countAliveNeighbors(grid, x, y)

			if grid[x][y] == 1 {
				if aliveNeighbors == 2 || aliveNeighbors == 3 {
					newGrid[x][y] = 1
				} else {
					newGrid[x][y] = 0
				}
			} else {
				if aliveNeighbors == 3 {
					newGrid[x][y] = 1
				} else {
					newGrid[x][y] = 0
				}
			}
		}
	}

	return newGrid
}

func Task_13() {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	for {
		displayGrid(grid)
		grid = updateGrid(grid)
		time.Sleep(1 * time.Second)
	}
}

func digitalRoot(n int) int {
	if n < 10 {
		return n
	}

	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return digitalRoot(sum)
}

func Task_14() {
	var number int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	result := digitalRoot(number)

	fmt.Printf("Цифровой корень числа %d равен %d\n", number, result)
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; num > 0; i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}
	return roman
}

func Task_15() {
	var number int

	fmt.Print("Введите арабское число: ")
	fmt.Scan(&number)

	if number < 1 || number > 3999 {
		fmt.Println("Пожалуйста, введите число от 1 до 3999.")
		return
	}

	roman := intToRoman(number)
	fmt.Printf("Римское число: %s\n", roman)
}
