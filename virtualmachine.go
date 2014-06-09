package vim25

import "encoding/xml"

type PowerOnVM_Task struct {
	XMLName xml.Name        `xml:"urn:vim25 PowerOnVM_Task"`
	This    *VirtualMachine `xml:"_this"`
}

type PowerOnVM_TaskResponse struct {
	Task *Task `xml:"returnval"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvim.VirtualMachine.html
type PowerOffVM_Task struct {
	XMLName xml.Name        `xml:"urn:vim25 PowerOffVM_Task"`
	This    *VirtualMachine `xml:"_this"`
}

type PowerOffVM_TaskResponse struct {
	Task *Task `xml:"returnval"`
}
