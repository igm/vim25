package vim25

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.ManagedObjectReference.html
type ManagedObjectReference struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// Managed Object Reference Types
type (
	VirtualMachine    ManagedObjectReference
	SessionManager    ManagedObjectReference
	PropertyCollector ManagedObjectReference
	ServiceInstance   ManagedObjectReference
	Folder            ManagedObjectReference
	Task              ManagedObjectReference
	OptionManager     ManagedObjectReference
	ViewManager       ManagedObjectReference
	ContainerView     ManagedObjectReference
)
