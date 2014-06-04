package main

import "github.com/igm/vim25"

func moList(moType string, pathSet ...string) (*vim25.RetrieveResult, error) {
	service := vim25.VimService{URL: vURL}
	response := new(vim25.RetrieveServiceContentResponse)
	if err := service.Invoke(vim25.RetrieveServiceContent{This: vim25.ServiceInstanceRef}, response); err != nil {
		return nil, err
	}
	sc := response.ServiceContent

	if err := service.Login(sc.SessionManager, vLogin, vPass); err != nil {
		return nil, err
	}

	ccv := &vim25.CreateContainerView{
		This:      sc.ViewManager,
		Container: (*vim25.ManagedObjectReference)(sc.RootFolder),
		Type:      []string{moType},
		Recursive: true,
	}

	containerViewResponse := new(vim25.CreateContainerViewResponse)
	if err := service.Invoke(ccv, containerViewResponse); err != nil {
		return nil, err
	}

	oSpec := &vim25.ObjectSpec{
		Obj:  containerViewResponse.ContainerView,
		Skip: true,
	}

	tSpec := &vim25.TraversalSpec{
		SelectionSpec: vim25.SelectionSpec{Name: "traverseEntities"},
		Path:          "view",
		Skip:          false,
		Type:          "ContainerView",
		TypeAttr:      "TraversalSpec",
	}
	oSpec.SelectSet = append(oSpec.SelectSet, tSpec)

	pSpec := &vim25.PropertySpec{
		Type:    moType,
		PathSet: pathSet,
	}

	fSpec := &vim25.PropertyFilterSpec{
		ObjectSet: []*vim25.ObjectSpec{oSpec},
		PropSet:   []*vim25.PropertySpec{pSpec},
	}

	rpse := vim25.RetrievePropertiesEx{
		This:    sc.PropertyCollector,
		SpecSet: []*vim25.PropertyFilterSpec{fSpec},
		Options: vim25.RetrieveOptions{},
	}

	rep := new(vim25.RetrievePropertiesExResponse)
	if err := service.Invoke(rpse, rep); err != nil {
		return nil, err
	}
	return rep.RetrieveResult, nil
}
