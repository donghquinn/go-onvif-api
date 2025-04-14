package main

import "org.donghyuns.com/onvif/ptz/ptz"

func main() {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	// device.GetServiceCapability()
	// device.GetDeviceInfo()
	// device.CreateProfile("Profile_2")
	// device.GetProfile("123456")
	// device.GetUserList()
	// device.CreateUser("dong", "dong", "1234")
	// device.GetDeviceCapability()
	// device.GetPresetList("123456")
	// device.MoveRelative("PTZNodeToken_1", -1.12391231, 10.123123)
	// device.GetStatus("123456")
	// device.GetConfiguration("123456")
	device.GetNodeList()
	// device.GetNodeInfo("PTZNodeToken_1")
}
