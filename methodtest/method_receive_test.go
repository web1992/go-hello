package methodtest

import "testing"

func TestReceive(t *testing.T) {

	r := rectangle{2, 3}

	r.area()
	r.double()
	r.area()
}
