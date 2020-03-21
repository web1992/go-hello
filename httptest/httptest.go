package httptest

import (
	"fmt"
	"net/http"
)

/**
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
// MyHandler 实现 Handler 接口
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>\n")
}

// HTTPTest test
// http://127.0.0.1:8080
func HTTPTest() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler, // 以&handler作为第二个参数
	}

	server.ListenAndServe()

}

// HTTPTest2 for http.handle
// http://127.0.0.1:8080/2
func HTTPTest2() {
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handler: &handler, // 以&handler作为第二个参数
	}

	http.Handle("/2", &handler)
	server.ListenAndServe()

}

// UseFunc user func
func UseFunc() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", helloFunc)

	fmt.Println("http server listen on http://127.0.0.1:8080")
	server.ListenAndServe()
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world func")
}

// UseChain test
func UseChain() {
	myHandler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/hello", log(&myHandler))
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}

// log
// 1. print log
// 2. do handler
func log(h http.Handler) http.Handler {
	count := 0
	f := func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Printf("Handler Function called %d times\n", count)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}

// body test
// curl -id "name=lognshuai&age=23" 127.0.0.1:8080/body
func body(w http.ResponseWriter, r *http.Request) {

	len := r.ContentLength

	body := make([]byte, len)

	if readLen, ok := r.Body.Read(body); ok != nil {
		if readLen > 0 {
			fmt.Fprintf(w, "%s\n", string(body))

		} else {
			fmt.Fprintf(w, "%s\n", "# body is empty.")
		}
	} else {
		fmt.Fprintf(w, "%s\n", "service error.")
	}

}
