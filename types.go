package vim25

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.DynamicProperty.html
type DynamicProperty struct {
	Name string  `xml:"name"`
	Val  AnyType `xml:"val"`
}

// xsd:anyType
type AnyType struct {
	XsiType string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Value   string `xml:",innerxml"`
}
