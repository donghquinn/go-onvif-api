package database

import (
	"net/http"

	"org.donghyuns.com/onvif/ptz/response"
	"org.donghyuns.com/onvif/ptz/utils"
)

func GetDeviceInfoCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody GetDeviceInfoRequest

	if unmarshalErr := utils.DecodeBody(req, &requestBody); unmarshalErr != nil {
		response.Response(res, GetDeviceInfoResponse{
			Status:  http.StatusBadRequest,
			Code:    "GDI001",
			Message: "Invalid Request",
		})
		return
	}

	result, getErr := GetDeviceInfo(requestBody.CctvId)

	if getErr != nil {
		response.Response(res, GetDeviceInfoResponse{
			Status:  http.StatusInternalServerError,
			Code:    "GDI002",
			Message: "Get Device Info from DB Error",
		})
		return
	}

	response.Response(res, GetDeviceInfoResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  result,
	})
}
