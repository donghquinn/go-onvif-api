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

	result := device.SetPreset(requestBody.ProfileToken, requestBody.PresetName)

	if result != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "CPT002",
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
	var requestBody GetPresetListRequest

	if decodeErr := utils.DecodeBody(req, &requestBody); decodeErr != nil {
		response.Response(res, PresetListResponse{
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
	result, getErr := device.GetPresetList(requestBody.ProfileToken)

	if getErr != nil {
		response.Response(res, PresetListResponse{
			Status:  http.StatusInternalServerError,
			Code:    "APT002",
			Message: "Get Preset List Error",
		})
		return
	}

	response.Response(res, PresetListResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  result,
	})
}
