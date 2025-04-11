package ptz

import (
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

func (d *OnvifDevice) GetDeviceCapa() (string, error) {
	response, capErr := d.CallMethod(device.GetCapabilities{Category: "PTZ"})

	if capErr != nil {
		log.Printf("[GET_CAPA] Get Device Capability: %v", capErr)
		return "", capErr
	}

	readRes, readErr := io.ReadAll(response.Body)

	if readErr != nil {
		log.Printf("[GET_CAPA] Read Response Body Error: %v", readErr)
		return "", readErr
	}

	log.Printf("[GET_CAPA] Capa Response: %v", readRes)

	return string(readRes), nil
}
