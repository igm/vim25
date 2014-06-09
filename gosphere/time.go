package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/igm/vim25"
)

func init() {
	app.Commands = append(app.Commands, cli.Command{
		Name:   "time",
		Usage:  "Show current time",
		Action: currentTime,
	})
}

func currentTime(c *cli.Context) {
	sc, err := ServiceContent(service)
	if err != nil {
		log.Fatal(err)
	}
	mustLogin(service, sc.SessionManager)

	body, err := service.SoapRequest(
		&vim25.Body{
			CurrentTimeRequest: &vim25.CurrentTime{This: si},
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body.CurrentTimeResponse.CurrentTime)
}
