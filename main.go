package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Создает псевдослучайный массив, заданной длины amountOfElements
func makePseudoRandArray(amountOfElements int) []int {
	rand.Seed(time.Now().UnixNano())
	unsortedArray := []int{}
	for i := 0; i < amountOfElements; i++ {
		unsortedArray = append(unsortedArray, rand.Intn(1000))
	}
	return unsortedArray
}

// Записывает в файл с названием FileName, передаваемый массив данных длины amountOfElements
func writeArrayInFile(array []int, FileName string) {
	file, err := os.Create(FileName + ".txt")

	if err != nil {
		fmt.Println("Невозможно создать файл:", err)
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(array); i++ {

		file.WriteString(strconv.Itoa(array[i]))
		file.WriteString("\n")
	}
}

// Функция реализует линейный поиск элемента в любом массиве
func linearSearchUnsorted(array []int, SearchElement int) {

	for index := 0; index < len(array); index++ {
		if array[index] == SearchElement {
			fmt.Println("Искомый элемент первый раз встречается на позиции ", index+1)
			break
		} else if index == len(array)-1 {
			fmt.Println("Искомый элемент не содержится в массиве")
			break
		}
	}

}

// Функция реализует бинарный поиск элемента в отсортированном массиве
func binarySearchSorted(array []int, amountOfElements int, SearchElement int) {
	left := 0 // задаем левую и правую границы поиска
	right := amountOfElements
	search := -1 // найденный индекс элемента равен -1 (элемент не найден)
	for left <= right {
		mid := (left + right) / 2 // ищем середину отрезка

		if SearchElement == array[mid] { // если ключевое поле равно искомому
			search = mid // мы нашли требуемый элемент,
			fmt.Println("Элемент находится на этой позиции", search+1)
			break // выходим из цикла
		}

		if SearchElement < array[mid] { // если искомое ключевое поле меньше найденной середины
			right = mid - 1
		} else { // иначе
			left = mid + 1 // смещаем левую границу, продолжим поиск в правой части
		}
	}
}

// Функция реализует сортировку выбором
func selectSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		if array[min] != array[i] {
			array[min], array[i] = array[i], array[min]
		}
	}
	return array
}

// Функция реализует пузырьковую сортировку
func bubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := len(array) - 1; j > 0; j-- {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
	return array
}

// Функция реализует сортировку вставками
func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}
	}
	return array
}

// Функция реализует быструю сортировку
func Partition(array []int, l int, r int) int { // Выбирает опорным элементом послейдний элемент массива
	pivot := array[r] // ставит элементы меньше опорного слева от него а больше справа
	less := l

	for i := l; i < r; i++ {
		if array[i] <= pivot {
			array[i], array[less] = array[less], array[i]
			less++
		}
	}
	array[r], array[less] = array[less], array[r]
	return less
}

func QuickSortImpl(array []int, l int, r int) []int { // Рекурсивная функция для быстрой сортировки
	if l < r {
		q := Partition(array, l, r)
		QuickSortImpl(array, l, q-1)
		QuickSortImpl(array, q+1, r)
	}
	return array
}

func QuickSort(array []int) []int {
	if len(array) > 0 {
		QuickSortImpl(array, 0, len(array)-1)
	}
	return array
}

func main() {
	return
}
