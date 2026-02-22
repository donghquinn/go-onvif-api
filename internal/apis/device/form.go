package deviceapi

type GetDeviceInfoRequest struct {
	CctvId string `json:"cctvId"`
}

type DeviceInfo struct {
	Endpoint string `json:"endpoint"`
}
