package ptz

import onvif2 "github.com/use-go/onvif/xsd/onvif"

type GetProfileRequest struct {
}

type MoveRelativeRequest struct {
	ProfileToken string  `json:"profileToken"`
	PanTiltX     float64 `json:"panTiltX"`
	PanTiltY     float64 `json:"panTiltY"`
	ZoomX        float64 `json:"zoomX"`
}

type CreateProfileRequest struct {
	ProfileName string `json:"profile"`
}

type GetStatusRequest struct {
	CctvId       string `json:"cctvId"`
	ProfileToken string `json:"profileToken"`
}

type GetConfigurationRequest struct {
	ProfileToken string `json:"profileToken"`
}

// ======= RESPONSE
type GetStatusResponse struct {
	Status  int              `json:"status"`
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Result  onvif2.PTZStatus `json:"result"`
}

type GetConfigurationResponse struct {
	Status  int                     `json:"status"`
	Code    string                  `json:"code"`
	Message string                  `json:"message"`
	Result  onvif2.PTZConfiguration `json:"result"`
}

// ======== ITEM
type GetStatusOnvifResponse struct {
	Status onvif2.PTZStatus `json:"status"`
}

type GetConfigurationOnvifResponse struct {
	Configuration onvif2.PTZConfiguration `json:"configuration"`
}

// type GetStatusOnvifResponse struct {
// 	GetStatusResponse GetStatusItem `xml:"GetStatusResponse"`
// }

// type PTZStatusItem struct {
// 	PTZStatus GetStatusItem `xml:"PTZStatus"`
// }

// type GetStatusItem struct {
// 	Position   PositionItem   `xml:"Position"`
// 	MoveStatus MoveStatusItem `xml:"MoveStatus"`
// 	UtcTime    string         `xml:"UtcTime"`
// }

// type PositionItem struct {
// 	PanTilt PanTiltItem `xml:"PanTilt"`
// 	Zoom    ZoomItem    `xml:"Zoom"`
// }

// type MoveStatusItem struct {
// 	PanTilt string `xml:"PanTilt"`
// 	Zoom    string `xml:"Zoom"`
// }

// type ZoomItem struct {
// 	X     string `xml:"x"`
// 	Space string `xml:"space"`
// }

// type PanTiltItem struct {
// 	X     string `xml:"x"`
// 	Y     string `xml:"y"`
// 	Space string `xml:"space"`
// }

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
