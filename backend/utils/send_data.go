package utils

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	//set the status code and send it as response
	w.WriteHeader(statusCode)
	// Creating Encoder object
	encoder := json.NewEncoder(w)
	// Converting text into JSON.
	encoder.Encode(data)
}
