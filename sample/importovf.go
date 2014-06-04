package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/igm/vim25"
)

func init() {
	commands["import"] = importVM
}

func importVM() {
	resourcePoolName := flag.Arg(1)
	datastoreName := flag.Arg(2)
	ovfFileName := flag.Arg(3)
	ovfFile, err := os.Open(ovfFileName)
	if err != nil {
		log.Fatal(err)
	}

	rp := vim25.ResourcePool{"ResourcePool", resourcePoolName}
	ds := vim25.Datastore{"Datastore", datastoreName}

	ovfContent, err := ioutil.ReadAll(ovfFile)
	if err != nil {
		log.Fatal(err)
	}
	service := vim25.VimService{URL: vURL}
	response := new(vim25.RetrieveServiceContentResponse)
	if err := service.Invoke(vim25.RetrieveServiceContent{This: vim25.ServiceInstanceRef}, response); err != nil {
		fmt.Println(err)
	}
	if err := service.Login(response.ServiceContent.SessionManager, vLogin, vPass); err != nil {
		log.Fatal(err)
	}

	cis := vim25.CreateImportSpec{
		This: response.ServiceContent.OvfManager,
		OvfDescriptor: vim25.OvfDescriptor{
			Value: string(ovfContent),
		},
		Datastore:    &ds,
		ResourcePool: &rp,
		Cisp: vim25.OvfCreateImportSpecParams{
			EntityName: "sampleVM",
			OvfManagerCommonParams: vim25.OvfManagerCommonParams{
				Locale: "US",
			},
		},
	}

	cisr := new(vim25.CreateImportSpecResponse)
	err = service.Invoke(&cis, &cisr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cisr.ImportSpec)

}
