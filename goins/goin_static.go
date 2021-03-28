package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	p1 = "/Users/zl/Documents/DEV/github/go-hello"
	p2 = "/Users/zl/Documents/DEV/github/go-hello"
	p3 = "/Users/zl/Documents/DEV/github/go-hello/arraytest/arraytest.go"
)

func main() {
	router := gin.Default()

	router.Static("/file1", p1)
	router.StaticFS("/file3", http.Dir(p2))

	router.StaticFile("/file2", p3)
	router.Run(":8082")
}
