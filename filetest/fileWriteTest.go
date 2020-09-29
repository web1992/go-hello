package filetest

import (
	"log"
	"os"
)

func writeFile() {
	file, err := os.Create("new.txt")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	b := []byte("Hello world!")

	wb, err := file.Write(b)
	_, _ = file.Write([]byte("\n"))
	_, _ = file.Write(b)

	if err != nil {
		log.Fatal(err)

	}

	log.Println("write byte is ", wb)
}
