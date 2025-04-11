# PTZ Controller Application
* SOAP Request with ONVIF
* It's a example of ONVIF requests and usages 

## Dependencies
* https://github.com/use-go/onvif

## Usage

### User and Profile
#### Create Profile
* Need profile name and reference token

#### Get Profile
* Need reference token(Profile Token) which is provided when creating profile

---

### Device
#### Device Info
* Firmware, SerialNumber, HardwareId, and Manufacturer

#### device Capabilities
* Network, Security, System, and Misc

---

### PTZ
#### Create Preset
* Require profile token and preset name

#### Preset List
* Require Profile Token

#### Apply Preset
* Require profile token and preset token
* Require Pan-Tilt values and Zoom values
    * X: X value on PTZ Vector
    * Y: Y value on PTZ Vector
    * Space: 