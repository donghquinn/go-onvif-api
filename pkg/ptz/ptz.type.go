package ptz

import onvif2 "github.com/use-go/onvif/xsd/onvif"

type GetProfileRequest struct {
}

type SetDefaultPositionRequest struct {
	CctvId       string `json:"cctvId"`
	ProfileToken string `json:"profileToken"`
}

type MoveToDefaultPositionRequest struct {
	CctvId       string  `json:"cctvId"`
	ProfileToken string  `json:"profileToken"`
	PanTiltX     float64 `json:"panTiltX"` // 아마 기존 위치로 돌아갈 때 x값 변화량인듯
	PanTiltY     float64 `json:"panTiltY"` // 아마 기존 위치로 돌아갈 때 y값 변화량인듯
	ZoomX        float64 `json:"zoomX"`    // 아마 기존 위치로 돌아갈 때 zoom의 x값 변화량인듯
	IsAbsolute   bool    `json:"isAbsolute"`
}

type MoveRelativeRequest struct {
	CctvId       string  `json:"cctvId"`
	ProfileToken string  `json:"profileToken"`
	PanTiltX     float64 `json:"panTiltX"`
	PanTiltY     float64 `json:"panTiltY"`
	ZoomX        float64 `json:"zoomX"`
}

type MoveContinousRequest struct {
	CctvId       string  `json:"cctvId"`
	ProfileToken string  `json:"profileToken"`
	PanTiltX     float64 `json:"panTiltX"`
	PanTiltY     float64 `json:"panTiltY"`
	ZoomX        float64 `json:"zoomX"`
	IsAbsolute   bool    `json:"isAbsolute"`
	Timeout      int     `json:"timeout"` // Second
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
	Status  int           `json:"status"`
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Result  GetStatusItem `json:"result"`
}

type GetConfigurationResponse struct {
	Status  int                     `json:"status"`
	Code    string                  `json:"code"`
	Message string                  `json:"message"`
	Result  onvif2.PTZConfiguration `json:"result"`
}

// ======== ITEM
type GetConfigurationOnvifResponse struct {
	Configuration onvif2.PTZConfiguration `json:"configuration"`
}

type GetStatusOnvifResponse struct {
	GetStatusResponse GetStatusItem `xml:"GetStatusResponse>PTZStatus"`
}

type GetStatusItem struct {
	Position   PositionItem   `xml:"Position"`
	MoveStatus MoveStatusItem `xml:"MoveStatus"`
	UtcTime    string         `xml:"UtcTime"`
}

type PositionItem struct {
	PanTilt PanTiltItem `xml:"PanTilt"`
	Zoom    ZoomItem    `xml:"Zoom"`
}

type MoveStatusItem struct {
	PanTilt string `xml:"PanTilt"`
	Zoom    string `xml:"Zoom"`
}

type ZoomItem struct {
	X     float64 `xml:"x,attr"`
	Space string  `xml:"space,attr"`
}

type PanTiltItem struct {
	X     float64 `xml:"x,attr"`
	Y     float64 `xml:"y,attr"`
	Space string  `xml:"space,attr"`
}
