package main

import (
	"fmt"
)

func main() {
	fmt.Println("---Start sorting slice---")
	SortSlice()

	fmt.Println("---Start two dimension array---")
	GetMaxElementRow()

	fmt.Println("---Start roman number convert---")
	s := "MMXIV"
	fmt.Println("Roman number:", s, "=", romanToInt(s))
	s = "MCXI"
	fmt.Println("Roman number:", s, "=", romanToInt(s))
	s = "MDCCCXXXIV"
	fmt.Println("Roman number:", s, "=", romanToInt(s))
	fmt.Println("")

	fmt.Println("---End program---")
}
