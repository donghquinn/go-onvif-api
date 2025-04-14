package ptz

import (
	"net/http"

	"org.donghyuns.com/onvif/ptz/database"
	"org.donghyuns.com/onvif/ptz/response"
	"org.donghyuns.com/onvif/ptz/utils"
)

func CreateUserCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody CreateUserRequest

	if unmarshalErr := utils.DecodeBody(req, &requestBody); unmarshalErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "CUR001",
			Message: "Invalid Request",
		})
		return
	}

	endpoint, getErr := database.GetDeviceInfo(requestBody.CctvId)
	if getErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "CUR002",
			Message: "Get Device Info Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint)
	if createErr := device.CreateUser(requestBody.UserName, requestBody.UserId, requestBody.Passwd); createErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "CUR003",
			Message: "Create User Error",
		})

		return
	}

	response.Response(res, response.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
	})
}

func CreateProfileCtl(res http.ResponseWriter, req *http.Request) {
	var requestBody CreateProfileRequest

	if unmarshalErr := utils.DecodeBody(req, &requestBody); unmarshalErr != nil {
		response.Response(res, CreateProfileResponse{
			Status:  http.StatusBadRequest,
			Code:    "CPF001",
			Message: "Invalid Request",
		})
		return
	}

	endpoint, getErr := database.GetDeviceInfo(requestBody.CctvId)
	if getErr != nil {
		response.Response(res, CreateProfileResponse{
			Status:  http.StatusInternalServerError,
			Code:    "CPF002",
			Message: "Get Device Info Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint)

	profileToken, createErr := device.CreateProfile(requestBody.ProfileName)

	if createErr != nil {
		response.Response(res, CreateProfileResponse{
			Status:  http.StatusInternalServerError,
			Code:    "CUR003",
			Message: "Create User Error",
		})

		return
	}

	response.Response(res, CreateProfileResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  profileToken,
	})
}
