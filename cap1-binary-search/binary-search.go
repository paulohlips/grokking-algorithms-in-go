package main

import "fmt"

func BinarySearch(list []int, x int) int{
	start := 0
	end := len(list) - 1
	center := 0
	numIterations := 0
//0, 1 ,2, 3, 4, 5, 6, 7
	for {
		center = start + calcCenter(start, end)
		fmt.Printf("start %d, end %d, center %d, iteration %d \n", start, end, center, numIterations)

		if list[center] == x {
			break
		}

		if list[center] < x {
			start = center + 1
		}

		if list[center] > x {
			end = center - 1
		}
		numIterations++

	}

	return numIterations
}

func calcCenter(start, end int) int {
	return (end - start)/2
}