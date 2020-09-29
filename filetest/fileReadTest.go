package filetest

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFile() {

	b, err := ioutil.ReadFile("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	txt := string(b)

	fmt.Println(txt)

}
