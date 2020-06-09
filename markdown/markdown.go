package markdown

import (
	"fmt"
	"gitlab.com/golang-commonmark/markdown"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func TestMd() string {
	md := markdown.New(markdown.XHTMLOutput(true))

	s := md.RenderToString([]byte("Header\n===\nText"))
	fmt.Println(s)
	return s
}

func Server() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", helloFunc)
	http.HandleFunc("/", homeFunc)
	server.ListenAndServe()

}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, TestMd())
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	md := markdown.New(markdown.XHTMLOutput(true))

	bytes, _ := readFile()
	s := md.RenderToString(bytes)

	fmt.Fprintf(w, s)
}

func readFile() ([]byte, error) {
	md, err := os.Open("docker-compose.md")

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
