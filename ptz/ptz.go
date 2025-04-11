package ptz

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"

	"github.com/use-go/onvif"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/ptz"

	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

type OnvifDevice struct {
	*onvif.Device
}

func DeviceConnect(endpoint string) *OnvifDevice {
	httpClient := http.Client{}

	deviceConfig := onvif.DeviceParams{
		Xaddr:      endpoint,
		Username:   "user",
		Password:   "123456",
		HttpClient: &httpClient,
	}

	device, deviceErr := onvif.NewDevice(deviceConfig)

	if deviceErr != nil {
		log.Printf("[NEW_DEVICE] Connect New Device: %v", deviceErr)
		return nil
	}

	return &OnvifDevice{
		device,
	}
}

func (d *OnvifDevice) GetServiceCapability() (string, error) {
	response, capErr := d.CallMethod(device.GetServiceCapabilities{})

	if capErr != nil {
		log.Printf("[GET_SRV_CAPA] Get Device Capability: %v", capErr)
		return "", capErr
	}

	marshaled, marshalErr := json.Marshal(response.Body)

	if marshalErr != nil {
		log.Printf("[GET_SRV_CAPA] Marshal JSON Error: %v", marshalErr)
		return "", marshalErr
	}

	log.Printf("[GET_SRV_CAPA] Capa Response: %v", marshaled)

	return string(marshaled), nil
}

func (d *OnvifDevice) GetDeviceCapa() (string, error) {
	response, capErr := d.CallMethod(device.GetCapabilities{Category: "PTZ"})

	if capErr != nil {
		log.Printf("[GET_DEVICE_CAPA] Get Device Capability: %v", capErr)
		return "", capErr
	}

	marshaled, marshalErr := json.Marshal(response.Body)

	if marshalErr != nil {
		log.Printf("[GET_DEVICE_CAPA] Marshal JSON Error: %v", marshalErr)
		return "", marshalErr
	}

	log.Printf("[GET_DEVICE_CAPA] Capa Response: %v", marshaled)

	return string(marshaled), nil
}

func (d *OnvifDevice) MoveRelative(x float64, y float64) error {
	ptzRes, ptzErr := d.CallMethod(ptz.RelativeMove{
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

	if ptzErr != nil {
		log.Printf("[MOVE_REL] Move Relative Error: %v", ptzErr)
		return ptzErr
	}

	ptzBody, readErr := io.ReadAll(ptzRes.Body)

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
