package deviceapi

import "fmt"

func (d *DeviceCtlService) GetDeviceEndpointFromDatabase(cctvId string) (DeviceInfo, error) {
	var deviceInfo DeviceInfo

	queryResult, err := d.dbCon.Client.Query("SELECT cctv_endpoint FROM m_fa_cctv WHERE cctv_id = $1", cctvId)
	if err != nil {
		return deviceInfo, fmt.Errorf("query device endpoint err: %v", err)
	}

	if err := queryResult.Scan(&deviceInfo.Endpoint); err != nil {
		return deviceInfo, fmt.Errorf("scan device endpoint err: %v", err)
	}

	return deviceInfo, nil
}
