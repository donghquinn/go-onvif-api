package ptz

import onvif2 "github.com/use-go/onvif/xsd/onvif"

type SetPresetRequest struct {
	ProfileToken string `json:"profileToken"`
	PresetName   string `json:"presetName"`
}

type ApplyPresetRequest struct {
	ProfileToken string  `json:"profileToken"`
	PresetToken  string  `json:"presetToken"`
	PanTiltX     float64 `json:"panTiltX"`
	PanTiltY     float64 `json:"panTiltY"`
	ZoomX        float64 `json:"zoomX"`
	IsAbsolute   bool    `json:"isAbsoulte"`
}

type GetPresetListRequest struct {
	ProfileToken string `json:"profileToken"`
}

// =========== RESPONSE
type PresetListResponse struct {
	Status  int                `json:"status"`
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Result  []onvif2.PTZPreset `json:"result"`
}

type GetPresetsResponse struct {
	Preset []onvif2.PTZPreset `json:"preset"`
}

type GetPresetListResponse struct {
	PresetList []onvif2.PTZPreset `json:"presetList"`
}
