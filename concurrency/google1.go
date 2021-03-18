package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(keyWord string) Result

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

func fakeSearch(kind string) Search {
	return func(keyWord string) Result {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		return Result(fmt.Sprintf("%s result %s \n", kind, keyWord))
	}
}

func Google(query string) (results []Result) {
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("golang")
	elapsed := time.Since(start)
	fmt.Println("result:", result)
	fmt.Println("elapsed:", elapsed)
}
