package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/igm/vim25"
)

var (
	login, pass, url string
	service          *vim25.Service

	app = cli.NewApp()
	si  = &vim25.ServiceInstance{"ServiceInstance", "ServiceInstance"}
)

var flags = []cli.Flag{
	cli.StringFlag{"host", os.Getenv("VSPHERE_URL"), "remote host URL, i.e. https://<ip>/sdk"},
	cli.StringFlag{"login,l", os.Getenv("VSPHERE_LOGIN"), "login"},
	cli.StringFlag{"password,p", os.Getenv("VSPHERE_PASS"), "password"},
	cli.BoolFlag{"skip-verify,s", "skip certificate verification"},
}

func init() {
	http.DefaultClient.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	app.Version = "0.0.1"
	app.Usage = "vSphere CLI tool"
	app.Author = "Igor Mihalik"
	app.Before = before
	app.Flags = flags
}

func before(c *cli.Context) error {
	login, pass, url = c.GlobalString("login"), c.GlobalString("password"), c.GlobalString("host")
	service = &vim25.Service{Url: url}
	return nil
}

func mustLogin(s *vim25.Service, sm *vim25.SessionManager) {
	login := &vim25.Login{This: sm, Username: login, Password: pass}
	body, err := s.SoapRequest(&vim25.Body{LoginRequest: login})
	if err != nil {
		log.Fatal(err)
	}
	if body.Fault != nil {
		log.Fatal(body.Fault)
	}
}

func main() { app.Run(os.Args) }
