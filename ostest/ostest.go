package ostest

import (
	"fmt"
	"log"
	"os"
)

func OsTest() {

	println("ostest")

	file, err := os.Open("go.mod") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}
