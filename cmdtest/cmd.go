package cmdtest

import (
	"fmt"
	"net/http"
	"os/exec"
)

const httpPort = "0.0.0.0:9080"

/**
GOARCH=amd64 GOOS=linux go build
*/
func main() {
	CmdServer()
}

/**
127.0.0.1:9080/carAppFunc
*/
func CmdServer() {

	server := http.Server{
		Addr: httpPort,
	}
	http.HandleFunc("/carAppFunc", carAppFunc)

	fmt.Println("http server listen on " + httpPort)
	server.ListenAndServe()
}

func carAppFunc(w http.ResponseWriter, r *http.Request) {

	command := `/root/run.sh `
	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		fmt.Fprintf(w, command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	fmt.Fprintf(w, "ok")

}
