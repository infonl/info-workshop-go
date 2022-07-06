package main

import (
	"fmt"
	bubbleSort "info-workshop-go/bubble-sort"
)

func main() {

	slice := bubbleSort.GenerateSlice(10)

	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubbleSort.BubbleSort(slice)

	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

}
