package cap4quicksort

func Sum(list []int) int {
	result := list[0]

	if len(list) == 1 {
		return list[0]
	}

	return result + Sum(list[1:])
}