package main

import "fmt"
import (
	"net/http"
	"io"
	"os/exec"
	"bytes"
	"log"
	"strings"
)

func cmdRunner(cmd string, args []string) bytes.Buffer {

	_, err := exec.LookPath(cmd)
	if err != nil{
		cmd = "ls"
		log.Print(err)
	}

	cmdExec := exec.Command(cmd)

	if cap(args) > 0{
		cmdExec = exec.Command(cmd, args...)
	}
	
	var out bytes.Buffer
	cmdExec.Stdout = &out
	cmdExec.Stderr = &out

	cmdExec.Run()

	fmt.Print(out.String())

	return out


}

func cmdHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		r.ParseMultipartForm(16*1024)//16k post data
		r.ParseForm()
		cmdData := r.Form.Get("cmd")
		argumentsData := r.Form.Get("args")

		var splited []string
		if len(argumentsData) == 0{
			splited = make([]string, 0)
		}else{
			splited = strings.Split(argumentsData, " ")
		}

		out := cmdRunner(cmdData, splited)
		fmt.Println(string(cmdData) + ":" + string(argumentsData))
		io.WriteString(w, out.String())

	}else {
		io.WriteString(w, "GET")
	}

}
func main() {
	log.Println("remote-invoker server started...")
	http.HandleFunc("/cmd/", cmdHandler)
	http.ListenAndServe(":8000", nil)
}
