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

func main() {
	return
}
