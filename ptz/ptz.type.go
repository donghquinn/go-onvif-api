package ptz

import (
	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

type GetPresetsResponse struct {
	Preset []onvif2.PTZPreset
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
