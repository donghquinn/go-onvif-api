package ptz

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/use-go/onvif/ptz"
	"github.com/use-go/onvif/xsd"

	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

// Move Relative
func (d *OnvifDevice) MoveRelative(profileToken string, panTiltX float64, panTiltY float64, zoomX float64) error {
	onvifRes, onvifErr := d.CallMethod(ptz.RelativeMove{
		ProfileToken: onvif2.ReferenceToken(profileToken),
		Translation: onvif2.PTZVector{
			PanTilt: onvif2.Vector2D{
				X:     panTiltX,
				Y:     panTiltY,
				Space: xsd.AnyURI(RelativePanTiltSpace),
			},
			Zoom: onvif2.Vector1D{
				X:     zoomX,
				Space: xsd.AnyURI(RelativeZoomSpace),
			},
		},
	})

	if onvifErr != nil {
		log.Printf("[MOVE_REL] Move Relative Error: %v", onvifErr)
		return onvifErr
	}

	// bodyRead, readErr := io.ReadAll(onvifRes.Body)
	// if readErr != nil {
	// 	return readErr
	// }

	if onvifRes.StatusCode != http.StatusOK {
		log.Printf("[MOVE_REL] Move Relative Status: %v", onvifRes.StatusCode)

		return fmt.Errorf("move relative response error: %v", onvifRes.StatusCode)
	}

	return nil
}

// 지속 이동
func (d *OnvifDevice) MoveContinuous(
	profileToken string,
	panTiltX float64,
	pantiltY float64,
	zoomX float64,
	isAbsolute bool,
	timeout time.Duration,
) error {
	panTiltSpace := AbsolutePanTiltSpace
	zoomSpace := AbsoluteZoomSpace

	if !isAbsolute {
		panTiltSpace = RelativePanTiltSpace
		zoomSpace = RelativeZoomSpace
	}

	panTilt := onvif2.Vector2D{
		X:     panTiltX,
		Y:     pantiltY,
		Space: xsd.AnyURI(panTiltSpace),
	}

	zoom := onvif2.Vector1D{
		X:     zoomX,
		Space: xsd.AnyURI(zoomSpace),
	}

	onvifRes, onvifErr := d.CallMethod(ptz.ContinuousMove{
		ProfileToken: onvif2.ReferenceToken(profileToken),
		Velocity:     onvif2.PTZSpeed{PanTilt: panTilt, Zoom: zoom},
		// Timeout:      xsd.Duration(timeout),
	})

	if onvifErr != nil {
		log.Printf("[MOVE_CONTINUOUS] Call Get Preset Method Error: %v", onvifErr)
	}

	if onvifRes.StatusCode != http.StatusOK {
		return fmt.Errorf("move continuous response error: %v", onvifRes.StatusCode)
	}

	return nil
}

// 설정 확인
func (d *OnvifDevice) GetConfiguration(profileToken string) GetConfigurationResponse {
	onvifRes, onvifErr := d.CallMethod(ptz.GetConfiguration{
		ProfileToken: onvif2.ReferenceToken(profileToken),
	})

	if onvifErr != nil {
		log.Printf("[GET_CONFIG] Call Get Preset Method Error: %v", onvifErr)
		return GetConfigurationResponse{
			Status:  http.StatusInternalServerError,
			Code:    "COF002",
			Message: "Read ONVIF Response Error",
		}
	}

	if onvifRes.StatusCode != http.StatusOK {
		log.Printf("[GET_CONFIG] Response is invalid: %d", onvifRes.StatusCode)
		return GetConfigurationResponse{
			Status:  http.StatusInternalServerError,
			Code:    "COF002",
			Message: "Read ONVIF Response Error",
		}
	}

	ptzBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_CONFIG] Read Response Error: %v", readErr)
		return GetConfigurationResponse{
			Status:  http.StatusInternalServerError,
			Code:    "STA004",
			Message: "Read ONVIF Response Error",
		}
	}

	var configurationResponse GetConfigurationOnvifResponse

	if unmarshal := xml.Unmarshal(ptzBody, &configurationResponse); unmarshal != nil {
		log.Printf("[GET_STATUS] Unmarshal ONVIF Response Error: %v", unmarshal)
		return GetConfigurationResponse{
			Status:  http.StatusInternalServerError,
			Code:    "STA005",
			Message: "Read ONVIF Response Error",
		}
	}

	return GetConfigurationResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "SUCCESS",
		Result:  configurationResponse.Configuration,
	}
}

// Move back to Default Position
func (d *OnvifDevice) GoToDefaultPosition(
	profileToken string,
	panTiltX float64,
	pantiltY float64,
	zoomX float64,
	isAbsolute bool,
) error {
	panTiltSpace := AbsolutePanTiltSpace
	zoomSpace := AbsoluteZoomSpace

	if !isAbsolute {
		panTiltSpace = RelativePanTiltSpace
		zoomSpace = RelativeZoomSpace
	}

	panTilt := onvif2.Vector2D{
		X:     panTiltX,
		Y:     pantiltY,
		Space: xsd.AnyURI(panTiltSpace),
	}

	zoom := onvif2.Vector1D{
		X:     zoomX,
		Space: xsd.AnyURI(zoomSpace),
	}

	onvifRes, onvifErr := d.CallMethod(ptz.GotoHomePosition{
		ProfileToken: onvif2.ReferenceToken(profileToken),
		Speed:        onvif2.PTZSpeed{PanTilt: panTilt, Zoom: zoom},
	})

	if onvifErr != nil {
		log.Printf("[GOTO_HOME] Move Relative Error: %v", onvifErr)
		return onvifErr
	}

	if onvifRes.StatusCode != http.StatusOK {
		return fmt.Errorf("move relative response error: %v", onvifRes.StatusCode)
	}

	return nil
}

// Create Default Position
func (d *OnvifDevice) CreateDefaultPosition(profileToken string) error {
	onvifRes, onvifErr := d.CallMethod(ptz.SetHomePosition{
		ProfileToken: onvif2.ReferenceToken(profileToken),
	})

	if onvifErr != nil {
		log.Printf("[SET_HOME_POS] Move Relative Error: %v", onvifErr)
		return onvifErr
	}

	if onvifRes.StatusCode != http.StatusOK {
		return fmt.Errorf("move relative response error: %v", onvifRes.StatusCode)
	}

	return nil
}
