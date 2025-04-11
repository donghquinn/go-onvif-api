package ptz

import "encoding/xml"

type ServiceCapabilitiesResponseBody struct {
	Response ServiceCapabilitiesResponse `xml:"GetServiceCapabilitiesResponse"`
}

type ServiceCapabilitiesResponse struct {
	Capabilities ServiceCapabilities `xml:"Capabilities"`
}

type ServiceCapabilities struct {
	Network  NetworkCapabilities  `xml:"Network"`
	Security SecurityCapabilities `xml:"Security"`
	System   SystemCapabilities   `xml:"System"`
	Misc     MiscCapabilities     `xml:"Misc"`
}

type NetworkCapabilities struct {
	IPFilter            bool `xml:"IPFilter,attr"`
	ZeroConfiguration   bool `xml:"ZeroConfiguration,attr"`
	IPVersion6          bool `xml:"IPVersion6,attr"`
	DynDNS              bool `xml:"DynDNS,attr"`
	Dot11Configuration  bool `xml:"Dot11Configuration,attr"`
	Dot1XConfigurations int  `xml:"Dot1XConfigurations,attr"`
	HostnameFromDHCP    bool `xml:"HostnameFromDHCP,attr"`
	NTP                 int  `xml:"NTP,attr"`
	DHCPv6              bool `xml:"DHCPv6,attr"`
}

type SecurityCapabilities struct {
	TLS10                bool   `xml:"TLS1.0,attr"`
	TLS11                bool   `xml:"TLS1.1,attr"`
	TLS12                bool   `xml:"TLS1.2,attr"`
	OnboardKeyGeneration bool   `xml:"OnboardKeyGeneration,attr"`
	AccessPolicyConfig   bool   `xml:"AccessPolicyConfig,attr"`
	DefaultAccessPolicy  bool   `xml:"DefaultAccessPolicy,attr"`
	Dot1X                bool   `xml:"Dot1X,attr"`
	RemoteUserHandling   bool   `xml:"RemoteUserHandling,attr"`
	X509Token            bool   `xml:"X.509Token,attr"`
	SAMLToken            bool   `xml:"SAMLToken,attr"`
	KerberosToken        bool   `xml:"KerberosToken,attr"`
	UsernameToken        bool   `xml:"UsernameToken,attr"`
	HttpDigest           bool   `xml:"HttpDigest,attr"`
	RELToken             bool   `xml:"RELToken,attr"`
	JsonWebToken         bool   `xml:"JsonWebToken,attr"`
	SupportedEAPMethods  string `xml:"SupportedEAPMethods,attr"`
	MaxUsers             int    `xml:"MaxUsers,attr"`
	MaxUserNameLength    int    `xml:"MaxUserNameLength,attr"`
	MaxPasswordLength    int    `xml:"MaxPasswordLength,attr"`
	HashingAlgorithms    string `xml:"HashingAlgorithms,attr"`
}

type SystemCapabilities struct {
	DiscoveryResolve          bool   `xml:"DiscoveryResolve,attr"`
	DiscoveryBye              bool   `xml:"DiscoveryBye,attr"`
	RemoteDiscovery           bool   `xml:"RemoteDiscovery,attr"`
	SystemBackup              bool   `xml:"SystemBackup,attr"`
	SystemLogging             bool   `xml:"SystemLogging,attr"`
	FirmwareUpgrade           bool   `xml:"FirmwareUpgrade,attr"`
	HttpFirmwareUpgrade       bool   `xml:"HttpFirmwareUpgrade,attr"`
	HttpSystemBackup          bool   `xml:"HttpSystemBackup,attr"`
	HttpSystemLogging         bool   `xml:"HttpSystemLogging,attr"`
	HttpSupportInformation    bool   `xml:"HttpSupportInformation,attr"`
	StorageConfiguration      bool   `xml:"StorageConfiguration,attr"`
	MaxStorageConfigurations  int    `xml:"MaxStorageConfigurations,attr"`
	GeoLocationEntries        int    `xml:"GeoLocationEntries,attr"`
	AutoGeo                   string `xml:"AutoGeo,attr"`
	StorageTypesSupported     string `xml:"StorageTypesSupported,attr"`
	DiscoveryNotSupported     bool   `xml:"DiscoveryNotSupported,attr"`
	NetworkConfigNotSupported bool   `xml:"NetworkConfigNotSupported,attr"`
	UserConfigNotSupported    bool   `xml:"UserConfigNotSupported,attr"`
}

