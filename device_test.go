package main_test

import (
	"testing"

	"org.donghyuns.com/onvif/ptz/ptz"
)

func TestDeviceCapa(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, capaErr := device.GetDeviceCapability()

	if capaErr != nil {
		t.Fatalf("Get Device Capabilities error: %v", capaErr)
	}
}

func TestServiceCapa(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, capaErr := device.GetServiceCapability()

	if capaErr != nil {
		t.Fatalf("Get Service Capabilities error: %v", capaErr)
	}
}
