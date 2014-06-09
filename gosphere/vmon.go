package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/igm/vim25"
)

func vmOn(c *cli.Context) {
	vmObj := c.Args().Get(0)

	sc, err := ServiceContent(service)
	if err != nil {
		log.Fatal(err)
	}
	mustLogin(service, sc.SessionManager)

	vmRef := &vim25.VirtualMachine{"VirtualMachine", vmObj}

	vmOn := &vim25.PowerOnVM_Task{This: vmRef}
	body, err := service.SoapRequest(&vim25.Body{PowerOnVM_Task: vmOn})
	if err != nil || body.Fault != nil {
		log.Fatal(err, body.Fault)
	}
	fmt.Println(body.PowerOnVM_TaskResponse.Task)
}

func vmOff(c *cli.Context) {
	vmObj := c.Args().Get(0)
	sc, err := ServiceContent(service)
	if err != nil {
		log.Fatal(err)
	}
	mustLogin(service, sc.SessionManager)

	vmRef := &vim25.VirtualMachine{"VirtualMachine", vmObj}

	vmOff := &vim25.PowerOffVM_Task{This: vmRef}
	body, err := service.SoapRequest(&vim25.Body{PowerOffVM_Task: vmOff})
	if err != nil || body.Fault != nil {
		log.Fatal(err, body.Fault)
	}
	fmt.Println(body.PowerOnVM_TaskResponse.Task)
}
