package database

import (
	"github.com/donghquinn/gdct"
	"org.donghyuns.com/onvif/ptz/configs"
)

func GetDeviceInfo(cctvId string) (DeviceInfo, error) {
	var deviceInfo DeviceInfo

	dbConfig := configs.DatabaseConfig

	conn, dbErr := gdct.InitPostgresConnection(gdct.DBConfig{
		Host:     dbConfig.Host,
		Port:     dbConfig.Port,
		UserName: dbConfig.User,
		Password: dbConfig.Passwd,
		Database: dbConfig.Database,
	})

	if dbErr != nil {
		return deviceInfo, dbErr
	}

	queryResult, queryErr := conn.PgSelectSingle("SELECT cctv_endpoint FROM m_fa_cctv WHERE cctv_id = $1", cctvId)

	if queryErr != nil {
		return deviceInfo, queryErr
	}

	if scanErr := queryResult.Scan(&deviceInfo.Endpoint); scanErr != nil {
		return deviceInfo, scanErr
	}

	return deviceInfo, nil
}
