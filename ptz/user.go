package ptz

import (
	"encoding/xml"
	"io"
	"log"

	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/media"
	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

func (d *OnvifDevice) GetProfile(token string) error {
	onvifRes, onvifErr := d.CallMethod(media.GetProfile{
		ProfileToken: onvif2.ReferenceToken(token),
	})

	if onvifErr != nil {
		log.Printf("[GET_PROFILE] Create User Method Error: %v", onvifErr)
		return nil
	}

	response, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_PROFILE] Create User Method Error: %v", onvifErr)
		return nil
	}

	var profileRes DefaultResponse[GetProfileResponseBody]

	if unmarshal := xml.Unmarshal(response, &profileRes); unmarshal != nil {
		log.Printf("[GET_PROFILE] Unmarshal Profile: %v", unmarshal)
		return unmarshal
	}

	log.Printf("[GET_PROFILE] ProfileList Res: %v", profileRes.Body.GetProfileResponse.Profile)

	return nil
}

func (d *OnvifDevice) CreateProfile(name string, token string) error {
	onvifRes, onvifErr := d.CallMethod(media.CreateProfile{
		Name:  onvif2.Name(name),
		Token: onvif2.ReferenceToken(token),
	})

	if onvifErr != nil {
		log.Printf("[CREATE_PROFILE] Create User Method Error: %v", onvifErr)
		return nil
	}

	response, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[CREATE_PROFILE] Create User Method Error: %v", onvifErr)
		return nil
	}

	log.Printf("[CREATE_PROFILE] ProfileList Res: %v", string(response))
	return nil
}

func (d *OnvifDevice) GetUserList() []onvif2.User {
	onvifRes, onvifErr := d.CallMethod(device.GetUsers{})

	if onvifErr != nil {
		log.Printf("[CREATE_USER] Create User Method Error: %v", onvifErr)
		return nil
	}

	response, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[CREATE_USER] Read Response Error: %v", readErr)
		return nil
	}

	var userListResponse DefaultResponse[GetUserResponseBody]

	if marshalErr := xml.Unmarshal(response, &userListResponse); marshalErr != nil {
		log.Printf("[GET_USER_LIST] Unmarshal XML Error: %v", marshalErr)
	}

	log.Printf("[GET_USER_LIST] Get User List Response: %v", userListResponse.Body.GetUsersResponse.User)

	return nil
}

func (d *OnvifDevice) CreateUser(userName string, userId string, passwd string) error {
	onvifRes, onvifErr := d.CallMethod(device.CreateUsers{
		User: onvif2.User{
			Username: userName,
			Password: passwd,
		},
	})

	if onvifErr != nil {
		log.Printf("[CREATE_USER] Create User Method Error: %v", onvifErr)
		return onvifErr
	}

	_, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[CREATE_USER] Read Response Error: %v", readErr)
		return readErr
	}

	return nil
}
