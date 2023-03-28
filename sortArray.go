package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sort() {
	rand.Seed(time.Now().Unix())
	slice := rand.Perm(10)

	fmt.Println("Slice:", slice)
	fmt.Println("Bubble sorted: ", sortBubble(slice))
}

func sortBubble(slice []int) []int {
	return slice
}
