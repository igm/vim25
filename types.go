package vim25

import (
	"encoding/xml"
	"time"
)

type RetrieveServiceContent struct {
	XMLName xml.Name         `xml:"urn:vim25 RetrieveServiceContent"`
	This    *ServiceInstance `xml:"urn:vim25 _this"`
}

type RetrieveServiceContentResponse struct {
	XMLName   xml.Name        `xml:"urn:vim25 RetrieveServiceContentResponse"`
	Returnval *ServiceContent `xml:"urn:vim25 returnval"`
}

type ServiceContent struct {
	RootFolder        *Folder                 `xml:"urn:vim25 rootFolder"`
	PropertyCollector *PropertyCollector      `xml:"urn:vim25 propertyCollector"`
	ViewManager       *ManagedObjectReference `xml:"urn:vim25 viewManager"`
	About             About                   `xml:"urn:vim25 about"`
	Setting           *OptionManager          `xml:"urn:vim25 setting"`
	SessionManager    *SessionManager         `xml:"urn:vim25 sessionManager"`
}

type About struct {
	Name       string `xml:"name"`
	FullName   string `xml:"fullName"`
	Vendor     string `xml:"vendor"`
	Version    string `xml:"version"`
	Build      uint64 `xml:"build"`
	ApiVersion string `xml:"apiVersion"`
}

type CurrentTime struct {
	XMLName xml.Name        `xml:"urn:vim25 CurrentTime"`
	This    ServiceInstance `xml:"urn:vim25 _this"`
}

type CurrentTimeResponse struct {
	CurrentTime time.Time `xml:"urn:vim25 returnval"`
}

type Login struct {
	XMLName  xml.Name        `xml:"urn:vim25 Login"`
	This     *SessionManager `xml:"_this"`
	Username string          `xml:"userName"`
	Password string          `xml:"password"`
}

type LoginResponse struct {
	UserSession UserSession `xml:"urn:vim25 returnval"`
}

type UserSession struct {
	Key            string    `xml:"key"`
	UserName       string    `xml:"userName"`
	FullName       string    `xml:"fullName"`
	CallCount      uint32    `xml:"callCount"`
	LoginTime      time.Time `xml:"loginTime"`
	LastActiveTime time.Time `xml:"lastActiveTime"`
	Locate         string    `xml:"locale"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vim.view.ViewManager.html#createContainerView
type CreateContainerView struct {
	XMLName   xml.Name                `xml:"urn:vim25 CreateContainerView"`
	This      *ManagedObjectReference `xml:"_this"`
	Container *ManagedObjectReference `xml:"container"`
	Type      []string                `xml:"type"`
	Recursive bool                    `xml:"recursive"`
}

type CreateContainerViewResponse struct {
	XMLName       xml.Name       `xml:"urn:vim25 CreateContainerViewResponse"`
	ContainerView *ContainerView `xml:"urn:vim25 returnval"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.ObjectSpec.html
type ObjectSpec struct {
	Obj       *ManagedObjectReference `xml:"obj"`
	Skip      bool                    `xml:"skip"`
	SelectSet []interface{}           `xml:"selectSet"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.SelectionSpec.html
type SelectionSpec struct {
	Name string `xml:"name"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.TraversalSpec.html
type TraversalSpec struct {
	SelectionSpec
	XsiType string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Type    string `xml:"type"`
	Path    string `xml:"path"`
	Skip    bool   `xml:"skip"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.PropertySpec.html
type PropertySpec struct {
	Type    string   `xml:"type"`
	PathSet []string `xml:"pathSet"`
	All     bool     `xml:"all,omitempty"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.FilterSpec.html
type PropertyFilterSpec struct {
	PropSet                       []*PropertySpec `xml:"propSet"`
	ObjectSet                     []*ObjectSpec   `xml:"objectSet"`
	ReportMissingObjectsInResults bool            `xml:"reportMissingObjectsInResults"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvmodl.query.PropertyCollector.RetrieveOptions.html
type RetrieveOptions struct {
	MaxObjects int `xml:"maxObjects,omitempty"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.html?path=7_0_0_2_6_17_8#retrievePropertiesEx
type RetrievePropertiesEx struct {
	XMLName xml.Name              `xml:"urn:vim25 RetrievePropertiesEx"`
	This    *PropertyCollector    `xml:"_this"`
	SpecSet []*PropertyFilterSpec `xml:"specSet"`
	Options RetrieveOptions       `xml:"options"`
}

type RetrievePropertiesExResponse struct {
	XMLName        xml.Name        `xml:"urn:vim25 RetrievePropertiesExResponse"`
	RetrieveResult *RetrieveResult `xml:"returnval"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.RetrieveResult.html
type RetrieveResult struct {
	Objects []ObjectContent `xml:"objects"`
	Token   string          `xml:"token"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.ObjectContent.html
type ObjectContent struct {
	MissingSet []MissingProperty       `xml:"missingSet"`
	Obj        *ManagedObjectReference `xml:"obj"`
	PropSet    []DynamicProperty       `xml:"propSet"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvmodl.query.PropertyCollector.MissingProperty.html
type MissingProperty struct {
	// Fault LocalizedMethodFault `xml:"fault"` // TODO(igm)
	Path string `xml:"path"`
}

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
