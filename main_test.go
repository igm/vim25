package vim25

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

func init() {
	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func TestSimple(t *testing.T) {
	si := &ManagedObjectReference{"ServiceInstance", "ServiceInstance"}
	service := VimService{URL: "https://127.0.0.1/sdk"}

	response := new(RetrieveServiceContentResponse)
	err := service.Invoke(RetrieveServiceContent{This: si}, response)
	if err != nil {
		fmt.Println(".....................", err)
		t.Error(err)
	}
}
