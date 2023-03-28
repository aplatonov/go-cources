package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func SortSlice() {
	fmt.Println("---Start sorting slice---")

	rand.Seed(time.Now().Unix())
	slice := rand.Perm(20)
	randomGeneratedSlice := getRandomSlice(20)

	fmt.Println("Slice:", slice)
	fmt.Println("Random generated slice:", randomGeneratedSlice)
	fmt.Println("Bubble sorted: ", sortBubble(slice))
	fmt.Println("Standart sort: ", sortStandart(randomGeneratedSlice))

	fmt.Println("")
}

func sortBubble(slice []int) []int {
	len := len(slice)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

func sortStandart(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return slice
}

func getRandomSlice(size int) []int {
	slice := []int{}
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(100))
	}

	return slice
}
