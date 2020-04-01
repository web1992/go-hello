package readfile

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ReadFile from file line by line
func ReadFile() {

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
