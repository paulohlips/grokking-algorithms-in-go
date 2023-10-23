package cap4quicksort

import "fmt"

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[len(list)/2]
	smaller := []int{}
	greater := []int{}

	for _, item := range list {
		if item < pivot {
			smaller = append(smaller, item)
		}

		if item > pivot {
			greater = append(greater, item)
		}
	}

	result := append(append(QuickSort(smaller), pivot), QuickSort(greater)...)
	fmt.Printf("%v", result)
	return result
}