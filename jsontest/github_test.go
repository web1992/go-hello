package jsontest

import (
	"fmt"
	"log"
	"testing"
)

func TestSearchIssues(t *testing.T) {
	s := []string{"json", "decoder"}
	r, err := SearchIssues(s)

	if nil != err {
		log.Fatal(err)
	}
	fmt.Printf("%d \n", r.TotalCount)

	for _, item := range r.Items {
		fmt.Printf("#%-5d %9.9s %.55s \n", item.Number, item.User.Login, item.Title)
	}
}
