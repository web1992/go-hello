package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result3 string
type Search3 func(query string) Result3

var (
	web3   = fakeSearch3("web3")
	image3 = fakeSearch3("image3")
	video3 = fakeSearch3("video3")
)

func fakeSearch3(kind string) Search3 {
	return func(query string) Result3 {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		return Result3(fmt.Sprintf("kind %s, result %s\n", kind, query))
	}
}

func Google3(query string) (results []Result3) {

	ch := make(chan Result3)

	go func() {
		ch <- web3(query)
	}()

	go func() {
		ch <- image3(query)
	}()

	go func() {
		ch <- video3(query)
	}()

	t := time.After(time.Duration(5) * time.Second)
	for {
		select {
		case r := <-ch:
			results = append(results, r)
		case tt := <-t:
			fmt.Println("time out after ", tt, " second")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google3("golang")
	elapsed := time.Since(start)
	fmt.Println("result:", result)
	fmt.Println("elapsed:", elapsed)
}
