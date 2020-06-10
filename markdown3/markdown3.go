package markdown3

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"io/ioutil"
	"log"
	"os"
)

func MarkDown3() {
	b, _ := readFile("docker-compose.md")

	var buf bytes.Buffer
	if err := goldmark.Convert(b, &buf); err != nil {
		panic(err)
	}

	bs := buf.Bytes()

	fmt.Printf(string(bs))
}

func readFile(fileName string) ([]byte, error) {
	md, err := os.Open(fileName)

	if err == nil {
		defer md.Close()
		bytes, err := ioutil.ReadAll(md)
		if err == nil {
			fmt.Printf("read bytes: %q\n", bytes)

			return bytes, nil
		} else {
			log.Fatal("ReadAll file error", err)
			return nil, err
		}

	} else {
		log.Fatal("Open file error", err)
	}
	return nil, err
}
