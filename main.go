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

func main() {
	return
}
