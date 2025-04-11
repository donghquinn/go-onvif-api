package ptz

import (
	"encoding/xml"
	"io"
	"log"

	"github.com/use-go/onvif/ptz"
)

func (d *OnvifDevice) GetNodeList() ([]PTZNode, error) {
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

	var responseBody GetNodesResponseBody

	if unmarshalErr := xml.Unmarshal(ptzBody, &responseBody); unmarshalErr != nil {
		log.Printf("[GET_NODE_LIST] Unmarshal Error: %v", unmarshalErr)
		return nil, unmarshalErr
	}

	return responseBody.Body.GetNodesResponse.PTZNode, nil
}
