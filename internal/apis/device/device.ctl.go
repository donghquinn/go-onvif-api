package deviceapi

import (
	"fmt"
	"log/slog"
	"net/http"

	"org.donghyuns.com/onvif/ptz/internal/response"
	"org.donghyuns.com/onvif/ptz/pkg/database"
)

type DeviceCtlService struct {
	logger *slog.Logger
	cctvId string
}

func NewDeviceCtlService() DeviceCtlService {
	return DeviceCtlService{}
}

// Get Service Capabilities
func (d *DeviceCtlService) GetServiceCapaCtl(res http.ResponseWriter, req *http.Request) {
	d.logger = slog.With("service", "get_device_capacity")

	cctvId := req.URL.Query().Get("cctv")
	if cctvId == "" {
		d.logger.Error("no cctv id provided")
		response.Response(res, http.StatusBadRequest, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "SCP001",
			Message: "Invalid Params",
		})

		return
	}
	d.cctvId = cctvId
	d.logger = d.logger.With("cctv_id", cctvId)

	endpoint, err := database.GetDeviceInfo(cctvId)
	if err != nil {
		d.logger.Error(fmt.Sprintf("get device endpoint from database err: %v", err))
		response.Response(res, http.StatusInternalServerError, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "SCP002",
			Message: "Get Device Info Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint) // TODO DB 조회
	result, err := device.GetServiceCapability()
	if err != nil {
		d.logger.Error(fmt.Sprintf("get device service capacity err: %v", err))
		response.Response(res, http.StatusInternalServerError, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "SCP003",
			Message: "Get Service Capabilities Error",
		})

		return
	}

	response.Response(res, http.StatusOK, response.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  result,
	})
	return
}

// Get Device Info
func (d *DeviceCtlService) GetDeviceInfoCtl(res http.ResponseWriter, req *http.Request) {
	d.logger = slog.With("service", "get_device_info")
	cctvId := req.URL.Query().Get("cctv")
	if cctvId == "" {
		d.logger.Error("no cctv id provided")
		response.Response(res, http.StatusBadRequest, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "DVF001",
			Message: "Invalid Params",
		})

		return
	}

	d.logger = d.logger.With("cctv_id", cctvId)
	d.cctvId = cctvId

	endpoint, err := database.GetDeviceInfo(cctvId)
	if err != nil {
		d.logger.Error(fmt.Sprintf("get device endpoint from database err: %v", err))
		response.Response(res, http.StatusInternalServerError, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "DVF002",
			Message: "Get Device Info Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint) // TODO DB 조회
	result, err := device.GetDeviceInfo()
	if err != nil {
		d.logger.Error(fmt.Sprintf("get device info err: %v", err))
		response.Response(res, http.StatusInternalServerError, response.CommonResponseWithMessage{
			Status: http.StatusInternalServerError,
			Code:   "DVF003",
		})

		return
	}

	response.Response(res, http.StatusOK, response.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  result.Response,
	})
	return
}

// Get Status
func (d *DeviceCtlService) GetDeviceStatusCtl(res http.ResponseWriter, req *http.Request) {
	cctvId := req.URL.Query().Get("cctv")
	profileToken := req.URL.Query().Get("profile")

	if cctvId == "" || profileToken == "" {
		response.Response(res, http.StatusBadRequest, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "STA001",
			Message: "Invalid Params",
		})

		return
	}

	endpoint, getErr := database.GetDeviceInfo(cctvId)
	if getErr != nil {
		response.Response(res, http.StatusInternalServerError, response.CommonResponseWithMessage{
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

	response.Response(res, result.Status, result)
	return
}

// Get Configuration
func (d *DeviceCtlService) GetDeviceConfigCtl(res http.ResponseWriter, req *http.Request) {
	cctvId := req.URL.Query().Get("cctv")
	profileToken := req.URL.Query().Get("profile")

	if cctvId == "" || profileToken == "" {
		response.Response(res, http.StatusBadRequest, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "COF001",
			Message: "Invalid Params",
		})

		return
	}

	endpoint, getErr := database.GetDeviceInfo(cctvId)
	if getErr != nil {
		response.Response(res, http.StatusInternalServerError, response.CommonResponseWithMessage{
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

	response.Response(res, result.Status, result)
	return
}
