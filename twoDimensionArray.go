package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetMaxElementRow() {
	rand.Seed(time.Now().Unix())
	arr := getRandomArray(5)

	printArray(arr)
	maxElement := arr[0][0]
	rowIndex := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > maxElement {
				maxElement = arr[i][j]
				rowIndex = i
			}
		}
	}
	fmt.Println("Строка с максимальным элементом:", arr[rowIndex])

	fmt.Println("")
}

func getRandomArray(dimension int) [][]int {
	slice := [][]int{}
	for i := 0; i < dimension; i++ {
		slice = append(slice, getRandomSlice(4))
	}

	return slice
}

func printArray(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
