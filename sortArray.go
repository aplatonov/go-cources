package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const size = 20

func SortSlice() {
	fmt.Println("---Start sorting slice---")

	rand.Seed(time.Now().Unix())
	slice := rand.Perm(size)

	fmt.Println("Slice:", slice)
	fmt.Println("Bubble sorted: ", sortBubble(slice))
	fmt.Println("Standart sort: ", sortStandart(slice))
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
