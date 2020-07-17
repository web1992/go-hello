package filetest

import "testing"

// curl --form "upload-key=@go.mod" http://127.0.0.1:5549/upload
func TestFile(t *testing.T) {
	main()
}
