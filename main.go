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
	for {
		fmt.Println("Введите количество элементов массива:")
		sortedArray := []int{}
		amountOfElements := 0
		fmt.Scan(&amountOfElements)
		if amountOfElements > 1 {
			array := makePseudoRandArray(amountOfElements)
			writeArrayInFile(array, "unsortedArray")
			fmt.Println("Массив создан")
			breakPoint := true  // Условие выхода из бесконечного цикла меню
			sortedFlag := false // Условие, что массив был отсортирован и можно использовать бинарный поиск
			for breakPoint == true {
				fmt.Println("Введите номер действия")
				fmt.Println("1. Пересоздать массив")
				fmt.Println("2. Отсортировать массив")
				fmt.Println("3. Произвести поиск")
				fmt.Println("0. Завершить работу")
				action := -1
				fmt.Scan(&action)
				switch action {
				case 1:
					fmt.Println("Введите количество элементов массива:")
					fmt.Scan(&amountOfElements)
					if amountOfElements > 1 {
						array = makePseudoRandArray(amountOfElements)
						writeArrayInFile(array, "unsortedArray")
						fmt.Println("Массив создан")
						sortedFlag = false
					} else {
						fmt.Println("Ошибка, введены неверные данные, повторите попытку")
					}
				case 2:
					fmt.Println("Введите номер действия")
					fmt.Println("1. Произвести сортировку выбором")
					fmt.Println("2. Произвести сортировку вставками")
					fmt.Println("3. Произвести сортировку методом пузырька")
					fmt.Println("4. Произвести быструю сортировку")
					fmt.Println("0. Вернуться")
					action := -1
					fmt.Scan(&action)
					switch action {
					case 1:
						fmt.Println("Запись начата")
						sortedArray = selectSort(array)
						writeArrayInFile(sortedArray, "sortedArray")
						sortedFlag = true
						fmt.Println("Запись окончена")
					case 2:
						fmt.Println("Запись начата")
						sortedArray = insertionSort(array)
						writeArrayInFile(sortedArray, "sortedArray")
						sortedFlag = true
						fmt.Println("Запись окончена")
					case 3:
						fmt.Println("Запись начата")
						sortedArray = bubbleSort(array)
						writeArrayInFile(sortedArray, "sortedArray")
						sortedFlag = true
						fmt.Println("Запись окончена")
					case 4:
						fmt.Println("Запись начата")
						sortedArray = QuickSort(array)
						writeArrayInFile(sortedArray, "sortedArray")
						sortedFlag = true
						fmt.Println("Запись окончена")
					case 0:
						fmt.Println("Возвращаемся в главное меню")
						break
					}
				case 3:
					if sortedFlag == true {
						fmt.Println("Введите номер действия")
						fmt.Println("1. Произвести линейный поиск в неотсортированном массиве")
						fmt.Println("2. Произвести бинарный поиск в отсортированном массиве")
						fmt.Println("0. Вернуться")
						action := -1
						fmt.Scan(&action)

						switch action {
						case 1:
							fmt.Println("Введите искомый элемент")
							searchElement := -1
							fmt.Scan(&searchElement)
							if int(searchElement) >= 0 {
								linearSearchUnsorted(array, searchElement)
							} else if int(searchElement) < 0 {
								linearSearchUnsorted(array, searchElement)
							} else {
								fmt.Println("Введено неверное значение")
							}

						case 2:
							fmt.Println("Введите искомый элемент")
							searchElement := -1
							fmt.Scan(&searchElement)
							if int(searchElement) >= 0 {
								binarySearchSorted(sortedArray, len(array), searchElement)
							} else if int(searchElement) < 0 {
								binarySearchSorted(sortedArray, len(array), searchElement)
							} else {
								fmt.Println("Введено неверное значение")
							}
						case 0:
							fmt.Println("Возвращаемся в главное меню")
							break
						}
					} else {
						fmt.Println("Введите номер действия")
						fmt.Println("Внимание! Созданный массив не был отсортирован, бинарный поиск - недоступен")
						fmt.Println("1. Произвести линейный поиск в неотсортированном массиве")
						fmt.Println("0. Вернуться")
						action := -1
						fmt.Scan(&action)

						switch action {
						case 1:
							fmt.Println("Введите искомый элемент")
							searchElement := -1
							fmt.Scan(&searchElement)
							if int(searchElement) >= 0 {
								linearSearchUnsorted(array, searchElement)
							} else if int(searchElement) < 0 {
								linearSearchUnsorted(array, searchElement)
							} else {
								fmt.Println("Введено неверное значение")
							}
						case 0:
							fmt.Println("Возвращаемся в главное меню")
							break
						}
					}
				case 0:
					fmt.Println("Завершаю работу")
					breakPoint = false
					break
				}
			}
			break

		} else {
			fmt.Println("Ошибка! Пожалуйста, введите натуральное число элементов массива больше 1")
		}

	}
	return
}
