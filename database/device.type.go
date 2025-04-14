package database

type GetDeviceInfoRequest struct {
	CctvId string `json:"cctvId"`
}

type GetDeviceInfoResponse struct {
	Status  int        `json:"status"`
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Result  DeviceInfo `json:"result"`
}

type DeviceInfo struct {
	Endpoint string `json:"endpoint"`
}
