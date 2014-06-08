package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/igm/vim25"
)

func vmOn(c *cli.Context) {
	fmt.Println(login, pass, url)
	service := &vim25.Service{Url: url}
	sc, err := ServiceContent(service)
	if err != nil {
		log.Fatal(err)
	}
	mustLogin(service, sc.SessionManager)
	fmt.Println(c.Args())
	vmObj := c.Args().Get(1)
	vmRef := &vim25.VirtualMachine{"VirtualMachine", vmObj}
	fmt.Println(vmRef)
}

// func vmPowerOn(c *cli.Context) {
// 	vLogin, vPass, vUrl := credentials(c)
//
// 	vmName := flag.Arg(1)
// 	if vmName == "" {
// 		log.Fatal("Missing Virtual Machine Obj reference value as param, i.e. vm-123")
// 	}
// 	service := vim25.VimService{URL: vURL}
// 	response := new(vim25.RetrieveServiceContentResponse)
// 	if err := service.Invoke(vim25.RetrieveServiceContent{This: vim25.ServiceInstanceRef}, response); err != nil {
// 		fmt.Println(err)
// 	}
// 	sc := response.ServiceContent
//
// 	if err := service.Login(sc.SessionManager, vLogin, vPass); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	var vmRef vim25.VirtualMachine = &vim25.ManagedObjectReference{"VirtualMachine", vmName}
// 	request := &vim25.PowerOnVM_Task{
// 		This: vmRef,
// 	}
// 	resp := new(vim25.PowerOnVm_TaskResponse)
// 	if err := service.Invoke(request, resp); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(resp.Task)
//
// }
