package main

import (
	"fmt"
	"sort"
)

func SortNumbers(numbers []int) []int {
	if numbers == nil {
		return nil
	}
	sort.Ints(numbers)
	
	return numbers
}
func main() {
		fmt.Println(SortNumbers([]int{5, 3, 4, 1, 2}))
}