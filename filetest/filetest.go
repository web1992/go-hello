package filetest

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var dir string
var port int
var staticHandler http.Handler

var (
	// 文件 key
	uploadFileKey = "upload-key"
	uploadDir     = "./down/"
)

func main() {

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", staticServer)

	if err := http.ListenAndServe(":5549", nil); err != nil {
		log.Fatalf("error to start http server:%s", err.Error())
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 接受文件
	file, header, err := r.FormFile(uploadFileKey)
	if err != nil {
		// ignore the error handler
	}
	log.Printf("selected file name is %s", header.Filename)
	// 将文件拷贝到指定路径下，或者其他文件操作
	dst, err := os.Create(uploadDir + header.Filename)
	if err != nil {
		// ignore
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		// ignore
	}
}

// 初始化参数
func init() {
	//dir = path.Dir(os.Args[0])
	//flag.Parse()
	//flag.IntVar(&port, "port", 5549, "服务器端口")
	//dir = path.Dir(uploadDir)
	fmt.Println("dir:", http.Dir(dir))
	staticHandler = http.FileServer(http.Dir(dir))
}

// 静态文件处理
func staticServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("path:" + req.URL.Path)
	if req.URL.Path != "/down/" {
		staticHandler.ServeHTTP(w, req)
		return
	}

	io.WriteString(w, "hello, world!\n")
}
