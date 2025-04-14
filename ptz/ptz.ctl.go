package ptz

import (
	"net/http"
	"time"

	"org.donghyuns.com/onvif/ptz/database"
	"org.donghyuns.com/onvif/ptz/response"
	"org.donghyuns.com/onvif/ptz/utils"
)

func RelativeMoveCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody MoveRelativeRequest

	if unmarshalErr := utils.DecodeBody(req, &requestBody); unmarshalErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "RMV001",
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

	if moveErr := device.MoveRelative(requestBody.ProfileToken, requestBody.PanTiltX, requestBody.PanTiltY, requestBody.ZoomX); moveErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "RMV003",
			Message: "Move Relative Error",
		})
		return
	}

	response.Response(res, response.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
	})
}

func ContinouseMoveCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody MoveContinousRequest

	if unmarshalErr := utils.DecodeBody(req, &requestBody); unmarshalErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "CMV001",
			Message: "Invalid Request",
		})
		return
	}

	endpoint, getErr := database.GetDeviceInfo(requestBody.CctvId)

	if getErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "CMV002",
			Message: "Get CCTV Endpoint Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint)
	if moveErr := device.MoveContinuous(
		requestBody.ProfileToken,
		requestBody.PresetToken,
		requestBody.PanTiltX,
		requestBody.PanTiltY,
		requestBody.ZoomX,
		requestBody.IsAbsolute,
		time.Duration(requestBody.Timeout)*time.Second,
	); moveErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "CMV003",
			Message: "Continous Move Error",
		})
		return
	}

	response.Response(res, response.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
	})

	return
}
