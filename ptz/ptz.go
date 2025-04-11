package ptz

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/use-go/onvif/ptz"
	"github.com/use-go/onvif/xsd"
	"org.donghyuns.com/onvif/ptz/util"

	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

// Move Relative
func (d *OnvifDevice) MoveRelative(profileToken string, x float64, y float64) error {
	onvifRes, onvifErr := d.CallMethod(ptz.RelativeMove{
		ProfileToken: onvif2.ReferenceToken(profileToken),
		Translation: onvif2.PTZVector{
			PanTilt: onvif2.Vector2D{
				X:     x,
				Y:     y,
				Space: xsd.AnyURI(RelativePanTiltSpace),
			},
			Zoom: onvif2.Vector1D{
				X:     0,
				Space: xsd.AnyURI(RelativeZoomSpace),
			},
		},
	})

	if onvifErr != nil {
		log.Printf("[MOVE_REL] Move Relative Error: %v", onvifErr)
		return onvifErr
	}

	if onvifRes.StatusCode != http.StatusOK {
		return fmt.Errorf("move relative response error: %v", onvifRes.StatusCode)
	}

	ptzBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[MOVE_REL] Read Response Error: %v", readErr)
		return readErr
	}

	log.Printf("[MOVE_REL] PTZ body Response: %v", string(ptzBody))
	return nil
}

// Create Preset
func (d *OnvifDevice) CreatePreset(profileToken, presetName string) error {
	presetToken := util.CreateToken()

	onvifRes, onvifErr := d.CallMethod(ptz.SetPreset{
		ProfileToken: onvif2.ReferenceToken(profileToken),
		PresetName:   xsd.String(presetName),
		PresetToken:  onvif2.ReferenceToken(presetToken),
	})

	if onvifErr != nil {
		log.Printf("[SET_PRESET] Call Get Preset Method Error: %v", onvifErr)
	}

	if onvifRes.StatusCode != http.StatusOK {
		return fmt.Errorf("set preset response error: %v", onvifRes.StatusCode)
	}

	return nil
}

// Apply Preset
func (d *OnvifDevice) ApplyPreset(
	profileToken,
	presetToken string,
	panTiltX float64,
	pantiltY float64,
	zoomX float64,
	isAbsolute bool,
) error {
	panTiltSpace := AbsolutePanTiltSpace
	zoomSpace := AbsoluteZoomSpace

	if !isAbsolute {
		panTiltSpace = RelativePanTiltSpace
		zoomSpace = RelativeZoomSpace
	}

	panTilt := onvif2.Vector2D{
		X:     panTiltX,
		Y:     pantiltY,
		Space: xsd.AnyURI(panTiltSpace),
	}

	zoom := onvif2.Vector1D{
		X:     zoomX,
		Space: xsd.AnyURI(zoomSpace),
	}

	onvifRes, onvifErr := d.CallMethod(ptz.GotoPreset{
		ProfileToken: onvif2.ReferenceToken(profileToken),
		PresetToken:  onvif2.ReferenceToken(presetToken),
		Speed:        onvif2.PTZSpeed{PanTilt: panTilt, Zoom: zoom},
	})

	if onvifErr != nil {
		log.Printf("[SET_PRESET] Call Get Preset Method Error: %v", onvifErr)
	}

	if onvifRes.StatusCode != http.StatusOK {
		return fmt.Errorf("set preset response error: %v", onvifRes.StatusCode)
	}

	return nil
}

// Get Preset List
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

	return presetList.Preset, nil
}
