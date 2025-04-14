package ptz

import (
	"net/http"

	"org.donghyuns.com/onvif/ptz/response"
)

// Get Service Capabilities
func GetServiceCapaCtl(res http.ResponseWriter, req *http.Request) {
	device := DeviceConnect("192.168.0.152:10000") // TODO DB 조회
	result, getErr := device.GetServiceCapability()

	if getErr != nil {
		response.Response(res, ServiceCapaResponse{
			Status: http.StatusInternalServerError,
			Code:   "SCP001",
		})

		return
	}

	response.Response(res, ServiceCapaResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  result,
	})
	return
}

// Get Device Info
func GetDeviceInfoCtl(res http.ResponseWriter, req *http.Request) {
	device := DeviceConnect("192.168.0.152:10000") // TODO DB 조회
	result, getErr := device.GetDeviceInfo()

	if getErr != nil {
		response.Response(res, DeviceInfoResponse{
			Status: http.StatusInternalServerError,
			Code:   "DVF001",
		})

		return
	}

	response.Response(res, DeviceInfoResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  result.Response,
	})
	return
}
