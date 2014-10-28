package vim25

import (
	"encoding/xml"
	"fmt"
)

// SOAP 1.1 Envelope
type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *Header  `xml:",omitempty"`
	Body    *Body
}

// SOAP 1.1 Header
type Header struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
}

// SOAP 1.1 Fault
type Fault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode"`
	String  string   `xml:"faultstring"`
	Detail  struct {
		Message string `xml:",innerxml"`
	} `xml:"detail"`
}

func (f *Fault) Error() string {
	return fmt.Sprintf("[%s] %s;%s", f.Code, f.String, f.Detail)
}

// SOAP 1.1 Body
type Body struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   *Fault

	// SearchIndex
	FindByUuidRequest  *FindByUuid
	FindByUuidResponse *FindByUuidResponse

	// ServiceInstance
	RetrieveServiceContentRequest  *RetrieveServiceContent
	RetrieveServiceContentResponse *RetrieveServiceContentResponse
	CurrentTimeRequest             *CurrentTime
	CurrentTimeResponse            *CurrentTimeResponse

	// SessionManager
	LoginRequest  *Login
	LoginResponse *LoginResponse

	// ViewManager
	CreateContainerViewRequest  *CreateContainerView
	CreateContainerViewResponse *CreateContainerViewResponse

	// PropertyCollector
	RetrievePropertiesExRequest          *RetrievePropertiesEx
	RetrievePropertiesExResponse         *RetrievePropertiesExResponse
	ContinueRetrievePropertiesExRequest  *ContinueRetrievePropertiesEx
	ContinueRetrievePropertiesExResponse *ContinueRetrievePropertiesExResponse
	CancelRetrievePropertiesExRequest    *CancelRetrievePropertiesEx
	CreatePropertyCollectorRequest       *CreatePropertyCollector
	CreatePropertyCollectorResponse      *CreatePropertyCollectorResponse
	DestroyPropertyCollectorRequest      *DestroyPropertyCollector
	DestroyPropertyCollectorResponse     *DestroyPropertyCollectorResponse

	// VirtualMachine
	PowerOnVM_Task          *PowerOnVM_Task
	PowerOnVM_TaskResponse  *PowerOnVM_TaskResponse
	PowerOffVM_Task         *PowerOffVM_Task
	PowerOffVM_TaskResponse *PowerOffVM_TaskResponse

	// OvfManager
	CreateImportSpecRequest  *CreateImportSpec
	CreateImportSpecResponse *CreateImportSpecResponse
	ImportVAppRequest        *ImportVApp
	ImportVAppResponse       *ImportVAppResponse
}
