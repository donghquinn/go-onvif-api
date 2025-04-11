package ptz

import (
	"encoding/xml"
)

type DefaultResponse struct {
	XMLName xml.Name     `xml:"Envelope"`
	Body    ResponseBody `xml:"Body"`
}

type ResponseBody struct {
	GetUsersResponse GetUsersResponse `xml:"GetUsersResponse"`
}

type GetUsersResponse struct {
	User []User `xml:"User"`
}

type User struct {
	Username  string `xml:"Username"`
	UserLevel string `xml:"UserLevel"`
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
