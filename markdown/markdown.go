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
	http.HandleFunc("/css.css", cssFunc)
	http.HandleFunc("/jj.css", jjFunc)
	http.HandleFunc("/", homeFunc)
	server.ListenAndServe()

}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, TestMd())
}

func cssFunc(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("content-type: text/css"))
	bytes, _ := readFile("github-markdown.css")
	fmt.Fprintf(w, string(bytes))
}

func jjFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/css")
	bytes, _ := readFile("code.css")
	fmt.Fprintf(w, string(bytes))
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	md := markdown.New(markdown.XHTMLOutput(true))

	bytes, _ := readFile("docker-compose.md")
	s := md.RenderToString(bytes)

	head := "<head>"
	headClose := "</head>"
	css := "<link type=\"text/css\" rel=\"stylesheet\" href=\"jj.css\">"
	//css2 := "<link type=\"text/css\" rel=\"stylesheet\" href=\"code2.css\">"

	page := head + css + headClose + s

	fmt.Fprintf(w, page)
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
