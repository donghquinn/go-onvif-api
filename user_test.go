package main_test

import (
	"testing"

	"org.donghyuns.com/onvif/ptz/ptz"
)

func TestCreateUser(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	if createErr := device.CreateUser("test123", "test123", "11111"); createErr != nil {
		t.Fail()
	}
}

func TestGetUserList(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, getUserListErr := device.GetUserList()

	if getUserListErr != nil {
		t.Fail()
	}
}

func TestCreateProfile(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, createErr := device.CreateProfile("test_profile")
	if createErr != nil {
		t.Fail()
	}

	// if len(profileToken) == 0 {
	// 	t.Fail()
	// }
}

func TestGetProfile(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, getProfileErr := device.GetProfile("123456")

	if getProfileErr != nil {
		t.Fail()
	}
}
