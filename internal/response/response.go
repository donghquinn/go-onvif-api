package response

import (
	"encoding/json"
	"net/http"
)

func Response(res http.ResponseWriter, statusCode int, message any) {
	// Set Content-Type header to application/json
	res.Header().Set("Content-Type", "application/json")

	// Marshal the message to JSON
	responseObject, _ := json.Marshal(message)

	res.WriteHeader(statusCode)
	res.Write(responseObject)
}
