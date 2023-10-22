package cap4quicksort

func Biggest(list []int) int {
	biggest := list[0]

	if len(list) == 1 {
		return biggest
	}

	num := Biggest(list[1:])

	if num > biggest {
		return num
	}

	return biggest
}
