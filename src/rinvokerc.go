package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"./utils"
)


func makePayload(cmdName string, args string, workdir string) string {
	v := url.Values{}
	v.Set("cmd", cmdName)
	v.Set("args", args)
	if len(workdir) ==0{
		v.Set("workdir", "/")
	}else{
		v.Set("workdir", workdir)
	}
	return v.Encode()
}

func main() {
	log.Println("remote-invoker client started...")
	args := os.Args

	cmdNameRAW := args[0]
	arguments := strings.Join(args[1:], " ")

	cmdNameList := strings.Split(cmdNameRAW, "/")

	cmdName:= cmdNameList[len(cmdNameList)-1]

	log.Println("loading config file for " + cmdName + " application...")

	cmdSection := utils.ConfigFileLoader(os.Getenv("HOME")+"/.rinvokerc", cmdName)

	hostIP, err := cmdSection.GetKey("hostIP")
	if err != nil {
		log.Println("Cannot found hostIP in config file")
	}
	port, err := cmdSection.GetKey("port")
	if err != nil {
		log.Println("Cannot found port in config file")
	}
	workdir, err := cmdSection.GetKey("workDir")
	if err != nil {
		log.Println("Cannot found workDir in config file")
	}

	log.Println(hostIP.Value())
	log.Println(port.Value())

	log.Println(cmdName)
	log.Println(arguments)

	req, err := http.NewRequest("POST", "http://"+hostIP.Value()+":"+port.Value()+"/cmd/", strings.NewReader(makePayload(cmdName, arguments, workdir.Value())))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")
	if err != nil {
		log.Println(err)
	}

	c := &http.Client{}

	resp, err := c.Do(req)

	if err != nil {
		fmt.Printf("http.Do() error: %v\n", err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() error: %v\n", err)
		return
	}

	fmt.Printf("read resp.Body successfully:\n%v\n", string(data))

}
