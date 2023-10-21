package cap2simplesort

import "fmt"

func FindSmaller(list []int) []int {
		smaller := list[0]
		index := 0

		for i, item := range list {
			if smaller > item {
				smaller = item
				index = i
			}
		}

	return []int{smaller, index}
}

func Sort(list []int) []int {
	//7, 2, 0, 8, 1
	sortedList := make([]int, len(list))

	fmt.Println(sortedList)

	for i := range list {
		result := FindSmaller(list)

		smaller := result[0]
		index := result[1]

		sortedList[i] = smaller
		list = append(list[:index], list[index+1:]...)
	}

	return sortedList
}
