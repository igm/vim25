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

	RetrieveServiceContentRequest  *RetrieveServiceContent
	RetrieveServiceContentResponse *RetrieveServiceContentResponse

	CurrentTimeRequest  *CurrentTime
	CurrentTimeResponse *CurrentTimeResponse

	LoginRequest  *Login
	LoginResponse *LoginResponse

	CreateContainerViewRequest  *CreateContainerView
	CreateContainerViewResponse *CreateContainerViewResponse

	RetrievePropertiesExRequest  *RetrievePropertiesEx
	RetrievePropertiesExResponse *RetrievePropertiesExResponse
}
