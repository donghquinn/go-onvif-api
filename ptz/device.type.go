package ptz

type DeviceCapabilitiesResponseBody struct {
	Response DeviceCapabilitiesResponse `xml:"GetServiceCapabilitiesResponse"`
}

type DeviceCapabilitiesResponse struct {
	Capabilities DeviceCapabilities `xml:"Capabilities"`
}

type DeviceCapabilities struct {
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
