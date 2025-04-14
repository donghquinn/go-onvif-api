package ptz

import (
	"encoding/xml"
	"io"
	"log"

	"github.com/use-go/onvif/ptz"
	"github.com/use-go/onvif/xsd/onvif"

	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

func (d *OnvifDevice) GetNodeList() ([]onvif2.PTZNode, error) {
	onvifRes, onvifErr := d.CallMethod(ptz.GetNodes{})

	if onvifErr != nil {
		log.Printf("[GET_NODE_LIST] Call Get Node List Method Error: %v", onvifErr)
		return nil, onvifErr
	}

	ptzBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_NODE_LIST] Read Response Error: %v", readErr)
		return nil, readErr
	}

	var responseBody []onvif2.PTZNode

	if unmarshalErr := xml.Unmarshal(ptzBody, &responseBody); unmarshalErr != nil {
		log.Printf("[GET_NODE_LIST] Unmarshal Error: %v", unmarshalErr)
		return nil, unmarshalErr
	}

	return responseBody, nil
}

func (d *OnvifDevice) GetNodeInfo(nodeToken string) (onvif2.PTZNode, error) {
	onvifRes, onvifErr := d.CallMethod(ptz.GetNode{
		NodeToken: onvif.ReferenceToken(nodeToken),
	})

	if onvifErr != nil {
		log.Printf("[GET_NODE] Call Get Node Method Error: %v", onvifErr)
		return onvif2.PTZNode{}, onvifErr
	}

	ptzBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_NODE] Read Response Error: %v", readErr)
		return onvif2.PTZNode{}, readErr
	}

	var responseBody onvif2.PTZNode

	if unmarshalErr := xml.Unmarshal(ptzBody, &responseBody); unmarshalErr != nil {
		log.Printf("[GET_NODE] Unmarshal Error: %v", unmarshalErr)
		return onvif2.PTZNode{}, unmarshalErr
	}

	return responseBody, nil
}
