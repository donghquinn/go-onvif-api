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

func SetDefaultPositionCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody SetDefaultPositionRequest

	if unmarshalErr := utils.DecodeBody(req, &requestBody); unmarshalErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "SDP001",
			Message: "Invalid Request",
		})
		return
	}

	endpoint, getErr := database.GetDeviceInfo(requestBody.CctvId)

	if getErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "SDP002",
			Message: "Get CCTV Endpoint Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint)
	if setErr := device.CreateDefaultPosition(requestBody.ProfileToken); setErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "SDP003",
			Message: "Set Default Position Error",
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

func MoveToDefaultPositionCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody MoveToDefaultPositionRequest

	if unmarshalErr := utils.DecodeBody(req, &requestBody); unmarshalErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "MTD001",
			Message: "Invalid Request",
		})
		return
	}

	endpoint, getErr := database.GetDeviceInfo(requestBody.CctvId)

	if getErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "MTD002",
			Message: "Get CCTV Endpoint Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint)
	if moveErr := device.GoToDefaultPosition(requestBody.ProfileToken, requestBody.PanTiltX, requestBody.PanTiltY, requestBody.ZoomX, requestBody.IsAbsolute); moveErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "MTD003",
			Message: "Move to Default Position Error",
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
