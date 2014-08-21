package vim25

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

var Debug = false

type Service struct {
	Url         string
	soapSession *http.Cookie
	HttpClient  *http.Client
}

func (s *Service) readSessionCookie(resp *http.Response) {
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "vmware_soap_session" {
			s.soapSession = cookie
		}
	}
}

func (s *Service) writeSessionCookie(req *http.Request) {
	if s.soapSession != nil {
		req.AddCookie(s.soapSession)
	}
}

func (s *Service) writeHttpHeader(req *http.Request) {
	req.Header.Set("content-type", "text/xml; charset=\"utf-8\"")
	req.Header.Set("user-agent", "Vim25 GoClient/0.1")
	req.Header.Set("Soapaction", "\"urn:vim25/5.1\"")
}

func (s *Service) SoapRequest(body *Body) (*Body, error) {
	xmlEnvelope, err := xml.Marshal(Envelope{
		Body: body,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", s.Url, bytes.NewReader(xmlEnvelope))
	if err != nil {
		return nil, err
	}
	if Debug {
		dump, _ := httputil.DumpRequest(req, true)
		fmt.Println(strings.Repeat("-", 80))
		fmt.Println(strings.Repeat("-", 80))
		log.Println(string(dump))
		fmt.Println(strings.Repeat("-", 80))
	}
	s.writeHttpHeader(req)
	s.writeSessionCookie(req)
	client := http.DefaultClient
	if nil != s.HttpClient {
		client = s.HttpClient
	}
	resp, err := client.Do(req)
	if Debug {
		dump, _ := httputil.DumpResponse(resp, true)
		fmt.Println(string(dump))
		fmt.Println(strings.Repeat("-", 80))
		fmt.Println(strings.Repeat("-", 80))
	}
	if err != nil {
		return nil, err
	}
	s.readSessionCookie(resp)
	defer resp.Body.Close()
	env := new(Envelope)
	dec := xml.NewDecoder(resp.Body)
	if err := dec.Decode(env); err != nil {
		return nil, err
	}
	return env.Body, nil
}
