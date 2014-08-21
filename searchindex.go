package vim25

import "encoding/xml"

type FindByInventoryPath struct {
	XMLName       xml.Name     `xml:"urn:vim25 FindByInventoryPath"`
	This          *SearchIndex `xml:"_this"`
	InventoryPath string       `xml:"inventoryPath"`
}

type FindByInventoryPathResponse struct {
	Folder *Folder `xml:"returnval"`
}

type FindByUuid struct {
	XMLName      xml.Name     `xml:"urn:vim25 FindByUuid"`
	This         *SearchIndex `xml:"_this"`
	Uuid         string       `xml:"uuid"`
	VmSearch     bool         `xml:"vmSearch"`
	InstanceUuid bool         `xml:"instanceUuid,omitempty"`
	//datacenter
}

type FindByUuidResponse struct {
	VmOrHost *ManagedObjectReference `xml:"returnval"`
}
