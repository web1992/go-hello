package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result4 string
type Search4 func(query string) Result4

var (
	web4 = fakeSearch4("web4")
	web5 = fakeSearch4("web5")

	image4 = fakeSearch4("image4")
	image5 = fakeSearch4("image5")

	video4 = fakeSearch4("video4")
	video5 = fakeSearch4("video5")
)

func fakeSearch4(kind string) Search4 {
	return func(query string) Result4 {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		return Result4(fmt.Sprintf("kind %s, result %s\n", kind, query))
	}
}

func First(query string, searchArr ...Search4) Result4 {

	ch := make(chan Result4)

	for i := range searchArr {
		go func(index int) {
			ch <- searchArr[index](query)
		}(i)
	}

	return <-ch

}

func Google4(query string) []Result4 {

	var rs []Result4

	ch := make(chan Result4)

	go func() {
		ch <- First(query, web4, web5)
	}()

	go func() {
		ch <- First(query, image4, image4)
	}()

	go func() {
		ch <- First(query, video4, video4)
	}()

	t := time.After(time.Duration(3) * time.Second)
	for {
		select {

		case r := <-ch:
			rs = append(rs, r)
		case <-t:
			fmt.Println("time out")
			return rs
		}
	}
	return rs
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google4("golang")
	elapsed := time.Since(start)
	fmt.Println("result:", result)
	fmt.Println("elapsed:", elapsed)
}
