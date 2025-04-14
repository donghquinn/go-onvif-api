package ptz

import (
	"net/http"

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

	device := DeviceConnect("192.168.0.152:10000")
	if moveErr := device.MoveRelative(requestBody.ProfileToken, requestBody.PanTiltX, requestBody.PanTiltY, requestBody.ZoomX); moveErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "RMV002",
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
