package main

import (
	"fmt"
	"time"
)

// 1. Написать программу, которая выводит текущее время и дату.
func printCurrentTime() {
	currentTime := time.Now()
	fmt.Println("Текущая дата и время:", currentTime.Format("02-01-2006 15:04:05"))
}

// 2. Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.
func printVariables() {
	var i int = 42
	var f float64 = 3.14
	var s string = "Привет, Go!"
	var b bool = true
	fmt.Println("Целое число:", i)
	fmt.Println("Число с плавающей точкой:", f)
	fmt.Println("Строка:", s)
	fmt.Println("Логическое значение:", b)
}

// 3. Использовать краткую форму объявления переменных для создания и вывода переменных.
func shortFormVariables() {
	i := 100
	f := 9.81
	s := "Краткая форма объявления"
	b := false
	fmt.Println("Целое число:", i)
	fmt.Println("Число с плавающей точкой:", f)
	fmt.Println("Строка:", s)
	fmt.Println("Логическое значение:", b)
}

// 4. Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.
func arithmeticOperations(a int, b int) {
	fmt.Println("Сложение:", a+b)
	fmt.Println("Вычитание:", a-b)
	fmt.Println("Умножение:", a*b)
	fmt.Println("Деление:", a/b)
}

// 5. Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.
func floatOperations(a float64, b float64) {
	fmt.Println("Сумма:", a+b)
	fmt.Println("Разность:", a-b)
}

// 6. Написать программу, которая вычисляет среднее значение трех чисел.
func averageOfThree(a, b, c float64) {
	average := (a + b + c) / 3
	fmt.Println("Среднее значение:", average)
}

func main() {
	// Выполнение каждой задачи

	// 1. Вывод текущей даты и времени
	printCurrentTime()

	// 2. Вывод переменных различных типов
	printVariables()

	// 3. Использование краткой формы объявления переменных
	shortFormVariables()

	// 4. Арифметические операции с двумя целыми числами
	arithmeticOperations(10, 5)

	// 5. Операции с числами с плавающей точкой
	floatOperations(10.5, 4.3)

	// 6. Среднее значение трёх чисел
	averageOfThree(3.5, 7.2, 1.8)
}
