package interfacetest

import "io"

type myWriter struct {
}

func (w myWriter) Write(p []byte) (n int, err error) {
	return
}

func TestMyStruct() {

	var _ io.Writer = (myWriter)(nil)
	var _ io.Writer = myWriter{}

}
