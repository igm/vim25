package vim25

import "encoding/xml"

type ImportVApp struct {
	XMLName xml.Name      `xml:"urn:vim25 ImportVApp"`
	This    *ResourcePool `xml:"_this"`
	Spec    ImportSpec    `xml:"spec"`
	Folder  *Folder       `xml:"folder"`
}

type ImportVAppResponse struct {
	HttpNfcLease *HttpNfcLease `xml:"returnval"`
}

type ImportSpec struct {
	XsiType string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr,ommitempty"`
	Content string `xml:",innerxml"`
}

type OvfManagerCommonParams struct {
	Locale           string      `xml:"locale"`
	DeploymentOption string      `xml:"deploymentOption"`
	MsgBundle        []*KeyValue `xml:"msgBundle,omitempty"`
	ImportOption     []string    `xml:"importOption,omitempty"`
}

type OvfCreateImportSpecParams struct {
	OvfManagerCommonParams
	DiskProvisioning   string                  `xml:"diskProvisioning,omitempty"`
	EntityName         string                  `xml:"entityName"`
	HostSystem         *ManagedObjectReference `xml:"hostSystem,omitempty"`
	InstantiationOst   *OvfConsumerOstNode     `xml:"instantiationOst,omitempty"`
	IpAllocationPolicy string                  `xml:"ipAllocationPolicy,omitempty"`
	ImProtocol         string                  `xml:"ipProtocol,omitempty"`
	NetworkMapping     []*OvfNetworkMapping    `xml:"networkMapping,omitempty"`
	PropertyMapping    []*KeyValue             `xml:"propertyMapping,omitempty"`
	ResourceMapping    []*OvfResourceMap       `xml:"resourceMapping,omitempty"`
}

// TODO
type OvfConsumerOstNode struct{}
type OvfNetworkMapping struct{}
type KeyValue struct{}
type OvfResourceMap struct{}
