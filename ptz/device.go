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

func (d *OnvifDevice) GetServiceCapability() (ServiceCapaOnvifResponse, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetServiceCapabilities{})

	if onvifErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Get Device Capability: %v", onvifErr)
		return ServiceCapaOnvifResponse{}, onvifErr
	}

	capaBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Read Response Error: %v", readErr)
		return ServiceCapaOnvifResponse{}, readErr
	}

	var deviceCapabilities ServiceCapaOnvifResponse

	if unmarshalErr := xml.Unmarshal(capaBody, &deviceCapabilities); unmarshalErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Unmarshal Response Error: %v", unmarshalErr)
		return ServiceCapaOnvifResponse{}, unmarshalErr
	}

	return deviceCapabilities, nil
}

func (d *OnvifDevice) GetDeviceInfo() (DeviceInformationResponseBody, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetDeviceInformation{})

	if onvifErr != nil {
		log.Printf("[GET_DEVICE_INFO] Get Device Capability: %v", onvifErr)
		return DeviceInformationResponseBody{}, onvifErr
	}

	capaBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_DEVICE_INFO] Read Response Error: %v", readErr)
		return DeviceInformationResponseBody{}, readErr
	}

	var deviceCapabilities DefaultResponse[DeviceInformationResponseBody]

	if unmarshalErr := xml.Unmarshal(capaBody, &deviceCapabilities); unmarshalErr != nil {
		log.Printf("[GET_DEVICE_INFO] Unmarshal Response Error: %v", unmarshalErr)
		return DeviceInformationResponseBody{}, unmarshalErr
	}

	return deviceCapabilities.Body, nil
}

func (d *OnvifDevice) GetDeviceCapability() (DeviceCapaOnvifResponse, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetCapabilities{Category: "PTZ"})

	if onvifErr != nil {
		log.Printf("[GET_DEVICE_CAPA] Get Device Capability: %v", onvifErr)
		return DeviceCapaOnvifResponse{}, onvifErr
	}

	capaBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Read Response Error: %v", readErr)
		return DeviceCapaOnvifResponse{}, readErr
	}

	var deviceCapabilities DeviceCapaOnvifResponse

	if unmarshalErr := xml.Unmarshal(capaBody, &deviceCapabilities); unmarshalErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Unmarshal Response Error: %v", unmarshalErr)
		return DeviceCapaOnvifResponse{}, unmarshalErr
	}

	return deviceCapabilities, nil
}
