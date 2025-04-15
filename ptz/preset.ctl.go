package ptz

import (
	"net/http"

	"org.donghyuns.com/onvif/ptz/database"
	"org.donghyuns.com/onvif/ptz/response"
	"org.donghyuns.com/onvif/ptz/utils"
)

// Set Preset Controller
func SetPresetCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody SetPresetRequest

	if decodeErr := utils.DecodeBody(req, &requestBody); decodeErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "CPT001",
			Message: "Invalid Request",
		})

		return
	}

	endpoint, getErr := database.GetDeviceInfo(requestBody.CctvId)
	if getErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "RMV002",
			Message: "Get CCTV Endpoint Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint)

	presetToken, result := device.SetPreset(requestBody.ProfileToken, requestBody.PresetName)

	if result != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "CPT002",
			Message: "Create Preset Error",
		})
		return
	}

	response.Response(res, SetPresetResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  presetToken,
	})
}

// Apply Preset
func ApplyPresetCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody ApplyPresetRequest

	if decodeErr := utils.DecodeBody(req, &requestBody); decodeErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "APT001",
			Message: "Invalid Request",
		})

		return
	}

	endpoint, getErr := database.GetDeviceInfo(requestBody.CctvId)
	if getErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "RMV002",
			Message: "Get CCTV Endpoint Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint) // TODO DB 조회
	result := device.ApplyPreset(requestBody.ProfileToken, requestBody.PresetToken, requestBody.PanTiltX, requestBody.PanTiltY, requestBody.ZoomX, requestBody.IsAbsolute)

	if result != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "APT002",
			Message: "Create Preset Error",
		})
		return
	}

	response.Response(res, response.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
	})
}

// Get Preset List
func GetPresetListCtl(res http.ResponseWriter, req *http.Request) {
	// var requestBody GetPresetListRequest

	// if decodeErr := utils.DecodeBody(req, &requestBody); decodeErr != nil {
	// 	response.Response(res, PresetListResponse{
	// 		Status:  http.StatusBadRequest,
	// 		Code:    "APT001",
	// 		Message: "Invalid Request",
	// 	})

	// 	return
	// }
	cctvId := req.URL.Query().Get("cctv")
	profileToken := req.URL.Query().Get("profile")
	if cctvId == "" || profileToken == "" {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "RMV002",
			Message: "Invalid Params",
		})

		return
	}

	endpoint, getErr := database.GetDeviceInfo(cctvId)
	if getErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "RMV002",
			Message: "Get CCTV Endpoint Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint) // TODO DB 조회
	result, getErr := device.GetPresetList(profileToken)

	if getErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "APT002",
			Message: "Get Preset List Error",
		})
		return
	}

	if result == nil || len(result) == 0 {
		result = []Preset{}
	}

	response.Response(res, PresetListResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  result,
	})
}
