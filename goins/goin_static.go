package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Static("/file1", "/Users/zl/Documents/DEV/github/go-hello")
	router.StaticFS("/file3", http.Dir("/Users/zl/Documents/DEV/github/go-hello"))

	router.StaticFile("/file2", "/Users/zl/Documents/DEV/github/go-hello/arraytest/arraytest.go")
	router.Run(":8082")
}
