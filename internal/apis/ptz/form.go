package ptzapi

type SetDefaultPositionRequest struct {
	CctvId       string `json:"cctvId"`
	ProfileToken string `json:"profileToken"`
}

type MoveToDefaultPositionRequest struct {
	CctvId       string  `json:"cctvId"`
	ProfileToken string  `json:"profileToken"`
	PanTiltX     float64 `json:"panTiltX"` // 아마 기존 위치로 돌아갈 때 x값 변화량인듯
	PanTiltY     float64 `json:"panTiltY"` // 아마 기존 위치로 돌아갈 때 y값 변화량인듯
	ZoomX        float64 `json:"zoomX"`    // 아마 기존 위치로 돌아갈 때 zoom의 x값 변화량인듯
	IsAbsolute   bool    `json:"isAbsolute"`
}

type MoveRelativeRequest struct {
	CctvId       string  `json:"cctvId"`
	ProfileToken string  `json:"profileToken"`
	PanTiltX     float64 `json:"panTiltX"`
	PanTiltY     float64 `json:"panTiltY"`
	ZoomX        float64 `json:"zoomX"`
}

type MoveContinousRequest struct {
	CctvId       string  `json:"cctvId"`
	ProfileToken string  `json:"profileToken"`
	PanTiltX     float64 `json:"panTiltX"`
	PanTiltY     float64 `json:"panTiltY"`
	ZoomX        float64 `json:"zoomX"`
	IsAbsolute   bool    `json:"isAbsolute"`
	Timeout      int     `json:"timeout"` // Second
}

type GetStatusRequest struct {
	CctvId       string `json:"cctvId"`
	ProfileToken string `json:"profileToken"`
}

type GetConfigurationRequest struct {
	ProfileToken string `json:"profileToken"`
}
