package ptz

import (
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

	log.Printf("asdcasdcsa: %v", string(capaBody))
	var deviceCapabilities DefaultResponse[ServiceCapabilitiesResponseBody]

	if unmarshalErr := xml.Unmarshal(capaBody, &deviceCapabilities); unmarshalErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Unmarshal Response Error: %v", unmarshalErr)
		return ServiceCapabilities{}, unmarshalErr
	}

	return deviceCapabilities.Body.Response.Capabilities, nil
}

// 디바이스 정보 조회
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

// 디바이스 캐파 조회
func (d *OnvifDevice) GetDeviceCapability() (ServiceCapabilities, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetCapabilities{Category: "PTZ"})

	if onvifErr != nil {
		log.Printf("[GET_DEVICE_CAPA] Get Device Capability: %v", onvifErr)
		return ServiceCapabilities{}, onvifErr
	}

	capaBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Read Response Error: %v", readErr)
		return ServiceCapabilities{}, readErr
	}

	var deviceCapabilities DefaultResponse[ServiceCapaOnvifResponse]

	if unmarshalErr := xml.Unmarshal(capaBody, &deviceCapabilities); unmarshalErr != nil {
		log.Printf("[GET_SERVICE_CAPA] Unmarshal Response Error: %v", unmarshalErr)
		return ServiceCapabilities{}, unmarshalErr
	}

	return deviceCapabilities.Body.Capabilities.Capabilities, nil
}

// 노드 상태 조회
func (d *OnvifDevice) GetStatus(profileToken string) GetStatusResponse {
	onvifRes, onvifErr := d.CallMethod(ptz.GetStatus{
		ProfileToken: onvif2.ReferenceToken(profileToken),
	})

	if onvifErr != nil {
		log.Printf("[GET_STATUS] Call Get Preset Method Error: %v", onvifErr)
		return GetStatusResponse{
			Status:  http.StatusInternalServerError,
			Code:    "STA001",
			Message: "Request Status ONVIF Error",
		}
	}

	if onvifRes.StatusCode != http.StatusOK {
		return GetStatusResponse{
			Status:  http.StatusInternalServerError,
			Code:    "STA001",
			Message: "Request Status ONVIF Error",
		}
	}

	ptzBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_STATUS] Read Response Error: %v", readErr)
		return GetStatusResponse{
			Status:  http.StatusInternalServerError,
			Code:    "STA002",
			Message: "Read ONVIF Response Error",
		}
	}

	var onvifStatusReponse DefaultResponse[GetStatusOnvifResponse]

	if unmarshal := xml.Unmarshal(ptzBody, &onvifStatusReponse); unmarshal != nil {
		log.Printf("[GET_STATUS] Unmarshal ONVIF Response Error: %v", unmarshal)
		return GetStatusResponse{
			Status:  http.StatusInternalServerError,
			Code:    "STA003",
			Message: "Read ONVIF Response Error",
		}
	}

	return GetStatusResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  onvifStatusReponse.Body.GetStatusResponse,
	}
}
