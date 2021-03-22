package main

import (
	"context"
	"log"
	"time"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {

	select {
	case <-time.After(d):
		log.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
		//default:
		//	log.Println("default")
	}

}

func main() {

	log.Println("start")
	ctx, cancel := context.WithCancel(context.Background())

	time.AfterFunc(1*time.Second, cancel)

	sleepAndTalk(ctx, 3*time.Second, "hello")
}
