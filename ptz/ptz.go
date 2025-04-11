package ptz

import (
	"encoding/xml"
	"io"
	"log"

	"github.com/use-go/onvif/ptz"

	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

func (d *OnvifDevice) MoveRelative(x float64, y float64) error {
	onvifRes, onvifErr := d.CallMethod(ptz.RelativeMove{
		ProfileToken: "Profile_2",
		Translation: onvif2.PTZVector{
			PanTilt: onvif2.Vector2D{
				X:     x,
				Y:     y,
				Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/TranslationGenericSpace",
			},
			Zoom: onvif2.Vector1D{
				X:     0,
				Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/TranslationGenericSpace",
			},
		},
	})

	if onvifErr != nil {
		log.Printf("[MOVE_REL] Move Relative Error: %v", onvifErr)
		return onvifErr
	}

	ptzBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[MOVE_REL] Read Response Error: %v", readErr)
		return readErr
	}

	log.Printf("[MOVE_REL] PTZ body Response: %v", string(ptzBody))
	return nil
}

func (d *OnvifDevice) GetPresetList(profileToken string) ([]onvif2.PTZPreset, error) {
	onvifRes, onvifErr := d.CallMethod(ptz.GetPresets{
		ProfileToken: onvif2.ReferenceToken(profileToken),
	})

	if onvifErr != nil {
		log.Printf("[GET_PRESET_LIST] Call Get Preset Method Error: %v", onvifErr)
	}

	ptzBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_PRESET_LIST] Read Response Error: %v", readErr)
		return nil, readErr
	}

	var presetList GetPresetsResponse

	if unmarshalErr := xml.Unmarshal(ptzBody, &presetList); unmarshalErr != nil {
		log.Printf("[GET_PRESET_LIST] Unmarshal Preset List Response Error: %v", unmarshalErr)
	}

	log.Printf("[GET_PRESET_LIST] Get Preset List: %v", presetList)
	// presetList := GetPresetsResponse{
	// 	Preset: []onvif2.PTZPreset{},
	// }

	return presetList.Preset, nil
}