type MiscCapabilities struct {
	AuxiliaryCommands string `xml:"AuxiliaryCommands,attr"`
}

// ===============
type DeviceInformationResponseBody struct {
	Response DeviceInformation `xml:"GetDeviceInformationResponse"`
}

type DeviceInformation struct {
	Manufacturer    string `xml:"Manufacturer"`
	Model           string `xml:"Model"`
	FirmwareVersion string `xml:"FirmwareVersion"`
	SerialNumber    string `xml:"SerialNumber"`
	HardwareId      string `xml:"HardwareId"`
}

// ===============

// GetDeviceCapabilities 요청 관련 구조체
type GetDeviceCapabilities struct {
	XMLName  xml.Name `xml:"tds:GetCapabilities"`
	Category string   `xml:"tds:Category,omitempty"` // 요청할 특정 카테고리 (선택적)
}

// GetDeviceCapabilitiesResponse 응답 구조체
type DeviceCapabilitiesResponseBody struct {
	Response DeviceCapabilitiesResponse `xml:"GetCapabilitiesResponse"`
}

type DeviceCapabilitiesResponse struct {
	Capabilities DeviceCapabilitiesType `xml:"Capabilities"`
}

type DeviceCapabilitiesType struct {
	Analytics AnalyticsCapabilities `xml:"Analytics"`
	Device    DeviceCapabilities    `xml:"Device"`
	Events    EventsCapabilities    `xml:"Events"`
	Imaging   ImagingCapabilities   `xml:"Imaging"`
	Media     MediaCapabilities     `xml:"Media"`
	PTZ       PTZCapabilities       `xml:"PTZ"`
	Extension CapabilitiesExtension `xml:"Extension,omitempty"`
}

// 각 카테고리별 세부 구조체

type AnalyticsCapabilities struct {
	XAddr                  string `xml:"XAddr"`
	RuleSupport            bool   `xml:"RuleSupport,attr"`
	AnalyticsModuleSupport bool   `xml:"AnalyticsModuleSupport,attr"`
}

type DeviceCapabilities struct {
	XAddr     string                `xml:"XAddr"`
	Network   NetworkCapabilities   `xml:"Network"`
	System    SystemCapabilities    `xml:"System"`
	Security  SecurityCapabilities  `xml:"Security"`
	IO        IOCapabilities        `xml:"IO,omitempty"`
	Extension DeviceCapabilitiesExt `xml:"Extension,omitempty"`
}

type EventsCapabilities struct {
	XAddr                                         string `xml:"XAddr"`
	WSSubscriptionPolicySupport                   bool   `xml:"WSSubscriptionPolicySupport,attr"`
	WSPullPointSupport                            bool   `xml:"WSPullPointSupport,attr"`
	WSPausableSubscriptionManagerInterfaceSupport bool   `xml:"WSPausableSubscriptionManagerInterfaceSupport,attr"`
}

type ImagingCapabilities struct {
	XAddr string `xml:"XAddr"`
}

type MediaCapabilities struct {
	XAddr                 string                     `xml:"XAddr"`
	StreamingCapabilities StreamingCapabilities      `xml:"StreamingCapabilities"`
	Extension             MediaCapabilitiesExtension `xml:"Extension,omitempty"`
}

type StreamingCapabilities struct {
	RTPMulticast        bool `xml:"RTPMulticast,attr"`
	RTP_TCP             bool `xml:"RTP_TCP,attr"`
	RTP_RTSP_TCP        bool `xml:"RTP_RTSP_TCP,attr"`
	NonAggregateControl bool `xml:"NonAggregateControl,attr,omitempty"`
}

type PTZCapabilities struct {
	XAddr string `xml:"XAddr"`
}

