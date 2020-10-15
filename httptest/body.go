package httptest

import (
	"fmt"
	log2 "log"
	"net/http"
)

func httpBody() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	err := server.ListenAndServe()
	if err != nil {
		log2.Fatal("create service error", err)
	}
	log2.Println("listen on 127.0.0.1:8080")
}

// body test
// curl -i -d "name=lognshuai&age=23" 127.0.0.1:8080/body
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
