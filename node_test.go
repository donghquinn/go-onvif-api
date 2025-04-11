package main_test

import (
	"testing"

	"org.donghyuns.com/onvif/ptz/ptz"
)

func TestGetNodeList(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, getUserListErr := device.GetNodeList()

	if getUserListErr != nil {
		t.Fail()
	}
}

func TestGetNodeInfo(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, getUserListErr := device.GetNodeInfo("PTZNodeToken_1")

	if getUserListErr != nil {
		t.Fail()
	}
}
