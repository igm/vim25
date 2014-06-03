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
	flag.StringVar(&VSPHERE_LOGIN, "u", os.Getenv("VSPHERE_LOGIN"), "vSphere username, default VSPHERE_LOGIN env variable")
	flag.StringVar(&VSPHERE_PASS, "p", os.Getenv("VSPHERE_PASS"), "vSphere password, default VSPHERE_PASS env variable")
	flag.StringVar(&VSPHERE_URL, "r", os.Getenv("VSPHERE_URL"), "vSphere URL, default VSPHERE_URL env variable")

	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func Usage(usage func()) func() {
	return func() {
		usage()
		fmt.Println("\nList of available commands:")
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
