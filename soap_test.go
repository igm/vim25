package vim25

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

var si = ServiceInstance{"ServiceInstance", "ServiceInstance"}
var jsonEnc = json.NewEncoder(os.Stdout)

var vSphereURL = os.Getenv("VSPHERE_URL")
var vSphereLogin = os.Getenv("VSPHERE_LOGIN")
var vSpherePass = os.Getenv("VSPHERE_PASS")

func init() {
	http.DefaultClient.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	// Debug = true
}

func TestServiceContent(t *testing.T) {
	service := Service{Url: vSphereURL}
	rsc := RetrieveServiceContent{This: &si}
	var sc *ServiceContent
	if body, err := service.SoapRequest(&Body{RetrieveServiceContentRequest: &rsc}); err == nil {
		sc = body.RetrieveServiceContentResponse.Returnval
	} else {
		log.Fatal(err)
	}
	jsonEnc.Encode(sc)
}

func TestLogin(t *testing.T) {
	service := Service{Url: vSphereURL}
	rsc := RetrieveServiceContent{This: &si}
	body, _ := service.SoapRequest(&Body{RetrieveServiceContentRequest: &rsc})
	login := &Login{
		This:     body.RetrieveServiceContentResponse.Returnval.SessionManager,
		Username: vSphereLogin,
		Password: vSpherePass,
	}
	body, _ = service.SoapRequest(&Body{LoginRequest: login})
	if body.Fault != nil {
		fmt.Println(body.Fault)
	}
	jsonEnc.Encode(body.LoginResponse)
}

func TestCurrentTime(t *testing.T) {
	service := Service{Url: vSphereURL}
	rsc := RetrieveServiceContent{This: &si}
	body, _ := service.SoapRequest(&Body{RetrieveServiceContentRequest: &rsc})
	login := &Login{This: body.RetrieveServiceContentResponse.Returnval.SessionManager, Username: vSphereLogin, Password: vSpherePass}
	service.SoapRequest(&Body{LoginRequest: login})
	ct := &CurrentTime{This: si}
	body, err := service.SoapRequest(&Body{CurrentTimeRequest: ct})
	if err != nil {
		log.Fatal(err)
	}
	jsonEnc.Encode(body.CurrentTimeResponse)
}
