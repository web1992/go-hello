package ostest

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// OsTest test
// 仅仅只读取指定的字节数
func OsTest() {

	println("ostest")

	file, err := os.Open("../go.mod") // For read access.
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}

// ReadmeRead test
func ReadmeRead() {
	fileReadme, err := os.Open("../README.md")

	if err == nil {
		defer fileReadme.Close()
		bytes, err := ioutil.ReadAll(fileReadme)
		if err == nil {
			fmt.Printf("read bytes: %q\n", bytes)
		} else {
			log.Fatal("ReadAll file error", err)
		}

	} else {
		log.Fatal("Open file error", err)
	}

}

// ReadmeFileAll test
// 使用for 循环读取文件
func ReadmeFileAll() {

	f, err := os.Open("../README.md")
	if nil != err {
		log.Fatal("Open file err is", err)
		return
	}
	defer f.Close()

	byteBuf := make([]byte, 100)
	for {
		readLen, err := f.Read(byteBuf)
		if nil == io.EOF {
			log.Fatal("End of file")
			return
		}
		if nil == err {
			fmt.Printf("read %d  bytes %q \n", readLen, byteBuf[:readLen])
		} else {
			log.Fatal("Read file err is ", err)
		}
	}

}

// ReadFileLineByLine test
// 按照行读取文件
func ReadFileLineByLine() {

	f, err := os.Open("../README.md")

	if nil != err {
		log.Fatal("Open File is ", err)
		return
	}

	defer f.Close()

	bnr := bufio.NewReader(f)

	for {
		line, err := bnr.ReadBytes('\n')
		if nil == err {
			fmt.Printf("line is %q \n", line)
			continue
		}
		if io.EOF == err {
			log.Fatal("End of file")
			break
		}

		log.Fatal("err for read file", err)
		break
	}
}
