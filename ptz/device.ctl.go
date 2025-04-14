package ptz

import (
	"net/http"

	"org.donghyuns.com/onvif/ptz/database"
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

// Get Status
func GetDeviceStatusCtl(res http.ResponseWriter, req *http.Request) {
	cctvId := req.URL.Query().Get("cctv")
	profileToken := req.URL.Query().Get("profile")

	if cctvId == "" || profileToken == "" {
		response.Response(res, GetProfileResponse{
			Status:  http.StatusBadRequest,
			Code:    "STA001",
			Message: "Invalid Params",
		})

		return
	}

	endpoint, getErr := database.GetDeviceInfo(cctvId)
	if getErr != nil {
		response.Response(res, GetProfileResponse{
			Status:  http.StatusInternalServerError,
			Code:    "STA002",
			Message: "Get Device Info Error",
		})
		return
	}

	// var requestBody GetStatusRequest

	// if unmarshalErr := utils.DecodeBody(req, requestBody); unmarshalErr != nil {
	// 	response.Response(res, GetStatusResponse{
	// 		Status:  http.StatusBadRequest,
	// 		Code:    "STA001",
	// 		Message: "Invalid Request",
	// 	})
	// 	return
	// }

	device := DeviceConnect(endpoint.Endpoint) // TODO DB 조회
	result := device.GetStatus(profileToken)

	response.Response(res, result)
	return
}

// Get Configuration
func GetDeviceConfigCtl(res http.ResponseWriter, req *http.Request) {
	cctvId := req.URL.Query().Get("cctv")
	profileToken := req.URL.Query().Get("profile")

	if cctvId == "" || profileToken == "" {
		response.Response(res, GetProfileResponse{
			Status:  http.StatusBadRequest,
			Code:    "COF001",
			Message: "Invalid Params",
		})

		return
	}

	endpoint, getErr := database.GetDeviceInfo(cctvId)
	if getErr != nil {
		response.Response(res, GetProfileResponse{
			Status:  http.StatusInternalServerError,
			Code:    "COF002",
			Message: "Get Device Info Error",
		})
		return
	}
	// var requestBody GetStatusRequest

	// if unmarshalErr := utils.DecodeBody(req, requestBody); unmarshalErr != nil {
	// 	response.Response(res, GetConfigurationResponse{
	// 		Status:  http.StatusBadRequest,
	// 		Code:    "COF001",
	// 		Message: "Invalid Request",
	// 	})
	// 	return
	// }

	device := DeviceConnect(endpoint.Endpoint)
	result := device.GetConfiguration(profileToken)

	response.Response(res, result)
	return
}
