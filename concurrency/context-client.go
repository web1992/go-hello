package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	ctx := context.Background()

	req, err := http.NewRequest("GEt", "https://cn.bing.com/", nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)

	if nil != err {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

}
