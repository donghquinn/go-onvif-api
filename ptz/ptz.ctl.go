package ptz

import (
	"net/http"

	"org.donghyuns.com/onvif/ptz/response"
	"org.donghyuns.com/onvif/ptz/utils"
)

func GetDeviceStatusCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody GetStatusRequest

	if unmarshalErr := utils.DecodeBody(req, requestBody); unmarshalErr != nil {
		response.Response(res, GetStatusResponse{
			Status:  http.StatusBadRequest,
			Code:    "STA001",
			Message: "Invalid Request",
		})
		return
	}

	device := DeviceConnect("192.168.0.152:10000")
	result := device.GetStatus(requestBody.ProfileToken)

	response.Response(res, result)
	return
}

func GetDeviceConfigCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody GetStatusRequest

	if unmarshalErr := utils.DecodeBody(req, requestBody); unmarshalErr != nil {
		response.Response(res, GetConfigurationResponse{
			Status:  http.StatusBadRequest,
			Code:    "COF001",
			Message: "Invalid Request",
		})
		return
	}

	device := DeviceConnect("192.168.0.152:10000")
	result := device.GetConfiguration(requestBody.ProfileToken)

	response.Response(res, result)
	return
}
