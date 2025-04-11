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

func (d *OnvifDevice) GetServiceCapability() (ServiceCapabilities, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetServiceCapabilities{})

	if onvifErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Get Device Capability: %v", onvifErr)
		return ServiceCapabilities{}, onvifErr
	}

	capaBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Read Response Error: %v", readErr)
		return ServiceCapabilities{}, readErr
	}

	var deviceCapabilities DefaultResponse[ServiceCapabilitiesResponseBody]

	if unmarshalErr := xml.Unmarshal(capaBody, &deviceCapabilities); unmarshalErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Unmarshal Response Error: %v", unmarshalErr)
		return ServiceCapabilities{}, unmarshalErr
	}

	return deviceCapabilities.Body.Response.Capabilities, nil
}

func (d *OnvifDevice) GetDeviceInfo() (DeviceInformation, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetDeviceInformation{})

	if onvifErr != nil {
		log.Printf("[GET_DEVICE_INFO] Get Device Capability: %v", onvifErr)
		return DeviceInformation{}, onvifErr
	}

	capaBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_DEVICE_INFO] Read Response Error: %v", readErr)
		return DeviceInformation{}, readErr
	}

	var deviceCapabilities DefaultResponse[DeviceInformationResponseBody]

	if unmarshalErr := xml.Unmarshal(capaBody, &deviceCapabilities); unmarshalErr != nil {
		log.Printf("[GET_DEVICE_INFO] Unmarshal Response Error: %v", unmarshalErr)
		return DeviceInformation{}, unmarshalErr
	}

	return deviceCapabilities.Body.Response, nil
}

func (d *OnvifDevice) GetDeviceCapability() (DeviceCapabilitiesType, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetCapabilities{Category: "PTZ"})

	if onvifErr != nil {
		log.Printf("[GET_DEVICE_CAPA] Get Device Capability: %v", onvifErr)
		return DeviceCapabilitiesType{}, onvifErr
	}

	capaBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Read Response Error: %v", readErr)
		return DeviceCapabilitiesType{}, readErr
	}

	log.Printf("asdcasdcds: %v", string(capaBody))

	var deviceCapabilities DefaultResponse[DeviceCapabilitiesResponseBody]

	if unmarshalErr := xml.Unmarshal(capaBody, &deviceCapabilities); unmarshalErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Unmarshal Response Error: %v", unmarshalErr)
		return DeviceCapabilitiesType{}, unmarshalErr
	}

	log.Printf("[GET_SERVICE_CAPA] Response: %v", deviceCapabilities.Body.Response.Capabilities)

	return deviceCapabilities.Body.Response.Capabilities, nil
}
