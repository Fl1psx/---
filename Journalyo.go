package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Grades []int
}

var students = make(map[string]Student)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Добавить студента")
		fmt.Println("2. Показать всех")
		fmt.Println("3. Фильтр по среднему баллу")
		fmt.Println("4. Выход")
		fmt.Print("Выберите: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addStudent(scanner)
		case "2":
			showStudents()
		case "3":
			filterStudents(scanner)
		case "4":
			return
		default:
			fmt.Println("Неверный выбор!")
		}
	}
}

func addStudent(scanner *bufio.Scanner) {
	fmt.Print("ФИО: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Оценки через пробел: ")
	scanner.Scan()
	gradesInput := scanner.Text()

	grades := parseGrades(gradesInput)
	students[name] = Student{Name: name, Grades: grades}
	fmt.Printf("Студент %s добавлен\n", name)
}

func parseGrades(input string) []int {
	var grades []int
	parts := strings.Fields(input) // лучше чем Split
	for _, p := range parts {
		if g, err := strconv.Atoi(p); err == nil && g >= 1 && g <= 5 {
			grades = append(grades, g)
		}
	}
	return grades
}

func average(grades []int) float64 {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for _, g := range grades {
		sum += g
	}
	return float64(sum) / float64(len(grades))
}

func showStudents() {
	if len(students) == 0 {
		fmt.Println("Нет студентов")
		return
	}

	for _, s := range students {
		avg := average(s.Grades)
		fmt.Printf("%s: %v (ср. %.2f)\n", s.Name, s.Grades, avg)
	}
}

func filterStudents(scanner *bufio.Scanner) {
	fmt.Print("Порог среднего балла: ")
	scanner.Scan()
	thresholdStr := scanner.Text()

	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		fmt.Println("Ошибка: введите число!")
		return
	}

	found := false
	for _, s := range students {
		avg := average(s.Grades)
		if avg < threshold {
			fmt.Printf("%s: %.2f (оценки: %v)\n", s.Name, avg, s.Grades)
			found = true
		}
	}

	if !found {
		fmt.Println("Студентов с таким средним баллом нет")
	}
}
