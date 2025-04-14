package ptz

import onvif2 "github.com/use-go/onvif/xsd/onvif"

type CreatePresetRequest struct {
	ProfileToken string `json:"profileToken"`
	PresetName   string `json:"presetName"`
}

type GetPresetsResponse struct {
	Preset []onvif2.PTZPreset
}