// IO 기능 관련 구조체
type IOCapabilities struct {
	InputConnectors int                     `xml:"InputConnectors,attr"`
	RelayOutputs    int                     `xml:"RelayOutputs,attr"`
	Extension       IOCapabilitiesExtension `xml:"Extension,omitempty"`
}

// 각종 확장 구조체들
type CapabilitiesExtension struct {
	DeviceIO        DeviceIOCapabilities        `xml:"DeviceIO,omitempty"`
	Display         DisplayCapabilities         `xml:"Display,omitempty"`
	Recording       RecordingCapabilities       `xml:"Recording,omitempty"`
	Search          SearchCapabilities          `xml:"Search,omitempty"`
	Replay          ReplayCapabilities          `xml:"Replay,omitempty"`
	Receiver        ReceiverCapabilities        `xml:"Receiver,omitempty"`
	AnalyticsDevice AnalyticsDeviceCapabilities `xml:"AnalyticsDevice,omitempty"`
	Extensions      CapabilitiesExtension2      `xml:"Extensions,omitempty"`
}

type DeviceCapabilitiesExt struct {
	// 필요에 따라 추가 확장 필드 정의
}

type MediaCapabilitiesExtension struct {
	ProfileCapabilities ProfileCapabilities `xml:"ProfileCapabilities,omitempty"`
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int `xml:"MaximumNumberOfProfiles,attr"`
}

type IOCapabilitiesExtension struct {
	Auxiliary         bool                     `xml:"Auxiliary,attr,omitempty"`
	AuxiliaryCommands string                   `xml:"AuxiliaryCommands,attr,omitempty"`
	Extension         IOCapabilitiesExtension2 `xml:"Extension,omitempty"`
}

type IOCapabilitiesExtension2 struct {
	// 필요에 따라 추가 확장 필드 정의
}

// 나머지 확장 구조체 (필요에 따라 추가 정의)
type DeviceIOCapabilities struct {
	XAddr        string `xml:"XAddr"`
	VideoSources int    `xml:"VideoSources,attr"`
	VideoOutputs int    `xml:"VideoOutputs,attr"`
	AudioSources int    `xml:"AudioSources,attr"`
	AudioOutputs int    `xml:"AudioOutputs,attr"`
	RelayOutputs int    `xml:"RelayOutputs,attr"`
}

type DisplayCapabilities struct {
	XAddr       string `xml:"XAddr"`
	FixedLayout bool   `xml:"FixedLayout,attr"`
}

type RecordingCapabilities struct {
	XAddr              string `xml:"XAddr"`
	ReceiverSource     bool   `xml:"ReceiverSource,attr"`
	MediaProfileSource bool   `xml:"MediaProfileSource,attr"`
	DynamicRecordings  bool   `xml:"DynamicRecordings,attr"`
	DynamicTracks      bool   `xml:"DynamicTracks,attr"`
	MaxStringLength    int    `xml:"MaxStringLength,attr"`
}

type SearchCapabilities struct {
	XAddr          string `xml:"XAddr"`
	MetadataSearch bool   `xml:"MetadataSearch,attr"`
}

type ReplayCapabilities struct {
	XAddr string `xml:"XAddr"`
}

type ReceiverCapabilities struct {
	XAddr                string `xml:"XAddr"`
	RTP_Multicast        bool   `xml:"RTP_Multicast,attr"`
	RTP_TCP              bool   `xml:"RTP_TCP,attr"`
	RTP_RTSP_TCP         bool   `xml:"RTP_RTSP_TCP,attr"`
	SupportedReceivers   int    `xml:"SupportedReceivers,attr"`
	MaximumRTSPURILength int    `xml:"MaximumRTSPURILength,attr"`
}

type AnalyticsDeviceCapabilities struct {
	XAddr       string                   `xml:"XAddr"`
	RuleSupport bool                     `xml:"RuleSupport,attr"`
	Extension   AnalyticsDeviceExtension `xml:"Extension,omitempty"`
}

type AnalyticsDeviceExtension struct {
	// 필요에 따라 추가 확장 필드 정의
}

type CapabilitiesExtension2 struct {
	// 필요에 따라 추가 확장 필드 정의
}
