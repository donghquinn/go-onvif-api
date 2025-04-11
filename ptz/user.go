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
	"org.donghyuns.com/onvif/ptz/util"
)

func (d *OnvifDevice) GetProfile(token string) (Profile, error) {
	onvifRes, onvifErr := d.CallMethod(media.GetProfile{
		ProfileToken: onvif2.ReferenceToken(token),
	})

	if onvifErr != nil {
		log.Printf("[GET_PROFILE] Create Profile Method Error: %v", onvifErr)
		return Profile{}, onvifErr
	}

	if onvifRes.StatusCode != http.StatusOK {
		log.Printf("[GET_PROFILE] Create Profile Status Code Error: %v", onvifErr)
		return Profile{}, fmt.Errorf("create user response failed. %v", onvifRes.StatusCode)
	}

	response, readErr := io.ReadAll(onvifRes.Body)

	if readErr != nil {
		log.Printf("[GET_PROFILE] Create Profile Method Error: %v", onvifErr)
		return Profile{}, readErr
	}

	var profileRes DefaultResponse[GetProfileResponseBody]

	if unmarshal := xml.Unmarshal(response, &profileRes); unmarshal != nil {
		log.Printf("[GET_PROFILE] Unmarshal Profile: %v", unmarshal)
		return Profile{}, unmarshal
	}

	log.Printf("[GET_PROFILE] ProfileList Res: %v", profileRes.Body.GetProfileResponse.Profile)

	return profileRes.Body.GetProfileResponse.Profile, nil
}

/*
Create Profile
@name: Profile Name
@Return: Profile Token
*/
func (d *OnvifDevice) CreateProfile(name string) (string, error) {
	referenceToken := util.CreateToken()

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

func (d *OnvifDevice) GetUserList() ([]User, error) {
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

	var userListResponse DefaultResponse[GetUserResponseBody]

	if marshalErr := xml.Unmarshal(response, &userListResponse); marshalErr != nil {
		log.Printf("[GET_USER_LIST] Unmarshal XML Error: %v", marshalErr)
	}

	return userListResponse.Body.GetUsersResponse.User, nil
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

	if onvifRes.StatusCode != http.StatusOK {
		log.Printf("[CREATE_USER] Create User Status Code Error: %v", onvifErr)
		return fmt.Errorf("create user response failed. %v", onvifRes.StatusCode)
	}

	return nil
}
