package ptz

import (
	"encoding/xml"

	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

type CreateUserRequest struct {
	CctvId   string `json:"cctvId"`
	UserName string `json:"userName"`
	UserId   string `json:"userId"`
	Passwd   string `json:"passwd"`
}

type CreateProfileRequest struct {
	CctvId      string `json:"cctvId"`
	ProfileName string `json:"profileName"`
}

type DefaultResponse[T any] struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    T        `xml:"Body"`
}

// =========== RESPONSE
type CreateProfileResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type GetProfileResponse struct {
	Status  int            `json:"status"`
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Result  onvif2.Profile `json:"result"`
}

// ================ USER
type GetUserResponseBody struct {
	GetUsersResponse GetUsersResponse `xml:"GetUsersResponse"`
}

type GetUsersResponse struct {
	User []User `xml:"User"`
}

type User struct {
	Username  string `xml:"Username"`
	UserLevel string `xml:"UserLevel"`
}

// ================ PROFILE
type GetProfileResponseBody struct {
	GetProfileResponse GetProfileResponse `xml:"GetProfileResponse"`
}

type Profile struct {
	Name string `xml:"Name"`
}

// type DefaultResponse struct {
// 	Body GetUserListResponse `xml:"s:Body"`
// }

// type GetUserListResponse struct {
// 	User []UserResponseItem `xml:"tds:GetUsersResponse"`
// }

// type UserResponseItem struct {
// 	UserName  string `xml:"tt:UserName"`
// 	UserLevel string `xml:"tt:UserLevel"`
// }
