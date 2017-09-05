package main

import "fmt"
import (
	"net/http"
	"io"
	"os/exec"
	"bytes"
	"log"
)

func cmdRunner(cmd string, args string) bytes.Buffer {

	cmdExec := exec.Command(cmd, args)
	_, err := exec.LookPath(cmd)
	if err != nil{
		log.Print(err)
		cmdExec = exec.Command("ls", args)
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
		out := cmdRunner(cmdData, argumentsData)
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
