package ptz

import onvif2 "github.com/use-go/onvif/xsd/onvif"

type GetPresetsResponse struct {
	Preset []onvif2.PTZPreset
}
