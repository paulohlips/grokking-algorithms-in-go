package cap3recursion

func Factorial(num int) int {
	if num == 1 {
		return num
	}

	return num * Factorial(num-1)
}
