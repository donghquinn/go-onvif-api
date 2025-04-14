package ptz

import (
	"net/http"

	"org.donghyuns.com/onvif/ptz/response"
	"org.donghyuns.com/onvif/ptz/utils"
)

func CreatePresetCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody CreatePresetRequest

	if decodeErr := utils.DecodeBody(req, &requestBody); decodeErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "CPT001",
			Message: "Invalid Request",
		})

		return
	}

	device := DeviceConnect("192.168.0.152:10000") // TODO DB 조회
	result := device.CreatePreset(requestBody.ProfileToken, requestBody.PresetName)

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
