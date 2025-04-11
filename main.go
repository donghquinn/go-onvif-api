package main

import "org.donghyuns.com/onvif/ptz/ptz"

func main() {
	device := ptz.DeviceConnect("192.168.0.152:10000")

	device.GetServiceCapability()

}
