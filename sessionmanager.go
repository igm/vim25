package vim25

import (
	"encoding/xml"
	"time"
)

type Login struct {
	XMLName  xml.Name        `xml:"urn:vim25 Login"`
	This     *SessionManager `xml:"_this"`
	Username string          `xml:"userName"`
	Password string          `xml:"password"`
}

type LoginResponse struct {
	UserSession UserSession `xml:"urn:vim25 returnval"`
}

type Logout struct {
	XMLName xml.Name        `xml:"urn:vim25 Logout"`
	This    *SessionManager `xml:"_this"`
}

type LogoutResponse struct {
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
