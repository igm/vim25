package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/igm/vim25"
)

func vmImport(c *cli.Context) {
	resourcePoolName := c.String("resourcepool")
	datastoreName := c.String("datastore")
	folderName := c.String("folder")
	ovfFileName := c.Args().Get(0)
	vmName := c.Args().Get(1)
	ovfFile, err := os.Open(ovfFileName)
	if err != nil {
		log.Fatal(err)
	}
	_ = ovfFile

	rp := vim25.ResourcePool{"ResourcePool", resourcePoolName}
	ds := vim25.Datastore{"Datastore", datastoreName}
	folder := vim25.Folder{"Folder", folderName}

	// real actions
	sc, err := ServiceContent(service)
	if err != nil {
		log.Fatal(err)
	}
	mustLogin(service, sc.SessionManager)

	ovfContent, err := ioutil.ReadAll(ovfFile)
	if err != nil {
		log.Fatal(err)
	}
	cis := &vim25.CreateImportSpec{
		This: sc.OvfManager,
		OvfDescriptor: vim25.OvfDescriptor{
			Value: string(ovfContent),
		},
		Datastore:    &ds,
		ResourcePool: &rp,
		Cisp: vim25.OvfCreateImportSpecParams{
			EntityName: vmName,
			OvfManagerCommonParams: vim25.OvfManagerCommonParams{
				Locale: "US",
			},
		},
	}
	body, err := service.SoapRequest(&vim25.Body{CreateImportSpecRequest: cis})
	if err != nil || body.Fault != nil {
		log.Fatal(err, body.Fault)
	}
	fmt.Println(body.CreateImportSpecResponse)
	fmt.Println(folder)

	// since we keep "innerxml", it uses "xsi" namespace prefix, so we need to replace it with what we use (or xml marshaller uses)
	// otherwise we would need to unmashall entire content and marshall it back
	replaced := strings.Replace(body.CreateImportSpecResponse.ImportSpec.Content, "xsi", "XMLSchema-instance", -1)
	body.CreateImportSpecResponse.ImportSpec.Content = replaced

	ivapp := &vim25.ImportVApp{
		This:   &rp,
		Spec:   body.CreateImportSpecResponse.ImportSpec,
		Folder: &folder,
	}
	body, err = service.SoapRequest(&vim25.Body{ImportVAppRequest: ivapp})
	if err != nil || body.Fault != nil {
		log.Fatal(err, body.Fault)
	}
	fmt.Println(body.ImportVAppResponse.HttpNfcLease)

	httpNfsLease := body.ImportVAppResponse.HttpNfcLease

	for {
		// check HttpNfcLease status
		pfs := &vim25.PropertyFilterSpec{
			PropSet: []*vim25.PropertySpec{
				{
					Type:    httpNfsLease.Type,
					All:     false,
					PathSet: []string{"state"},
				},
			},
			ObjectSet: []*vim25.ObjectSpec{
				{
					XsiType: "ObjectSpec",
					Obj:     (*vim25.ManagedObjectReference)(httpNfsLease),
					Skip:    false,
				},
			},
		}

		rpe := &vim25.RetrievePropertiesEx{
			This:    sc.PropertyCollector,
			SpecSet: []*vim25.PropertyFilterSpec{pfs},
		}

		body, err = service.SoapRequest(&vim25.Body{RetrievePropertiesExRequest: rpe})

		if err != nil || body.Fault != nil {
			log.Fatal(err, body.Fault)
		}
		fmt.Println(body.RetrievePropertiesExResponse)
	}
}
