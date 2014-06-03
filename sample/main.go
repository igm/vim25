package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var vLogin, vPass, vURL string

var commands = map[string]func(){}

func init() {
	flag.StringVar(&vLogin, "u", os.Getenv("VSPHERE_LOGIN"), "vSphere username, default VSPHERE_LOGIN env variable")
	flag.StringVar(&vPass, "p", os.Getenv("VSPHERE_PASS"), "vSphere password, default VSPHERE_PASS env variable")
	flag.StringVar(&vURL, "r", os.Getenv("VSPHERE_URL"), "vSphere URL, default VSPHERE_URL env variable")

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
