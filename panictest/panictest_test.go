package panictest

import "testing"

// https://blog.go-zh.org/defer-panic-and-recover
func TestPanic(t *testing.T) {
	main()
}

func TestPanic2(t *testing.T) {
	main2()
}
