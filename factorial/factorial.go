package factorial

func factorial(num int) int {

	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}

}

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sum(arr[1:])
}
