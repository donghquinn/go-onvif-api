package ptz

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"

	"github.com/use-go/onvif"
	"github.com/use-go/onvif/device"
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

func (d *OnvifDevice) GetServiceCapability() (DeviceCapabilities, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetServiceCapabilities{})

	if onvifErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Get Device Capability: %v", onvifErr)
		return DeviceCapabilities{}, onvifErr
	}

	capaBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Read Response Error: %v", readErr)
		return DeviceCapabilities{}, readErr
	}

	var deviceCapabilities DeviceCapabilitiesResponseBody

	if unmarshalErr := xml.Unmarshal(capaBody, &deviceCapabilities); unmarshalErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Unmarshal Response Error: %v", unmarshalErr)
		return DeviceCapabilities{}, unmarshalErr
	}

	return deviceCapabilities.Response.Capabilities, nil
}

func (d *OnvifDevice) GetDeviceCapability() (string, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetCapabilities{Category: "PTZ"})

	if onvifErr != nil {
		log.Printf("[GET_DEVICE_CAPA] Get Device Capability: %v", onvifErr)
		return "", onvifErr
	}

	marshaled, marshalErr := xml.Marshal(onvifRes.Body)

	if marshalErr != nil {
		log.Printf("[GET_DEVICE_CAPA] Marshal JSON Error: %v", marshalErr)
		return "", marshalErr
	}

	log.Printf("[GET_DEVICE_CAPA] Capa Response: %v", marshaled)

	return string(marshaled), nil
}
