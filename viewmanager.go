package vim25

import "encoding/xml"

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vim.view.ViewManager.html#createContainerView
type CreateContainerView struct {
	XMLName   xml.Name                `xml:"urn:vim25 CreateContainerView"`
	This      *ViewManager            `xml:"_this"`
	Container *ManagedObjectReference `xml:"container"`
	Type      []string                `xml:"type"`
	Recursive bool                    `xml:"recursive"`
}

type CreateContainerViewResponse struct {
	XMLName       xml.Name       `xml:"urn:vim25 CreateContainerViewResponse"`
	ContainerView *ContainerView `xml:"urn:vim25 returnval"`
}
