package ptz

import (
	"net/http"

	"org.donghyuns.com/onvif/ptz/database"
	"org.donghyuns.com/onvif/ptz/response"
)

func GetNodeListCtl(res http.ResponseWriter, req *http.Request) {
	cctvId := req.URL.Query().Get("cctv")
	if cctvId == "" {
		response.Response(res, GetProfileResponse{
			Status:  http.StatusBadRequest,
			Code:    "NDL001",
			Message: "Invalid Params",
		})

		return
	}

	endpoint, getErr := database.GetDeviceInfo(cctvId)

	if getErr != nil {
		response.Response(res, CreateProfileResponse{
			Status:  http.StatusInternalServerError,
			Code:    "NLT002",
			Message: "Get Device Info Error",
		})
		return
	}

	device := DeviceConnect(endpoint.Endpoint)
	nodeList, getErr := device.GetNodeList()

	if getErr != nil {
		response.Response(res, NodeListResponse{
			Status:  http.StatusInternalServerError,
			Code:    "NLT001",
			Message: "Get Node List Error",
		})
		return
	}

	response.Response(res, NodeListResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  nodeList,
	})
	return
}

func GetNodeDetailCtl(res http.ResponseWriter, req *http.Request) {
	cctvId := req.URL.Query().Get("cctv")
	nodeProfile := req.URL.Query().Get("profile")

	if cctvId == "" || nodeProfile == "" {
		response.Response(res, GetProfileResponse{
			Status:  http.StatusBadRequest,
			Code:    "NDL001",
			Message: "Invalid Params",
		})

		return
	}

	endpoint, getErr := database.GetDeviceInfo(cctvId)

	if getErr != nil {
		response.Response(res, CreateProfileResponse{
			Status:  http.StatusInternalServerError,
			Code:    "NLT002",
			Message: "Get Device Info Error",
		})
		return
	}
	// var requestBody NodeDetailRequest

	// if decodeErr := utils.DecodeBody(req, &requestBody); decodeErr != nil {
	// 	response.Response(res, NodeDetailResponse{
	// 		Status:  http.StatusBadRequest,
	// 		Code:    "NDL001",
	// 		Message: "Invalid Request",
	// 	})
	// }

	device := DeviceConnect(endpoint.Endpoint)
	nodeData, getErr := device.GetNodeInfo(nodeProfile)

	if getErr != nil {
		response.Response(res, NodeDetailResponse{
			Status:  http.StatusInternalServerError,
			Code:    "NDL002",
			Message: "Get Node Data Error",
		})
		return
	}

	response.Response(res, NodeDetailResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  nodeData,
	})
	return
}
