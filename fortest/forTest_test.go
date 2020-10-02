package fortest

import "testing"

func TestForTest(t *testing.T) {

	ForAppend()
}

func TestForTest2(t *testing.T) {

	ForGetVar()
}

func TestForTest3(t *testing.T) {

	i := 0
	for {
		ForMap()
		if i > 1000 {
			break
		}
		i++
	}
}
