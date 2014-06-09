package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/igm/vim25"
)

func ServiceContent(service *vim25.Service) (*vim25.ServiceContent, error) {
	rsc := vim25.RetrieveServiceContent{This: si}
	body, err := service.SoapRequest(&vim25.Body{RetrieveServiceContentRequest: &rsc})
	if err != nil {
		return nil, err
	}
	return body.RetrieveServiceContentResponse.Returnval, nil
}

func vmList(c *cli.Context) {
	sc, err := ServiceContent(service)
	if err != nil {
		log.Fatal(err)
	}
	mustLogin(service, sc.SessionManager)

	ccv := &vim25.CreateContainerView{
		This:      sc.ViewManager,
		Container: (*vim25.ManagedObjectReference)(sc.RootFolder),
		Type:      []string{"VirtualMachine"},
		Recursive: true,
	}

	body, err := service.SoapRequest(&vim25.Body{CreateContainerViewRequest: ccv})
	if err != nil {
		log.Fatal(err)
	}
	cv := body.CreateContainerViewResponse.ContainerView

	oSpec := &vim25.ObjectSpec{
		XsiType: "ObjectSpec",
		Obj:     (*vim25.ManagedObjectReference)(cv),
		Skip:    true,
	}

	tSpec := &vim25.TraversalSpec{
		SelectionSpec: vim25.SelectionSpec{Name: "traverseEntities"},
		XsiType:       "TraversalSpec",
		Path:          "view",
		Skip:          false,
		Type:          "ContainerView",
	}
	oSpec.SelectSet = append(oSpec.SelectSet, tSpec)

	pSpec := &vim25.PropertySpec{
		Type:    "VirtualMachine",
		PathSet: []string{"name"},
	}

	fSpec := &vim25.PropertyFilterSpec{
		ObjectSet: []*vim25.ObjectSpec{oSpec},
		PropSet:   []*vim25.PropertySpec{pSpec},
	}

	rpse := &vim25.RetrievePropertiesEx{
		This:    sc.PropertyCollector,
		SpecSet: []*vim25.PropertyFilterSpec{fSpec},
		Options: vim25.RetrieveOptions{},
	}

	body, err = service.SoapRequest(&vim25.Body{RetrievePropertiesExRequest: rpse})
	if err != nil || body.Fault != nil {
		log.Fatal(err, body.Fault)
	}

	rep := body.RetrievePropertiesExResponse
	for _, rep := range rep.RetrieveResult.Objects {
		fmt.Printf("Managed Object Reference: (%s,%s)\n", rep.Obj.Type, rep.Obj.Value)
		for _, prop := range rep.PropSet {
			fmt.Printf("\t %s:[type:'%s', value:'%s']\n", prop.Name, prop.Val.XsiType, prop.Val.Value)
		}
	}
}
