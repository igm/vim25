package vim25

import (
	"encoding/xml"
	"time"
)

type CurrentTime struct {
	XMLName xml.Name         `xml:"urn:vim25 CurrentTime"`
	This    *ServiceInstance `xml:"urn:vim25 _this"`
}

type CurrentTimeResponse struct {
	CurrentTime time.Time `xml:"urn:vim25 returnval"`
}

type RetrieveServiceContent struct {
	XMLName xml.Name         `xml:"urn:vim25 RetrieveServiceContent"`
	This    *ServiceInstance `xml:"urn:vim25 _this"`
}

type RetrieveServiceContentResponse struct {
	XMLName   xml.Name        `xml:"urn:vim25 RetrieveServiceContentResponse"`
	Returnval *ServiceContent `xml:"urn:vim25 returnval"`
}

type ServiceContent struct {
	RootFolder        *Folder            `xml:"urn:vim25 rootFolder"`
	PropertyCollector *PropertyCollector `xml:"urn:vim25 propertyCollector"`
	ViewManager       *ViewManager       `xml:"urn:vim25 viewManager"`
	About             About              `xml:"urn:vim25 about"`
	Setting           *OptionManager     `xml:"urn:vim25 setting"`
	SessionManager    *SessionManager    `xml:"urn:vim25 sessionManager"`
	SearchIndex       *SearchIndex       `xml:"urn:vim25 searchIndex"`
	OvfManager        *OvfManager        `xml:"ovfManager"`
}

type About struct {
	Name       string `xml:"name"`
	FullName   string `xml:"fullName"`
	Vendor     string `xml:"vendor"`
	Version    string `xml:"version"`
	Build      uint64 `xml:"build"`
	ApiVersion string `xml:"apiVersion"`
}
