package ptz

import (
	"encoding/xml"
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
	Status  int     `json:"status"`
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Result  Profile `json:"result"`
}

type GetProfileListResponse struct {
	Status  int       `json:"status"`
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Result  []Profile `json:"result"`
}

type GetUserListResponse struct {
	Status  int                `json:"status"`
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Result  []UserResponseItem `json:"result"`
}

// ================ USER

type GetUserOnvifListResponse struct {
	User []UserResponseItem `xml:"GetUsersResponse>User"`
}

type UserResponseItem struct {
	UserName  string `xml:"Username"`
	UserLevel string `xml:"UserLevel"`
}

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
	GetProfileResponse Profile `xml:"GetProfileResponse>Profile"`
}

type Profile struct {
	Token                       string                      `xml:"token,attr"`
	Fixed                       bool                        `xml:"fixed,attr"`
	Name                        string                      `xml:"Name"`
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration `xml:"VideoAnalyticsConfiguration"`
	PTZConfiguration            PTZConfiguration            `xml:"PTZConfiguration"`
	MetadataConfiguration       MetadataConfiguration       `xml:"MetadataConfiguration"`
}

type VideoAnalyticsConfiguration struct {
	Token                        string                       `xml:"token,attr"`
	Name                         string                       `xml:"Name"`
	UseCount                     int                          `xml:"UseCount"`
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `xml:"AnalyticsEngineConfiguration"`
	RuleEngineConfiguration      RuleEngineConfiguration      `xml:"RuleEngineConfiguration"`
}

type AnalyticsEngineConfiguration struct {
	Modules []AnalyticsModule `xml:"AnalyticsModule"`
}

type AnalyticsModule struct {
	Name       string     `xml:"Name,attr"`
	Type       string     `xml:"Type,attr"`
	Parameters Parameters `xml:"Parameters"`
}

type RuleEngineConfiguration struct {
	Rules []Rule `xml:"Rule"`
}

type Rule struct {
	Name       string     `xml:"Name,attr"`
	Type       string     `xml:"Type,attr"`
	Parameters Parameters `xml:"Parameters"`
}

type Parameters struct {
	SimpleItems  []SimpleItem  `xml:"SimpleItem"`
	ElementItems []ElementItem `xml:"ElementItem"`
}

type SimpleItem struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

type ElementItem struct {
	Name       string      `xml:"Name,attr"`
	CellLayout *CellLayout `xml:"CellLayout"`
}

type CellLayout struct {
	Columns        int            `xml:"Columns,attr"`
	Rows           int            `xml:"Rows,attr"`
	Transformation Transformation `xml:"Transformation"`
}

type Transformation struct {
	Translate Translate `xml:"Translate"`
	Scale     Scale     `xml:"Scale"`
}

type Translate struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type Scale struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type PTZConfiguration struct {
	Token                                  string        `xml:"token,attr"`
	MoveRamp                               int           `xml:"MoveRamp,attr"`
	PresetRamp                             int           `xml:"PresetRamp,attr"`
	PresetTourRamp                         int           `xml:"PresetTourRamp,attr"`
	Name                                   string        `xml:"Name"`
	UseCount                               int           `xml:"UseCount"`
	NodeToken                              string        `xml:"NodeToken"`
	DefaultAbsolutePantTiltPositionSpace   string        `xml:"DefaultAbsolutePantTiltPositionSpace"`
	DefaultAbsoluteZoomPositionSpace       string        `xml:"DefaultAbsoluteZoomPositionSpace"`
	DefaultRelativePanTiltTranslationSpace string        `xml:"DefaultRelativePanTiltTranslationSpace"`
	DefaultRelativeZoomTranslationSpace    string        `xml:"DefaultRelativeZoomTranslationSpace"`
	DefaultContinuousPanTiltVelocitySpace  string        `xml:"DefaultContinuousPanTiltVelocitySpace"`
	DefaultContinuousZoomVelocitySpace     string        `xml:"DefaultContinuousZoomVelocitySpace"`
	DefaultPTZSpeed                        PTZSpeed      `xml:"DefaultPTZSpeed"`
	DefaultPTZTimeout                      string        `xml:"DefaultPTZTimeout"`
	PanTiltLimits                          PanTiltLimits `xml:"PanTiltLimits"`
	ZoomLimits                             ZoomLimits    `xml:"ZoomLimits"`
	Extension                              PTZExtension  `xml:"Extension"`
}

type PTZSpeed struct {
	PanTilt PanTilt `xml:"PanTilt"`
	Zoom    Zoom    `xml:"Zoom"`
}

type PanTilt struct {
	X     float64 `xml:"x,attr"`
	Y     float64 `xml:"y,attr"`
	Space string  `xml:"space,attr"`
}

type Zoom struct {
	X     float64 `xml:"x,attr"`
	Space string  `xml:"space,attr"`
}

type PanTiltLimits struct {
	Range PTZRange `xml:"Range"`
}

type ZoomLimits struct {
	Range PTZRange `xml:"Range"`
}

type PTZRange struct {
	URI    string  `xml:"URI"`
	XRange MinMax  `xml:"XRange"`
	YRange *MinMax `xml:"YRange"` // Zoom에는 YRange 없음
}

type MinMax struct {
	Min float64 `xml:"Min"`
	Max float64 `xml:"Max"`
}

type PTZExtension struct {
	PTControlDirection PTControlDirection `xml:"PTControlDirection"`
}

type PTControlDirection struct {
	EFlip   FlipMode `xml:"EFlip"`
	Reverse FlipMode `xml:"Reverse"`
}

type FlipMode struct {
	Mode string `xml:"Mode"`
}

type MetadataConfiguration struct {
	Token           string         `xml:"token,attr"`
	CompressionType string         `xml:"CompressionType"`
	GeoLocation     string         `xml:"GeoLocation"`
	ShapePolygon    string         `xml:"ShapePolygon"`
	Name            string         `xml:"Name"`
	UseCount        int            `xml:"UseCount"`
	PTZStatus       PTZStatusFlags `xml:"PTZStatus"`
	Analytics       string         `xml:"Analytics"`
	Multicast       Multicast      `xml:"Multicast"`
	SessionTimeout  string         `xml:"SessionTimeout"`
}

type PTZStatusFlags struct {
	Status   string `xml:"Status"`
	Position string `xml:"Position"`
}

type Multicast struct {
	Address   Address `xml:"Address"`
	Port      int     `xml:"Port"`
	TTL       int     `xml:"TTL"`
	AutoStart bool    `xml:"AutoStart"`
}

type Address struct {
	Type        string `xml:"Type"`
	IPv4Address string `xml:"IPv4Address"`
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
