package vim25

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"testing"
)

var VSPHERE_LOGIN string = os.Getenv("VSPHERE_LOGIN")
var VSPHERE_PASS string = os.Getenv("VSPHERE_PASS")

func XXTestSimple2(t *testing.T) {
	e := &RetrieveServiceContentResponse{}
	b, err := xml.Marshal(e)
	fmt.Println(string(b), err)
}

func TestCurrentTime(t *testing.T) {
	si := &ManagedObjectReference{"ServiceInstance", "ServiceInstance"}
	service := VimService{URL: "https://127.0.0.1/sdk"}
	response := new(CurrentTimeResponse)
	err := service.Invoke(CurrentTime{This: si}, response)
	if err == nil {
		t.Error(err)
	}
}

func TestPowerOnVM(t *testing.T) {
	si := &ManagedObjectReference{"ServiceInstance", "ServiceInstance"}
	service := VimService{URL: "https://127.0.0.1/sdk"}

	response := new(RetrieveServiceContentResponse)
	err := service.Invoke(RetrieveServiceContent{This: si}, response)
	if err != nil {
		t.Error(err)
	}

	sc := response.ServiceContent

	loginResponse := new(LoginResponse)
	err = service.Invoke(Login{
		This:     sc.SessionManager,
		Username: VSPHERE_LOGIN,
		Password: VSPHERE_PASS,
	}, loginResponse)
	if err != nil {
		t.Error(err)
	}

	request := &PowerOnVM_Task{
		This: &ManagedObjectReference{"VirtualMachine", "vm-467"},
	}
	resp := new(PowerOnVm_TaskResponse)
	err = service.Invoke(request, resp)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(resp.Task)
	}
}

func TestListVMs(t *testing.T) {
	si := &ManagedObjectReference{"ServiceInstance", "ServiceInstance"}
	service := VimService{URL: "https://127.0.0.1/sdk"}

	response := new(RetrieveServiceContentResponse)
	err := service.Invoke(RetrieveServiceContent{This: si}, response)
	if err != nil {
		t.Error(err)
	}

	sc := response.ServiceContent

	loginResponse := new(LoginResponse)
	err = service.Invoke(Login{
		This:     sc.SessionManager,
		Username: VSPHERE_LOGIN,
		Password: VSPHERE_PASS,
	}, loginResponse)
	if err != nil {
		t.Error(err)
	}

	ccv := &CreateContainerView{
		This:      sc.ViewManager,
		Container: sc.RootFolder,
		Type:      []string{"VirtualMachine"},
		Recursive: true,
	}

	containerViewResponse := new(CreateContainerViewResponse)
	err = service.Invoke(ccv, containerViewResponse)
	if err != nil {
		t.Error(err)
	}
	oSpec := &ObjectSpec{
		Obj:  containerViewResponse.ContainerView,
		Skip: true,
	}

	tSpec := &TraversalSpec{
		SelectionSpec: SelectionSpec{Name: "traverseEntities"},
		Path:          "view",
		Skip:          false,
		Type:          "ContainerView",
		TypeAttr:      "TraversalSpec",
	}
	oSpec.SelectSet = append(oSpec.SelectSet, tSpec)

	pSpec := &PropertySpec{
		Type:    "VirtualMachine",
		PathSet: []string{"name"},
	}

	fSpec := &PropertyFilterSpec{
		ObjectSet: []*ObjectSpec{oSpec},
		PropSet:   []*PropertySpec{pSpec},
	}

	rpse := RetrievePropertiesEx{
		This:    sc.PropertyCollector,
		SpecSet: []*PropertyFilterSpec{fSpec},
		Options: RetrieveOptions{},
	}

	rep := new(RetrievePropertiesExResponse)
	err = service.Invoke(rpse, rep)
	if err != nil {
		t.Error(err)
	} else {
		out, _ := json.MarshalIndent(rep.RetrieveResult.Objects, "", "   ")
		fmt.Println(string(out))
	}
}
