package txttest

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Curl() {

	f, err := os.Open("1.txt")
	r, err := os.Create("r.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer r.Close()

	rd := bufio.NewReader(f)
	c := 0
	m := make(map[string]int)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		if c != 0 {
			//fmt.Println(line)
			array := strings.Split(line, "\t")

			userId := array[1]
			outId := array[7]
			d := array[8]
			if m[d] >= 10 {
				continue
			}
			m[d]++

			fmt.Fprintf(r, "http://localhost:8080/point/job/sendPoint?outId=%s&uid=%s&amount=2 \r ", outId, userId)

		}
		c++
	}

}
