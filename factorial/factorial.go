package factorial

func factorial(num int) int {

	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}

}
