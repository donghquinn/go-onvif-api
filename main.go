package main

import "org.donghyuns.com/onvif/ptz/ptz"

func main() {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	// device.CreateProfile("Profile_2", "123456")
	// device.GetProfile("123456")
	// device.GetUserList()
	// device.CreateUser("dong", "dong", "1234")
	// device.GetDeviceCapability()
	device.GetPresetList("123456")
}
