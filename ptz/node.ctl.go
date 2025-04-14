package ptz

import (
	"net/http"

	"org.donghyuns.com/onvif/ptz/response"
	"org.donghyuns.com/onvif/ptz/utils"
)

func GetNodeListCtl(res http.ResponseWriter, req *http.Request) {
	device := DeviceConnect("192.168.0.152:10000")
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
	var requestBody NodeDetailRequest

	if decodeErr := utils.DecodeBody(req, &requestBody); decodeErr != nil {
		response.Response(res, NodeDetailResponse{
			Status:  http.StatusBadRequest,
			Code:    "NDL001",
			Message: "Invalid Request",
		})
	}

	device := DeviceConnect("192.168.0.152:10000")
	nodeData, getErr := device.GetNodeInfo(requestBody.NodeProfile)

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
