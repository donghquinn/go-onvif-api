package ptz

import (
	"encoding/xml"
	"io"
	"log"

	"github.com/use-go/onvif/ptz"
	"github.com/use-go/onvif/xsd/onvif"
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

	var responseBody DefaultResponse[GetNodesResponseBodyContent]

	if unmarshalErr := xml.Unmarshal(ptzBody, &responseBody); unmarshalErr != nil {
		log.Printf("[GET_NODE_LIST] Unmarshal Error: %v", unmarshalErr)
		return nil, unmarshalErr
	}

	return responseBody.Body.GetNodesResponse.Node, nil
}

func (d *OnvifDevice) GetNodeInfo(nodeToken string) (PTZNode, error) {
	onvifRes, onvifErr := d.CallMethod(ptz.GetNode{
		NodeToken: onvif.ReferenceToken(nodeToken),
	})

	if onvifErr != nil {
		log.Printf("[GET_NODE] Call Get Node Method Error: %v", onvifErr)
		return PTZNode{}, onvifErr
	}

	ptzBody, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_NODE] Read Response Error: %v", readErr)
		return PTZNode{}, readErr
	}

	var responseBody DefaultResponse[GetNodeResponseBody]

	if unmarshalErr := xml.Unmarshal(ptzBody, &responseBody); unmarshalErr != nil {
		log.Printf("[GET_NODE] Unmarshal Error: %v", unmarshalErr)
		return PTZNode{}, unmarshalErr
	}

	return responseBody.Body.GetNodeResponse.Node, nil
}
