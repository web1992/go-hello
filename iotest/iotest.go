package iotest

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func IoTest() {
	fmt.Println("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"

	}
	fmt.Println("a grade of", grade, "is", status)

}

func OsStat() {
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("size is", fileInfo.Size())
}
