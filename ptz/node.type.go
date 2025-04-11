package ptz

import "encoding/xml"

type GetNodesResponseBodyContent struct {
	GetNodesResponse GetNodesResponse `xml:"GetNodesResponse"`
}

type GetNodeResponseBody struct {
	GetNodeResponse GetNodeResponse `xml:"GetNodeResponse"`
}

type GetNodeResponse struct {
	Node PTZNode `xml:"PTZNode"`
}

type GetNodesResponse struct {
	Node []PTZNode `xml:"PTZNode"`
}

// PTZNode는 PTZ 노드의 핵심 정보를 포함합니다
type PTZNode struct {
	NodeToken              string           `xml:"token,attr"`
	FixedHomePosition      bool             `xml:"FixedHomePosition,attr"`
	GeoMove                bool             `xml:"GeoMove,attr"`
	Name                   string           `xml:"Name"`
	SupportedPTZSpaces     PTZSpaces        `xml:"SupportedPTZSpaces"`
	MaximumNumberOfPresets int              `xml:"MaximumNumberOfPresets"`
	HomeSupported          bool             `xml:"HomeSupported"`
	AuxiliaryCommands      []string         `xml:"AuxiliaryCommands"`
	Extension              PTZNodeExtension `xml:"Extension"`
}

// PTZSpaces는 지원되는 모든 PTZ 공간 정의를 포함합니다
type PTZSpaces struct {
	AbsolutePanTiltPositionSpace    SpaceDescription2D `xml:"AbsolutePanTiltPositionSpace"`
	AbsoluteZoomPositionSpace       SpaceDescription1D `xml:"AbsoluteZoomPositionSpace"`
	RelativePanTiltTranslationSpace SpaceDescription2D `xml:"RelativePanTiltTranslationSpace"`
	RelativeZoomTranslationSpace    SpaceDescription1D `xml:"RelativeZoomTranslationSpace"`
	ContinuousPanTiltVelocitySpace  SpaceDescription2D `xml:"ContinuousPanTiltVelocitySpace"`
	ContinuousZoomVelocitySpace     SpaceDescription1D `xml:"ContinuousZoomVelocitySpace"`
	PanTiltSpeedSpace               SpaceDescription1D `xml:"PanTiltSpeedSpace"`
	ZoomSpeedSpace                  SpaceDescription1D `xml:"ZoomSpeedSpace"`
}

// SpaceDescription1D는 1차원 공간(Zoom)을 정의합니다
type SpaceDescription1D struct {
	URI    string     `xml:"URI"`
	XRange FloatRange `xml:"XRange"`
}

// SpaceDescription2D는 2차원 공간(PanTilt)을 정의합니다
type SpaceDescription2D struct {
	URI    string     `xml:"URI"`
	XRange FloatRange `xml:"XRange"`
	YRange FloatRange `xml:"YRange"`
}

// FloatRange는 공간의 최소/최대 값 범위를 정의합니다
type FloatRange struct {
	Min float64 `xml:"Min"`
	Max float64 `xml:"Max"`
}

// PTZNodeExtension은 확장 기능을 포함합니다
type PTZNodeExtension struct {
	SupportedPresetTour PresetTourSupport `xml:"SupportedPresetTour"`
}

// PresetTourSupport는 프리셋 투어 기능을 정의합니다
type PresetTourSupport struct {
	MaximumNumberOfPresetTours int      `xml:"MaximumNumberOfPresetTours"`
	PTZPresetTourOperation     []string `xml:"PTZPresetTourOperation"`
}

// 프리셋 생성 및 관리를 위한 구조체들

// CreatePreset 요청 구조체
type CreatePreset struct {
	XMLName      xml.Name `xml:"tptz:CreatePreset"`
	ProfileToken string   `xml:"tptz:ProfileToken"`
	PresetName   string   `xml:"tptz:PresetName,omitempty"`
	PresetToken  string   `xml:"tptz:PresetToken,omitempty"`
}

// CreatePresetResponse 응답 구조체
type CreatePresetResponse struct {
	PresetToken string `xml:"tptz:PresetToken"`
}

// PTZPosition은 프리셋 위치 정보를 담고 있습니다
type PTZPosition struct {
	PanTilt Vector2D `xml:"PanTilt"`
	Zoom    Vector1D `xml:"Zoom"`
}

// Vector2D는 2차원 벡터(PanTilt)를 정의합니다
type Vector2D struct {
	X     float64 `xml:"x,attr"`
	Y     float64 `xml:"y,attr"`
	Space string  `xml:"space,attr"`
}

// Vector1D는 1차원 벡터(Zoom)를 정의합니다
type Vector1D struct {
	X     float64 `xml:"x,attr"`
	Space string  `xml:"space,attr"`
}
