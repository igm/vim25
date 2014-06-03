package vim25

import (
	"encoding/xml"
	"time"
)

type ManagedObjectReference struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type (
	VirtualMachine    *ManagedObjectReference
	SessionManager    *ManagedObjectReference
	PropertyCollector *ManagedObjectReference
	ServiceInstance   *ManagedObjectReference
	Folder            *ManagedObjectReference
	Task              *ManagedObjectReference
)

type RetrieveServiceContent struct {
	XMLName xml.Name        `xml:"urn:vim25 RetrieveServiceContent"`
	This    ServiceInstance `xml:"_this"`
}

type RetrieveServiceContentResponse struct {
	ServiceContent ServiceContent `xml:"urn:vim25 returnval"`
}

type ServiceContent struct {
	RootFolder        Folder                  `xml:"rootFolder"`
	AccountManager    *ManagedObjectReference `xml:"accountManager"`
	ViewManager       *ManagedObjectReference `xml:"viewManager"`
	PropertyCollector PropertyCollector       `xml:"propertyCollector"`
	SessionManager    SessionManager          `xml:"sessionManager"`
}

type CurrentTime struct {
	XMLName xml.Name        `xml:"urn:vim25 CurrentTime"`
	This    ServiceInstance `xml:"_this"`
}

type CurrentTimeResponse struct {
	CurrentTime time.Time `xml:"urn:vim25 returnval"`
}

type Login struct {
	XMLName  xml.Name       `xml:"urn:vim25 Login"`
	This     SessionManager `xml:"_this"`
	Username string         `xml:"userName"`
	Password string         `xml:"password"`
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

type CreateContainerView struct {
	XMLName   xml.Name                `xml:"urn:vim25 CreateContainerView"`
	This      *ManagedObjectReference `xml:"_this"`
	Container *ManagedObjectReference `xml:"container"`
	Type      []string                `xml:"type"`
	Recursive bool                    `xml:"recursive"`
}

type CreateContainerViewResponse struct {
	XMLName       xml.Name                `xml:"urn:vim25 CreateContainerViewResponse"`
	ContainerView *ManagedObjectReference `xml:"urn:vim25 returnval"`
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
	TypeAttr string `xml:"xsi:type,attr"`
	Type     string `xml:"type"`
	Path     string `xml:"path"`
	Skip     bool   `xml:"skip"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.PropertySpec.html
type PropertySpec struct {
	Type    string   `xml:"type"`
	PathSet []string `xml:"pathSet"`
	All     bool     `xml:"all,omitempty"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.FilterSpec.html
type PropertyFilterSpec struct {
	PropSet   []*PropertySpec `xml:"propSet"`
	ObjectSet []*ObjectSpec   `xml:"objectSet"`
	// ReportMissingObjectsInResults bool            `xml:"reportMissingObjectsInResults"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvmodl.query.PropertyCollector.RetrieveOptions.html
type RetrieveOptions struct {
	MaxObjects int `xml:"maxObjects,omitempty"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.html?path=7_0_0_2_6_17_8#retrievePropertiesEx
type RetrievePropertiesEx struct {
	XMLName xml.Name              `xml:"urn:vim25 RetrievePropertiesEx"`
	This    PropertyCollector     `xml:"_this"`
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
	Type  string `xml:"type,attr"`
	Value string `xml:",innerxml"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html?path=7_0_0_2_6_14_8#powerOn
type PowerOnVM_Task struct {
	XMLName xml.Name       `xml:"urn:vim25 PowerOnVM_Task"`
	This    VirtualMachine `xml:"_this"`
}

type PowerOnVm_TaskResponse struct {
	Task Task `xml:"returnval"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvim.VirtualMachine.html
type PowerOffVM_Task struct {
	XMLName xml.Name       `xml:"urn:vim25 PowerOffVM_Task"`
	This    VirtualMachine `xml:"_this"`
}

type PowerOffVm_TaskResponse struct {
	Task Task `xml:"returnval"`
}
