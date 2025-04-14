package ptz

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/media"
	onvif2 "github.com/use-go/onvif/xsd/onvif"
	"org.donghyuns.com/onvif/ptz/utils"
)

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

	if onvifRes.StatusCode != http.StatusOK {
		log.Printf("[CREATE_USER] Create User Status Code Error: %v", onvifRes.StatusCode)
		return fmt.Errorf("create user response failed. %v", onvifRes.StatusCode)
	}

	return nil
}

/*
Create Profile
@name: Profile Name
@Return: Profile Token
*/
func (d *OnvifDevice) CreateProfile(name string) (string, error) {
	referenceToken := utils.CreateToken()

	onvifRes, onvifErr := d.CallMethod(media.CreateProfile{
		Name:  onvif2.Name(name),
		Token: onvif2.ReferenceToken(referenceToken),
	})

	if onvifErr != nil {
		log.Printf("[CREATE_PROFILE] Create User Method Error: %v", onvifErr)
		return "", nil
	}

	if onvifRes.StatusCode != http.StatusOK {
		log.Printf("[CREATE_PROFILE] Create User Status Code Error: %v", onvifErr)
		return "", fmt.Errorf("create user response failed. %v", onvifRes.StatusCode)
	}

	return referenceToken, nil
}

func (d *OnvifDevice) GetProfile(token string) (onvif2.Profile, error) {
	onvifRes, onvifErr := d.CallMethod(media.GetProfile{
		ProfileToken: onvif2.ReferenceToken(token),
	})

	if onvifErr != nil {
		log.Printf("[GET_PROFILE] Create Profile Method Error: %v", onvifErr)
		return onvif2.Profile{}, onvifErr
	}

	if onvifRes.StatusCode != http.StatusOK {
		log.Printf("[GET_PROFILE] Create Profile Status Code Error: %v", onvifErr)
		return onvif2.Profile{}, fmt.Errorf("create user response failed. %v", onvifRes.StatusCode)
	}

	response, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_PROFILE] Create Profile Method Error: %v", onvifErr)
		return onvif2.Profile{}, readErr
	}

	var profileRes onvif2.Profile

	if unmarshal := xml.Unmarshal(response, &profileRes); unmarshal != nil {
		log.Printf("[GET_PROFILE] Unmarshal Profile: %v", unmarshal)
		return onvif2.Profile{}, unmarshal
	}

	return profileRes, nil
}

func (d *OnvifDevice) GetUserList() ([]onvif2.User, error) {
	onvifRes, onvifErr := d.CallMethod(device.GetUsers{})

	if onvifErr != nil {
		log.Printf("[GET_USER_LIST] Create User List Method Error: %v", onvifErr)
		return nil, onvifErr
	}

	if onvifRes.StatusCode != http.StatusOK {
		log.Printf("[GET_USER_LIST] Create User List Status Code Error: %v", onvifErr)
		return nil, fmt.Errorf("create user list response failed. %v", onvifRes.StatusCode)
	}

	response, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_USER_LIST] Read Response Error: %v", readErr)
		return nil, readErr
	}

	var userListResponse []onvif2.User

	if marshalErr := xml.Unmarshal(response, &userListResponse); marshalErr != nil {
		log.Printf("[GET_USER_LIST] Unmarshal XML Error: %v", marshalErr)
	}

	return userListResponse, nil
}
