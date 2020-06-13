package filetest

import "testing"

// curl --form "upload-key=@go.mod" http://127.0.0.1:8080/upload
func TestFile(t *testing.T) {
	Main()
}
