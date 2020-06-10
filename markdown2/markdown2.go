package markdown2

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"io/ioutil"
	"log"
	"os"
)

// Markdown2Run
func Markdown2Run() {

	md := []byte("## markdown document")
	output := markdown.ToHTML(md, nil, nil)

	fmt.Println(string(output))

	b, _ := readFile("docker-compose.md")
	output2 := markdown.ToHTML(b, nil, nil)

	fmt.Println(string(output2))

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
