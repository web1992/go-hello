package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result2 string
type Search2 func(keyWord string) Result2

var (
	web2   = fakeSearch2("web2")
	image2 = fakeSearch2("image2")
	video2 = fakeSearch2("video2")
)

func fakeSearch2(kind string) Search2 {
	return func(keyWord string) Result2 {
		time.Sleep(time.Duration(1000) * time.Microsecond)
		return Result2(fmt.Sprintf("%s result %s \n", kind, keyWord))
	}
}

func Google2(query string) (results []Result2) {

	ch := make(chan Result2)

	go func() {
		ch <- web2(query)
	}()

	go func() {
		ch <- image2(query)
	}()

	go func() {
		ch <- video2(query)
	}()

	for i := 0; i < 3; i++ {
		results = append(results, <-ch)
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google2("golang")
	elapsed := time.Since(start)
	fmt.Println("result:", result)
	fmt.Println("elapsed:", elapsed)
}
