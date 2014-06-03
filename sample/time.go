package main

import (
	"fmt"
	"log"

	"github.com/igm/vim25"
)

func init() {
	commands["time"] = currentTime
}

func currentTime() {
	service := vim25.VimService{URL: vURL}
	response := new(vim25.RetrieveServiceContentResponse)
	if err := service.Invoke(vim25.RetrieveServiceContent{This: vim25.ServiceInstanceRef}, response); err != nil {
		fmt.Println(err)
	}
	if err := service.Login(response.ServiceContent.SessionManager, vLogin, vPass); err != nil {
		log.Fatal(err)
	}
	resTime := new(vim25.CurrentTimeResponse)
	if err := service.Invoke(vim25.CurrentTime{This: vim25.ServiceInstanceRef}, resTime); err != nil {
		log.Fatal(err)
	}
	fmt.Println(resTime.CurrentTime)
}
