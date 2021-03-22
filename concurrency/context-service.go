package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleFunc)
	log.Fatal(http.ListenAndServe("0.0.0.0:8082", nil))
}

func handleFunc(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log.Println("Handler started")
	defer log.Println("Handler stopped")

	select {
	case <-time.After(time.Second * 5):
		fmt.Fprintf(w, "hello")
	case <-ctx.Done():
		e := ctx.Err()
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
