package vim25

import "encoding/xml"

// http://pubs.vmware.com/vsphere-55/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvim.OvfManager.html&resultof=%22CreateImportSpec%22%20%22createimportspec%22%20
type CreateImportSpec struct {
	XMLName       xml.Name                  `xml:"urn:vim25 CreateImportSpec"`
	This          *OvfManager               `xml:"_this"`
	OvfDescriptor OvfDescriptor             `xml:"ovfDescriptor"`
	ResourcePool  *ResourcePool             `xml:"resourcePool"`
	Datastore     *Datastore                `xml:"datastore"`
	Cisp          OvfCreateImportSpecParams `xml:"cisp"`
}

type CreateImportSpecResponse struct {
	ImportSpec ImportSpec             `xml:"returnval>importSpec"`
	FileItem   []OvfFileItem          `xml:"returnval>fileItem"`
	Warning    []LocalizedMethodFault `xml:"warning"`
	Error      []LocalizedMethodFault `xml:"error"`
}

type OvfDescriptor struct {
	Value string `xml:",chardata"`
}

type LocalizedMethodFault struct {
	Fault            MethodFault `xml:"fault"`
	LocalizedMessage string      `xml:"localizedMessage"`
}

type MethodFault struct {
	DynamicType     string            `xml:"dynamicType"`
	DynamicProperty []DynamicProperty `xml:"dynamicProperty"`
}

type OvfFileItem struct {
	DeviceID          string `xml:"deviceId"`
	Path              string `xml:"path"`
	CompressionMethod string `xml:"compressionMethod"`
	ChunkSize         int64  `xml:"chunkSize"`
	Size              int64  `xml:"size"`
	CimType           int32  `xml:"cimType"`
	Create            bool   `xml:"create"`
}
