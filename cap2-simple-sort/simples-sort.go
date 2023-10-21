package cap2simplesort

import "fmt"

func Sort(list []int) []int {
	//7, 2, 0, 8, 1
	sortedList := make([]int, len(list))

	// smaller := list[0]

	for i := range list {
		fmt.Println(i)
	}

	return sortedList
}

func FindSmaller(list []int) int {
		smaller := list[0]

		for _, item := range list {
			if smaller > item {
				smaller = item
			}
		}

	return smaller
}