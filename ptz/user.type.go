package ptz

import (
	"encoding/xml"
)

type DefaultResponse[T any] struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    T        `xml:"Body"`
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

type GetProfileResponse struct {
	Profile Profile `xml:"Profile"`
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
