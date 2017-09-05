package main

import "fmt"
import (
	"net/http"
	"io"
)

func cmd(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		r.ParseMultipartForm(16*1024)//16k post data
		r.ParseForm()
		cmdData := r.Form.Get("cmd")
		argumentsData := r.Form.Get("args")

		io.WriteString(w, string(cmdData) + ":" + string(argumentsData))

	}else {
		io.WriteString(w, "GET")
	}

}
func main() {
	fmt.Print("hello")
	http.HandleFunc("/cmd/", cmd)
	http.ListenAndServe(":8000", nil)
}
