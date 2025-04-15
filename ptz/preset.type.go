package ptz

type SetPresetRequest struct {
	CctvId       string `json:"cctvId"`
	ProfileToken string `json:"profileToken"`
	PresetName   string `json:"presetName"`
}

type ApplyPresetRequest struct {
	CctvId       string  `json:"cctvId"`
	ProfileToken string  `json:"profileToken"`
	PresetToken  string  `json:"presetToken"`
	PanTiltX     float64 `json:"panTiltX"`
	PanTiltY     float64 `json:"panTiltY"`
	ZoomX        float64 `json:"zoomX"`
	IsAbsolute   bool    `json:"isAbsoulte"`
}

type GetPresetListRequest struct {
	CctvId       string `json:"cctvId"`
	ProfileToken string `json:"profileToken"`
}

// =========== RESPONSE
type PresetListResponse struct {
	Status  int      `json:"status"`
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Result  []Preset `json:"result"`
}

type SetPresetResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// ======= ONVIF
type GetPresetListResponse struct {
	GetPresetsResponse GetPresetList `xml:"GetPresetsResponse"`
}

type GetPresetList struct {
	Presets []Preset `xml:"Preset"`
}

// Preset은 PTZ 프리셋을 나타냅니다
type Preset struct {
	Token       string                `xml:"token,attr"`
	Name        string                `xml:"Name"`
	PTZPosition PresetListPTZPosition `xml:"PTZPosition"`
}

type PresetListPTZPosition struct {
	PanTilt PresetListVector2D `xml:"PanTilt"`
	Zoom    PresetListVector1D `xml:"Zoom"`
}

// Vector2D는 2차원 벡터(PanTilt)를 정의합니다
type PresetListVector2D struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

// Vector1D는 1차원 벡터(Zoom)를 정의합니다
type PresetListVector1D struct {
	X float64 `xml:"x,attr"`
}

type CreatePresetOnvifResponse struct {
	SetPresetResponse CreatedPresetInfo `json:"SetPresetResponse"`
}

type CreatedPresetInfo struct {
	PresetToken string `json:"PresetToken"`
}

// type PTZPosition struct {
// 	PanTilt PanTilt `xml:"PanTilt"`
// 	Zoom    Zoom    `xml:"Zoom"`
// }

// type PanTilt struct {
// 	X     float64 `xml:"x,attr"`
// 	Y     float64 `xml:"y,attr"`
// 	Space string  `xml:"space,attr,omitempty"`
// }

// type Zoom struct {
// 	X     float64 `xml:"x,attr"`
// 	Space string  `xml:"space,attr,omitempty"`
// }
