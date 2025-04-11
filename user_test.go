package main_test

import (
	"testing"

	"org.donghyuns.com/onvif/ptz/ptz"
)

func TestCreateUser(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	if createErr := device.CreateUser("test", "test", "123456"); createErr != nil {
		t.Fatalf("Create User Error: %v", createErr)
	}
}

func TestGetUserList(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")
	_, getUserListErr := device.GetUserList()

	if getUserListErr != nil {
		t.Fatalf("Get User List Error: %v", getUserListErr)
	}
}

func TestCreateProfile(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")
	if createErr := device.CreateProfile("test_profile", "123456"); createErr != nil {
		t.Fatalf("Create Profile Error: %v", createErr)
	}
}

func TestGetProfile(t *testing.T) {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	_, getProfileErr := device.GetProfile("123456")

	if getProfileErr != nil {
		t.Fatalf("Get Profile Error: %v", getProfileErr)
	}
}
