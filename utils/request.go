package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func DecodeBody(req *http.Request, instance interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(instance)

	if err != nil {
		log.Printf("[DECODE] Decode JSON Error: %v", err)
		return err
	}

	defer req.Body.Close()

	return nil
}

func ParseBody(req *http.Response, instance interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(instance)

	if err != nil {
		log.Printf("[PARSE] Parsing and Decode JSON Error: %v", err)
		return err
	}

	defer req.Body.Close()

	return nil
}
