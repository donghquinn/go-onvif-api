package main_test

import (
	"testing"

	"org.donghyuns.com/onvif/ptz/ptz"
)

func TestDeviceInfo(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, capaErr := device.GetDeviceInfo()

	if capaErr != nil {
		t.Fail()
	}
}

func TestDeviceCapa(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, capaErr := device.GetDeviceCapability()

	if capaErr != nil {
		t.Fail()
	}
}

func TestServiceCapa(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, capaErr := device.GetServiceCapability()

	if capaErr != nil {
		t.Fail()
	}
}
