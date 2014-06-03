package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var VSPHERE_LOGIN string = os.Getenv("VSPHERE_LOGIN")
var VSPHERE_PASS string = os.Getenv("VSPHERE_PASS")
var VSPHERE_URL string = os.Getenv("VSPHERE_URL")

var commands = map[string]func(){}

func init() {
	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func Usage(usage func()) func() {
	return func() {
		usage()
		fmt.Println("set VSPHERE_URL, VSPHERE_LOGIN and VSPHERE_PASS as env variables. URL in the form https://<host>/sdk")
		fmt.Println("List of available commands:")
		for cmd, _ := range commands {
			fmt.Println(" -", cmd)
		}
	}
}

func main() {
	flag.Parse()
	flag.Usage = Usage(flag.Usage)
	command := flag.Arg(0)
	if cmd, exists := commands[command]; exists {
		cmd()
		return
	}
	flag.Usage()
}
