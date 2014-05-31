package vim25

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Envelope struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	XmlSchema string   `xml:"xmlns:xsi,attr"`
	Header    Header
	Body      Body
}

type Header struct {
}

type Fault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode"`
	String  string   `xml:"faultstring"`
	Detail  Detail   `xml:"detail"`
}

type Detail struct {
	Message string `xml:",innerxml"`
}

func (f *Fault) Error() string { return fmt.Sprintf("[%s] %s;%s", f.Code, f.String, f.Detail) }

type Body struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Content string   `xml:",innerxml"`
	Fault   *Fault
}

type VimService struct {
	url         string
	soapSession *http.Cookie
}

func (s *VimService) invoke(request interface{}, response interface{}) error {
	requestXML, err := xml.Marshal(request)
	if err != nil {
		return err
	}
	e := &Envelope{
		XmlSchema: "http://www.w3.org/2001/XMLSchema-instance",
		Body: Body{
			Content: string(requestXML),
		},
	}
	d, err := xml.Marshal(e)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, bytes.NewReader(d))
	req.Header.Set("content-type", "text/xml; charset=\"utf-8\"")
	req.Header.Set("user-agent", "VMware VI Client/5.0.0")
	req.Header.Set("Soapaction", "\"urn:vim25/5.1\"")
	if s.soapSession != nil {
		req.AddCookie(s.soapSession)
	}
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "vmware_soap_session" {
			s.soapSession = cookie
		}
	}
	if err != nil {
		return err
	}
	// fmt.Println(resp)
	dec := xml.NewDecoder(resp.Body)
	e = new(Envelope)
	dec.Decode(e)
	// fmt.Println(e.Body.Content)
	xml.Unmarshal([]byte(e.Body.Content), response)
	if e.Body.Fault != nil {
		return e.Body.Fault
	}
	return nil
}
